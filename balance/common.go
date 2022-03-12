package balance

import (
	"Amanisis/model/ServerModel"
	"github.com/gin-gonic/gin"
)

func GetServer(list []ServerModel.Server, c *gin.Context) (num int) {
	return
}

//根据时间戳选择五分钟内活动的服务返回
func GetUsedServer(list []ServerModel.Server) []ServerModel.Server {
	ServerList := make([]ServerModel.Server, 100)
	return ServerList
}
