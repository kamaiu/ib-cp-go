//go:generate easyjson -all $GOFILE
package model

import "github.com/kamaiu/ib-cp-go/client/v1/rest/model"

type MarketDataHistoryRequest struct {
	// duration of time back
	// {1-30}min, {1-8}h, {1-1000}d, {1-792}w, {1-182}m, {1-15}y
	Period string `json:"period"`

	// data granularity
	// 	1min, 2min, 3min, 5min, 10min, 15min, 30min, 1h, 2h, 3h, 4h, 8h, 1d, 1w, 1m
	Bar string `json:"bar"`

	// outside regular trading hours
	// true/false
	OutsideRth bool `json:"outsideRth"`

	// The type of data received
	// TRADES, MIDPOINT, BID_ASK
	Source string `json:"source"`

	// historical values returned
	// Can specify multiple separated by forward slash e.g. %o/%c
	Format string `json:"format"`
}

type MarketDataHistoryMessage struct {
	model.HistoryData
	Topic TopicType `json:"topic"`
}
