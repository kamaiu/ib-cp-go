package ws

type LiveOrdersMessage struct {
	Topic TopicType                 `json:"topic"`
	Args  []*LiveOrdersMessageOrder `json:"args"`
}

type LiveOrdersMessageOrder struct {
	Acct                string  `json:"acct"`
	Conid               int64   `json:"conid"`
	OrderId             int64   `json:"orderId"`
	CashCcy             string  `json:"cashCcy"`
	SizeAndFills        string  `json:"sizeAndFills"`
	OrderDesc           string  `json:"orderDesc"`
	Description1        string  `json:"description1"`
	Ticker              string  `json:"ticker"`
	SecType             string  `json:"secType"`
	ListingExchange     string  `json:"listingExchange"`
	RemainingQuantity   float64 `json:"remainingQuantity"`
	FilledQuantity      float64 `json:"filledQuantity"`
	CompanyName         string  `json:"companyName"`
	Status              string  `json:"status"`
	OrigOrderType       string  `json:"origOrderType"`
	SupportsTaxOpt      string  `json:"supportsTaxOpt"`
	LastExecutionTime   string  `json:"lastExecutionTime"`
	LastExecutionTime_r int64   `json:"lastExecutionTime_r"`
	OrderType           string  `json:"orderType"`
	Side                string  `json:"side"`
	TimeInForce         string  `json:"timeInForce"`
	Price               float64 `json:"price"`
	BgColor             string  `json:"bgColor"`
	FgColor             string  `json:"fgColor"`
}
