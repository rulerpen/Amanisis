package ServerModel

import (
	"Amanisis/model"
	"fmt"
)

func GetServerListByName(name string) []Server {
	serverList := make([]Server, 100)
	model.MysqlClient.Where("name=?", name).Find(&serverList)
	fmt.Println(serverList)
	return serverList
}

func IsExist(name string, ip string) bool {
	serverList := make([]Server, 100)
	model.MysqlClient.Where("name=?", name).Where("ip=?", ip).Find(&serverList)
	if len(serverList) > 0 {
		fmt.Println("test", serverList[0])
		return true
	}
	return false
}
