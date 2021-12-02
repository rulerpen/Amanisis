package registerModel

type Register struct {
	Name     string `gorm:"column:name"`
	Ip       string `gorm:"column:ip"`
	Strategy string `gorm:"strategy"`
	Extra    string `gorm:"extra"`
}

func (Register) TableName() string {
	return "Register"
}
