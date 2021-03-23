package ws

type MarketDataHistory struct {
	ServerId          string                  `json:"serverId"`
	Symbol            string                  `json:"symbol"`
	Text              string                  `json:"text"`
	PriceFactor       float64                 `json:"priceFactor"`
	StartTime         string                  `json:"startTime"`
	High              string                  `json:"high"`
	Low               string                  `json:"low"`
	TimePeriod        string                  `json:"timePeriod"`
	BarLength         int64                   `json:"barLength"`
	MdAvailability    string                  `json:"mdAvailability"`
	MktDataDelay      string                  `json:"mktDataDelay"`
	OutsideRth        bool                    `json:"outsideRth"`
	VolumeFactor      int64                   `json:"volumeFactor"`
	PriceDisplayRule  int64                   `json:"priceDisplayRule"`
	PriceDisplayValue string                  `json:"priceDisplayValue"`
	NegativeCapable   bool                    `json:"negativeCapable"`
	MessageVersion    int64                   `json:"messageVersion"`
	Data              []*MarketDataHistoryBar `json:"data"`
	Points            int64                   `json:"points"`
	Topic             TopicType               `json:"topic"`
}

type MarketDataHistoryBar struct {
}
