package rest

import (
	"crypto/tls"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
	"time"
	"unsafe"
)

const (
	DefaultUrl = "localhost:5000"
)

const (
	maxRedirectsCount = 16
	acceptEncoding    = "gzip, deflate, br"
)

type StatusCodeError struct {
	Code int
}

func (s StatusCodeError) Error() string {
	return strconv.Itoa(s.Code)
}

type Client struct {
	client *fasthttp.HostClient
	host   string
}

func NewClient(url string, ssl bool) *Client {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c := &fasthttp.HostClient{
		Addr:                          url,
		Name:                          "IB-CP",
		NoDefaultUserAgentHeader:      true,
		Dial:                          fasthttp.Dial,
		DialDualStack:                 false,
		IsTLS:                         ssl,
		TLSConfig:                     &tls.Config{InsecureSkipVerify: true, ServerName: "*"},
		MaxConns:                      120,
		MaxConnDuration:               0,
		MaxIdleConnDuration:           time.Minute * 5,
		MaxIdemponentCallAttempts:     fasthttp.DefaultMaxIdemponentCallAttempts,
		ReadBufferSize:                1024 * 64,
		WriteBufferSize:               1024 * 64,
		ReadTimeout:                   time.Second * 60,
		WriteTimeout:                  time.Second * 60,
		MaxResponseBodySize:           fasthttp.DefaultMaxRequestBodySize,
		DisableHeaderNamesNormalizing: false,
		DisablePathNormalizing:        false,
		MaxConnWaitTimeout:            time.Second * 5,
		RetryIf:                       nil,
	}
	host := ""
	if ssl {
		host = "https://" + url + "/v1/portal"
	} else {
		host = "http://" + url + "/v1/portal"
	}
	return &Client{
		client: c,
		host:   host,
	}
}

type call struct {
	req  *fasthttp.Request
	resp *fasthttp.Response
}

func newCall(
	conn *Client,
	method string,
	url *bytebufferpool.ByteBuffer,
) *call {
	ctx := &call{
		req:  fasthttp.AcquireRequest(),
		resp: fasthttp.AcquireResponse(),
	}
	ctx.req.Header.SetMethod(method)
	ctx.req.SetRequestURIBytes(url.B)
	// Release host buffer
	bytebufferpool.Put(url)
	if conn != nil {
		//if len(conn.agent) > 0 {
		//	ctx.req.Header.Set(fasthttp.HeaderUserAgent, conn.agent)
		//}
		//ctx.req.Header.Set(fasthttp.HeaderAuthorization, conn.auth)
	}
	ctx.req.Header.Set(fasthttp.HeaderContentType, "application/json")
	//ctx.req.Header.Set(fasthttp.HeaderAcceptEncoding, acceptEncoding)
	return ctx
}

func (c *call) release() {
	fasthttp.ReleaseRequest(c.req)
	fasthttp.ReleaseResponse(c.resp)
}

func (c *Client) Do(
	method string,
	url *bytebufferpool.ByteBuffer, // owned
	handle func(status int, body []byte, err error) error,
) error {
	ctx := newCall(c, method, url)
	defer ctx.release()

	scheme := ctx.req.URI().Scheme()
	_ = scheme
	err := c.client.Do(ctx.req, ctx.resp)
	if err != nil {
		if handle != nil {
			return handle(0, nil, err)
		}
		return err
	}
	statusCode := ctx.resp.StatusCode()
	body, err := readBody(ctx.resp)

	if handle != nil {
		return handle(statusCode, body, err)
	}
	return err
}

func readBody(r *fasthttp.Response) (body []byte, err error) {
	if r == nil {
		return nil, nil
	}
	encoding := r.Header.Peek(fasthttp.HeaderContentEncoding)
	if len(encoding) > 0 {
		switch *(*string)(unsafe.Pointer(&encoding)) {
		case "deflate":
			body, err = r.BodyInflate()
		case "brotli":
			body, err = r.BodyUnbrotli()
		case "gzip":
			body, err = r.BodyGunzip()
		default:
			body = r.Body()
		}
	} else {
		body = r.Body()
	}
	return body, err
}
