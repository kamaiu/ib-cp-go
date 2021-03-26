//go:generate easyjson -all $GOFILE
package model

type IncomingMessage struct {
	Topic TopicType `json:"topic"`
}
