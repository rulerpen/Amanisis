package register

import (
	"Amanisis/common/dto/registerDto"
	"Amanisis/log"
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
		Ip:   c.ClientIP(),
		Time: int(time.Now().Unix()),
	}
	if ServerModel.IsExist(data.Name, c.ClientIP()) {
		server := ServerModel.Server{
			Name: data.Name,
			Ip:   c.ClientIP(),
			Time: 0,
		}
		ServerModel.Update(server)
		c.JSON(200, "success")
		return
	}
	ServerModel.Create(createData)
	c.JSON(200, "success")
}
