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

func IsExist(name string, ip string) bool {
	serverList := make([]Register, 100)
	model.MysqlClient.Where("name=?", name).Where("ip=?", ip).Find(&serverList)
	if len(serverList) > 0 {
		return true
	}
	return false
}
