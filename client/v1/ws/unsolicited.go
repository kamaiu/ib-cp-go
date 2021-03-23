package ws

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
	}
}

type Notifications struct {
	Topic TopicType `json:"topic"`
	Args  struct {
		Id    string `json:"id"`
		Text  string `json:"text"`
		Title string `json:"title"`
		Url   string `json:"url"`
	}
}

type Bulletins struct {
	Topic TopicType `json:"topic"`
	Args  struct {
		Id      string `json:"id"`
		Message string `json:"message"`
	}
}
