package DefineServer

import (
	"Amanisis/DefineServer/define1"
	"Amanisis/DefineServer/define2"
	"Amanisis/DefineServer/http"
)

func InitServer() {
	go http.InitHttpServer()
	go define1.InitDefine1Server()
	go define2.InitDefine2Server()
}
