package ServerModel

type Server struct {
	Name string `gorm:"column:name"`
	Ip   string `gorm:"column:ip"`
	Time int    `gorm:"column:time"`
}

func (Server) TableName() string {
	return "server"
}
