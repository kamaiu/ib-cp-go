package ws

import (
	"context"
	"errors"
	"github.com/fasthttp/websocket"
	"github.com/kamaiu/ib-cp-go/client/v1/ws/model"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

const (
	WsUrl          = "wss://localhost:5000/v1/client/ws"
	MaxWsFrameSize = 1024 * 1024 * 4 // 4mb
)

var (
	ErrNilHandler  = errors.New("nil handler")
	ErrFrameTooBig = errors.New("frame too big")
)

type StatusCodeError struct {
	Code int
}

func (s StatusCodeError) Error() string {
	return strconv.Itoa(s.Code)
}

type WS struct {
	auth      string
	c         *websocket.Conn
	ctx       context.Context
	cancel    context.CancelFunc
	closeCode int
	closeText string

	history map[int64]interface{}
	quotes  map[int64]*MarketDataSubscription
	wg      sync.WaitGroup
	mu      sync.RWMutex
}

func OpenWS(url, auth string) (*WS, error) {
	ctx, cancel := context.WithCancel(context.Background())
	wc, resp, err := websocket.DefaultDialer.DialContext(ctx, url, http.Header{})
	if err != nil {
		cancel()
		return nil, err
	}
	if resp.StatusCode != 200 && resp.StatusCode != 101 {
		cancel()
		return nil, StatusCodeError{resp.StatusCode}
	}

	ws := &WS{
		auth:   auth,
		c:      wc,
		ctx:    ctx,
		cancel: cancel,
	}
	ws.wg.Add(1)
	go ws.run()

	return ws, nil
}

func (ws *WS) run() {
	defer func() {
		ws.wg.Done()
	}()

	c := ws.c

	c.SetCloseHandler(func(code int, text string) error {
		ws.mu.Lock()
		defer ws.mu.Unlock()
		ws.closeCode = code
		ws.closeText = text
		return nil
	})

	var (
		mt, n    int
		r        io.Reader
		err      error
		b        = make([]byte, 4096)
		incoming = &model.IncomingMessage{}
		topic    model.TopicType
		conid    int64
	)
	_ = conid

LOOP:
	for {
		mt, r, err = c.NextReader()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break LOOP
		}
		switch mt {
		case websocket.TextMessage, websocket.BinaryMessage:
			n, err = r.Read(b)
			b = b[:n]
			if err != nil {
				if err == io.EOF {
					err = nil
				}
				break LOOP
			}
			if len(b) == cap(b) {
				for {
					if len(b) == cap(b) {
						if len(b)+1024 >= MaxWsFrameSize {
							err = ErrFrameTooBig
							break LOOP
						}
						b = append(b, make([]byte, 1024)...)[:len(b)]
					}
					n, err = r.Read(b[len(b):cap(b)])
					b = b[:len(b)+n]
					if err != nil {
						if err == io.EOF {
							err = nil
						}
						break LOOP
					}
				}
			}

			err = incoming.UnmarshalJSON(b)
			if err != nil {
				continue
			}

			topic = incoming.Topic
			switch incoming.Topic {
			case model.TopicSystem:
				msg := &model.SystemConnection{}
				err = msg.UnmarshalJSON(b)
			case model.TopicNotifications:
				msg := &model.Notifications{}
				err = msg.UnmarshalJSON(b)
			case model.TopicLiveOrders:
				msg := &model.LiveOrders{}
				err = msg.UnmarshalJSON(b)
			case model.TopicStatus:
				msg := &model.AuthStatus{}
				err = msg.UnmarshalJSON(b)
			case model.TopicProfitLoss:
				msg := &model.PNL{}
				err = msg.UnmarshalJSON(b)
			default:
				t := (string)(topic)
				if strings.Index(t, "smd+") == 0 {
					conid, err = strconv.ParseInt(strings.TrimSpace(t[4:]), 10, 64)
					msg := &model.MarketDataHistoryMessage{}
					err = msg.UnmarshalJSON(b)
				} else if strings.Index((string)(topic), "smh+") == 0 {
					conid, err = strconv.ParseInt(strings.TrimSpace(t[4:]), 10, 64)
					msg := &model.MarketDataMessage{}
					err = msg.UnmarshalJSON(b)
				}
			}

			// Reset buffer
			b = b[0:cap(b)]
		}
	}
}
