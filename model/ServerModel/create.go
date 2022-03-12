package ServerModel

import (
	"Amanisis/model"
	"fmt"
)

func Create(data Server) {
	fmt.Println(data)
	model.MysqlClient.Create(data)
}
