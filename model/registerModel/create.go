package registerModel

import (
	"github.com/jinzhu/gorm"
)

func Create(data Register){
	db, _ := gorm.Open("mysql","root:15955638129@tcp(8.140.30.33:3306)/amanisis?charset=utf8")

	defer db.Close()



	db.Create(data)
}