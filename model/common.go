package model

import (
	"Amanisis/config"
	"fmt"
	"github.com/jinzhu/gorm"
)
var MysqlClient *gorm.DB

func MysqlInit(mysqlConf config.MysqlServer){
	url := mysqlConf.User + ":" + mysqlConf.Password + "@tcp(" + mysqlConf.Host + ":" + fmt.Sprint(mysqlConf.Port)+")/"+mysqlConf.DBName+"?charset=utf8"
	//db, _ := gorm.Open("mysql","root:15955638129@tcp(8.140.30.33:3306)/amanisis?charset=utf8")
	fmt.Println(url)
	MysqlClient, _ = gorm.Open("mysql",url)
}
