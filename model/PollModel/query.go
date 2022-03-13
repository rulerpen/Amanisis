package PollModel

import (
	"Amanisis/model"
	"Amanisis/model/ServerModel"
)

func GetNum(server ServerModel.Server) Balance {
	var balance Balance
	model.MysqlClient.Where("server=?", server.Name).First(&balance)
	return balance
}
