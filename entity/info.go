package entity

type Info struct {
	Body        string `json:"body"`
	SendTime    int    `json:"send_time"`
	ReceiveTime int    `json:"receive_time"`
	Theme       string `json:"theme"`
}

func NewInfo(body string, sendTime int, receiveTime int, theme string) Info {
	return Info{Body: body, SendTime: sendTime, ReceiveTime: receiveTime, Theme: theme}
}
