package balance

import (
	"Amanisis/model/PollModel"
	"Amanisis/model/ServerModel"
)

func PollBalance(list []ServerModel.Server, ip string) ServerModel.Server {
	list = GetUsedServer(list)
	lens := len(list)
	var curIndex int
	var data PollModel.Balance
	data = PollModel.GetNum(list[0])
	curIndex = (data.Info + 1) % lens
	PollModel.Update(data, curIndex)
	return list[curIndex]

}
