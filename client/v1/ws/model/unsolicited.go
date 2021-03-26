//go:generate easyjson -all $GOFILE
package model

type SystemConnection struct {
	Topic   TopicType `json:"topic"`
	Success string    `json:"success"`
	HB      int64     `json:"hb"`
}

type AuthStatus struct {
	Topic TopicType `json:"topic"`
	Args  struct {
		Authenticated bool `json:"authenticated"`
		Competing     bool `json:"competing"`
	} `json:"args"`
}

type Notifications struct {
	Topic TopicType `json:"topic"`
	Args  struct {
		Id    string `json:"id"`
		Text  string `json:"text"`
		Title string `json:"title"`
		Url   string `json:"url"`
	} `json:"args"`
}

type Bulletins struct {
	Topic TopicType `json:"topic"`
	Args  struct {
		Id      string `json:"id"`
		Message string `json:"message"`
	} `json:"args"`
}
