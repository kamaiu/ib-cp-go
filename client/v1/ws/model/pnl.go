//go:generate easyjson -all $GOFILE
package model

type PNL struct {
	Topic TopicType          `json:"topic"`
	Args  map[string]*PNLArg `json:"args"`
}

type PNLArg struct {
	RowType int64   `json:"rowType"`
	DPL     float64 `json:"dpl"`
	UPL     float64 `json:"upl"`
}
