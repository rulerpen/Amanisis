package PollModel

type Balance struct {
	Server string `gorm:"column:server"`
	Info   int    `gorm:"column:info"`
}

func (Balance) TableName() string {
	return "balance"
}
