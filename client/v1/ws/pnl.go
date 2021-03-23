package ws

type PNLSubscribe struct{}

func (p PNLSubscribe) Sub() string {
	return "spl+{}"
}

func (p PNLSubscribe) Unsub() string {
	return "upl{}"
}

type PNL struct {
	Topic TopicType          `json:"topic"`
	Args  map[string]*PNLArg `json:"args"`
}

type PNLArg struct {
	RowType int64   `json:"rowType"`
	DPL     float64 `json:"dpl"`
	UPL     float64 `json:"upl"`
}
