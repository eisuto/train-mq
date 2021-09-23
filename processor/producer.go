package processor

import . "TrainMQ/entity"

func Push(info Info, wayMap map[string]Queue){
	q := wayMap[info.Theme]
	q.Push(info)
	wayMap[info.Theme] = q
}
