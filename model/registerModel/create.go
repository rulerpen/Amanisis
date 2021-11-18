package registerModel

import (
	"Amanisis/model"
	"fmt"
)

func Create(data Register){
	fmt.Println(data)
	model.MysqlClient.Create(data)
}