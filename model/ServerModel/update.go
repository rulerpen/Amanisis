package ServerModel

import (
	"Amanisis/model"
	"fmt"
	"time"
)

func Update(data Server) {
	fmt.Println(data)
	model.MysqlClient.Model(data).Where("name=?", data.Name).Where("ip=?", data.Ip).Update("time", time.Now().Unix())
}
