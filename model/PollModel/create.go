package PollModel

import (
	"Amanisis/model"
	"fmt"
)

func Create(data Balance) {
	fmt.Println(data)
	model.MysqlClient.Create(data)
}
