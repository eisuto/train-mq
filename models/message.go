package models

type Message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Topic   string `json:"topic"`
}
