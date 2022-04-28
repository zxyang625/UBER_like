package models

import (
	Err "pkg/error"
)

type Passenger struct {
	UserId     int64  `gorm:"column:user_id; PRIMARY_KEY" json:"user_id,omitempty"`
	Name       string `gorm:"column:name" json:"name,omitempty"`
	Username   string `gorm:"column:username" json:"username,omitempty"`
	Password   string `gorm:"column:password" json:"password,omitempty"`
	Age        int    `gorm:"column:age" json:"age,omitempty"`
	AccountNum int64  `gorm:"column:account_num" json:"account_num,omitempty"`
}

func (p *Passenger) TableName() string {
	return "passenger"
}

func GetPassenger(username string, password string) (*Passenger, error) {
	passenger := &Passenger{}
	err := db.Model(&Passenger{}).
		Where("username = ?", username).First(passenger).Error
	if err != nil || passenger.Password != password {
		return nil, Err.New(Err.MysqlNoUserOrWrongPWD, "user not exist or wrong password")
	}
	return passenger, nil
}
