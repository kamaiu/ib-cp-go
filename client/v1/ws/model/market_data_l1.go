//go:generate easyjson -all $GOFILE
package model

import "github.com/kamaiu/ib-cp-go/client/v1/rest/model"

type MarketDataRequest struct {
	Fields []string `json:"fields"`
}

type MarketDataMessage struct {
	model.Iserver_Marketdata_Snapshot_GET_200_List_Item
	Topic TopicType `json:"topic"`
}
