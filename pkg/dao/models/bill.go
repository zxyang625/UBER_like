package models

type Bill struct {
	BillNum       int64   `gorm:"column:bill_num; PRIMARY_KEY" json:"bill_num,omitempty"`
	Price         float32 `gorm:"column:price" json:"price,omitempty"`
	StartTime     int64   `gorm:"column:start_time" json:"start_time,omitempty"`
	EndTime       int64   `gorm:"column:end_time" json:"end_time,omitempty"`
	Origin        string  `gorm:"column:origin" json:"origin,omitempty"`
	Destination   string  `gorm:"column:destination" json:"destination,omitempty"`
	PassengerName string  `gorm:"column:passenger_name" json:"passenger_name,omitempty"`
	DriverName    string  `gorm:"column:driver_name" json:"driver_name,omitempty"`
	Payed         bool    `gorm:"column:payed" json:"payed,omitempty"`
	PassengerId   int64   `gorm:"column:passenger_id" json:"passenger_id,omitempty"`
	DriverId      int64 `gorm:"column:driver_id" json:"driver_id,omitempty"`
}

func (b *Bill) TableName() string {
	return "bill"
}

func GetBill(billNum int64) (*Bill, error) {
	bill := &Bill{}
	err := db.Model(&Bill{}).Select([]string{"bill_num", "price", "start_time", "end_time", "origin", "destination", "passenger_name", "driver_name", "payed", "passenger_id", "driver_id"}).
		Where("bill_num = ?", billNum).First(bill).Error
	return bill, err
}

func GetBillList(userId int64) ([]*Bill, error) {
	billList := make([]*Bill, 10)
	err := db.Model(&Bill{}).Where("passenger_id = ?", userId).Find(&billList).Error
	if err != nil {
		return nil, err
	}
	return billList, nil
}

func UpdateBill(billNum int64, bill *Bill) (err error) {
	err = db.Model(&Bill{BillNum: billNum}).Updates(bill).Error
	return
}

func AddBill(bill *Bill) (err error) {
	err = db.Create(bill).Error
	return
}

func DelBill(billNum int64) (err error) {
	err = db.Delete(&Bill{BillNum: billNum}).Error
	return
}

func SetPayedAndGetPrice(billNum int64) (float32, error) {
	err := db.Model(&Bill{}).Where("bill_num = ?", billNum).Update("payed", true).Error
	if err != nil {
		return 0, err
	}
	bill := &Bill{}
	err = db.Model(&Bill{}).Select("price").Where("bill_num = ?", billNum).First(bill).Error
	if err != nil {
		return 0, err
	}
	return bill.Price, nil
}
