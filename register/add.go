package register

import (
	"Amanisis/common/dto/registerDto"
	"Amanisis/log"
	"Amanisis/model/PollModel"
	"Amanisis/model/ServerModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Add(c *gin.Context) {
	log.LogClient.Record(c)
	var data registerDto.AddServer
	c.BindJSON(&data)
	fmt.Println(data)
	createData := ServerModel.Server{
		Name: data.Name,
		Ip:   c.ClientIP() + ":" + data.Port,
		Time: int(time.Now().Unix()),
	}
	if ServerModel.IsExist(data.Name, c.ClientIP()+":"+data.Port) {
		server := ServerModel.Server{
			Name: data.Name,
			Ip:   c.ClientIP() + ":" + data.Port,
			Time: 0,
		}
		ServerModel.Update(server)
		c.JSON(200, "success")
		return
	}
	ServerModel.Create(createData)
	balance := PollModel.Balance{
		Server: data.Name,
		Info:   0,
	}
	PollModel.Create(balance)
	c.JSON(200, "success")
}
