package http

import (
	"Amanisis/httpclient"
	"Amanisis/model/ServerModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

func GetUrl(ip string, name string) string {
	strings.ReplaceAll(name, ".", "/")
	str := ip + "/" + name
	return str
}
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
	fmt.Println(serverList, len(serverList))
	fmt.Println(serverList[0], GetUrl(serverList[0].Ip, serverList[0].Name))
	//	fmt.Println(httpclient.Post("127.0.0.1:9001/ping", "1"))
	fmt.Println(GetUrl(serverList[0].Ip, serverList[0].Name))
	body, _ := ioutil.ReadAll(c.Request.Body)
	resp := httpclient.Post(GetUrl(serverList[0].Ip, serverList[0].Name), string(body))
	respMap := httpclient.JSONToMap(resp)
	c.JSON(200, respMap)
}
