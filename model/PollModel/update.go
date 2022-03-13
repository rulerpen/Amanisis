package PollModel

import "Amanisis/model"

func Update(data Balance, turn int) {
	model.MysqlClient.Model(data).Where("server=?", data.Server).Update("info", turn)
}
