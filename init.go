package main

import (
	"Amanisis/config"
	"Amanisis/log"
	"Amanisis/model"
	"Amanisis/server"
	toml "github.com/pelletier/go-toml"
	"os"
)

func init() {
	//LogFile, _ := os.OpenFile("./log/test.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)
	//log.SetOutput(LogFile)
	conf := new(config.Config)
	TomlFile, _ := os.OpenFile("./config/app.toml", os.O_RDWR, 0644)
	decoder := toml.NewDecoder(TomlFile)
	decoder.Decode(conf)
	log.LogInit(conf.LogServers)
	model.MysqlInit(conf.MysqlServers)
	server.InitServer()
}
