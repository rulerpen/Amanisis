package server

import "Amanisis/server/http"

func InitServer() {
	go http.InitHttpServer()
}
