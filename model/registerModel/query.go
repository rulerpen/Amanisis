package registerModel

import (
	"Amanisis/model"
	"fmt"
)

func GetServerListByName(name string) []Register {
	serverList := make([]Register, 100)
	model.MysqlClient.Where("name=?", name).Find(&serverList)
	fmt.Println(serverList)
	return serverList
}
