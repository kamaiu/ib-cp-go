package ws

import "github.com/kamaiu/ib-cp-go/client/v1/ws/model"

type MarketDataHistorySubscription struct {
	Conid   int64
	Request model.MarketDataRequest
	handler func(data *model.MarketDataHistoryMessage)
}

func (w *WS) SubscribeDataHistory(
	conid int64,
	request *model.MarketDataRequest,
	handler func(data *model.MarketDataHistoryMessage),
) (error, *MarketDataHistorySubscription) {
	return nil, nil
}
