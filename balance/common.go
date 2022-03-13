package balance

import (
	"Amanisis/model/ServerModel"
	"fmt"
	"time"
)

func GetServer(list []ServerModel.Server, ip string) ServerModel.Server {
	a := ServerModel.Server{
		Name: "",
		Ip:   "",
		Time: 0,
	}
	return a
}

//根据时间戳选择五分钟内活动的服务返回
func GetUsedServer(list []ServerModel.Server) []ServerModel.Server {
	ServerList := make([]ServerModel.Server, 0)
	for _, i := range list {
		f := int(time.Now().Unix())
		if f-i.Time < 30000 {
			ServerList = append(ServerList, i)
		} else {
			if f-i.Time < 0 {
				fmt.Println("TimeWrong")
				return nil
			}
		}
	}
	return ServerList
}
