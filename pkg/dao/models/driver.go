package models

import Err "pkg/error"

type Driver struct {
	UserId     int64  `gorm:"column:user_id" json:"user_id,omitempty"`
	Name       string `gorm:"column:name" json:"name,omitempty"`
	Username   string `gorm:"column:username" json:"username,omitempty"`
	Password   string `gorm:"column:password" json:"password,omitempty"`
	Age        int    `gorm:"column:age" json:"age,omitempty"`
	AccountNum int64  `gorm:"column:account_num" json:"account_num,omitempty"`
}

func (d *Driver) TableName() string {
	return "driver"
}

func GetDriver(username string, password string) (*Driver, error) {
	driver := &Driver{}
	err := db.Model(&Driver{}).Select([]string{"user_id", "name", "username", "password", "age"}).
		Where("username = ?", username).First(driver).Error
	if err != nil || driver.Password != password {
		return nil, Err.New(Err.MysqlNoUserOrWrongPWD, "user not exist or wrong password")
	}
	return driver, err
}