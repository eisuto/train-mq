package processor

import (
	. "TrainMQ/entity"
)
func Pop(theme string, wayMap map[string]Queue) Info {
	q := wayMap[theme]
	if q.Len() == 0 {
		return NewInfo("No message.", 0, 0, theme)
	}
	val := q.Pop()
	wayMap[theme] = q
	return val.(Info)
}

