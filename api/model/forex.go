package model

type ForexContract struct {
	ID     int64
	Symbol string
}

type ForexQuote struct {
	Bid     float64 `json:"b"`
	Ask     float64 `json:"a"`
	BidSize float64 `json:"d"`
	AskSize float64 `json:"e"`
}

type ForexBar struct {
	Bid     OHLC `json:"b"`
	BidSize OHLC `json:"d"`
	Ask     OHLC `json:"a"`
	AskSize OHLC `json:"e"`
	Spread  OHLC `json:"s"`
	Volume  OHLC `json:"v"`
}
