package model

import "time"

type Stock struct {
	Symbol string  `json:"s"`
	Chains *Chains `json:"c"`
}

type Chains struct {
	Chains []*OptionChain
}

type OptionChain struct {
	Expires time.Time `json:"e"`

	Strikes []*OptionContract `json:"s"`
}

type OptionRight byte

const (
	OptionRightCall = 'C'
	OptionRightPut  = 'P'
)

type OptionContract struct {
	ID        int64       `json:"i"`
	Strike    float64     `json:"s"`
	Right     OptionRight `json:"r"`
	Exchanges string      `json:"e"`
}

type OptionQuote struct {
	Bid        float64 `json:"b"`
	BidSize    float64 `json:"o"`
	Ask        float64 `json:"a"`
	AskSize    float64 `json:"s"`
	Volume     float64 `json:"c"`
	Delta      float64 `json:"d"`
	Gamma      float64 `json:"g"`
	Theta      float64 `json:"t"`
	Vega       float64 `json:"v"`
	Volatility float64 `json:"i"`
}

type OHLC struct {
	Open  float64 `json:"o"`
	High  float64 `json:"h"`
	Low   float64 `json:"l"`
	Close float64 `json:"c"`
}

// Option greeks
type Greeks struct {
	Delta float64 `json:"d"`
	Gamma float64 `json:"g"`
	Theta float64 `json:"t"`
	Vega  float64 `json:"v"`
}

type GreeksOHLC struct {
	Open  Greeks `json:"o"`
	High  Greeks `json:"h"`
	Low   Greeks `json:"l"`
	Close Greeks `json:"c"`
}

type OptionBar struct {
	Time       int64      `json:"t"`
	Bid        OHLC       `json:"b"`
	Ask        OHLC       `json:"a"`
	Bids       float64    `json:"d"`
	Asks       float64    `json:"e"`
	Greeks     GreeksOHLC `json:"g"`
	Volatility OHLC       `json:"v"`
}
