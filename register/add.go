package register

import (
	"Amanisis/common/dto/registerDto"
	"Amanisis/model/registerModel"
	"github.com/gin-gonic/gin"
	"log"
)
func Add(c *gin.Context){
	raw, _:= c.GetRawData()
	rawStr := string(raw)
	log.Println(c.FullPath(),rawStr)
	var data registerDto.AddServer
	c.BindJSON(&data)
	createData := registerModel.Register{
		Name:     data.Name,
		Ip:       c.ClientIP(),
		Strategy: data.Strategy,
	}
	registerModel.Create(createData)
	c.JSON(200,"success")
}
