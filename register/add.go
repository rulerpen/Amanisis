package register

import (
	"Amanisis/common/dto/registerDto"
	"Amanisis/log"
	"Amanisis/model/registerModel"
	"fmt"
	"github.com/gin-gonic/gin"
)
func Add(c *gin.Context){
	log.LogClient.Record(c)
	var data registerDto.AddServer
	c.BindJSON(&data)
	fmt.Println(data)
	createData := registerModel.Register{
		Name:     data.Name,
		Ip:       c.ClientIP(),
		Strategy: data.Strategy,
	}
	registerModel.Create(createData)
	c.JSON(200,"success")
}
