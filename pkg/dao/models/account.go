package models

import (
	Err "pkg/error"
)

type Account struct {
	AccountNum  int64   `gorm:"column:account_num; PRIMARY_KEY" json:"account_num,omitempty"`
	PayPassword string  `gorm:"column:pay_password" json:"pay_password,omitempty"`
	Asset       float32 `gorm:"column:asset" json:"asset,omitempty"`
}

func (a Account) TableName() string {
	return "account"
}

func GetAccount(accountNum int64, payPassword string) (*Account, error) {
	account := &Account{}
	db.Model(&Account{}).Select("account_num", "pay_password", "asset").Where("account_num = ?", accountNum).First(&account)
	if account == nil || account.PayPassword != payPassword {
		return nil, Err.New(Err.MysqlNoAccountOrWrongPWD, "no match account or wrong pay password")
	}
	return account, nil
}

func UpdateAccount(accountNum int64, account *Account) (err error) {
	err = db.Model(&Account{AccountNum: accountNum}).Updates(account).Error
	return
}

func AddAccount(account *Account) (err error) {
	err = db.Create(account).Error
	return
}

func DelAccount(accountNum int64) (err error) {
	err = db.Delete(&Account{AccountNum: accountNum}).Error
	return
}

func GetAsset(accountNum int64) (float32, error) {
	account := &Account{}
	err := db.Model(&Account{}).Select("Asset").Where("account_num = ?", accountNum).First(&account).Error
	if err != nil {
		return 0, err
	}
	return account.Asset, nil
}
