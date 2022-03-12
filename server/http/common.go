package http

import (
	"Amanisis/model/ServerModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func InitHttpServer() {
	r := gin.Default()
	r.POST("/*path", server)
	r.Run(":8888")
}

func server(c *gin.Context) {
	path := c.Param("path")
	server := strings.ReplaceAll(path, "/", ".")
	server = strings.TrimLeft(server, ".")
	fmt.Println(server)
	serverList := ServerModel.GetServerListByName(server)
	fmt.Println(serverList)
}
