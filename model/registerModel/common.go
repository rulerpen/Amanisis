package registerModel

type Register struct {
	Name string `gorm:"column:name"`
	Ip string `gorm:"column:ip"`
	Strategy string `gorm:"strategy"`
}

func (Register) TableName() string {
	return "Register"
}