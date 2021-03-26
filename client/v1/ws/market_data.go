package ws

import (
	"github.com/kamaiu/ib-cp-go/client/v1/ws/model"
	"os"
	"sync"
)

type MarketDataSubscription struct {
	Conid   int64
	Request model.MarketDataRequest
	handler func(data *model.MarketDataMessage)
	ws      *WS
	mu      sync.Mutex
}

func (w *WS) SubscribeData(
	conid int64,
	request *model.MarketDataRequest,
	handler func(data *model.MarketDataMessage),
) (*MarketDataSubscription, error) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	existing := w.quotes[conid]
	if existing != nil {
		return nil, os.ErrExist
	}
	return nil, nil
}
