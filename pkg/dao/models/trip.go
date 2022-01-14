package models

type Trip struct {
	TripNum       int64  `gorm:"column:trip_num" json:"trip_num,omitempty"`
	PassengerId   int64  `gorm:"column:passenger_id" json:"passenger_id,omitempty"`
	DriverId      int64  `gorm:"column:driver_id" json:"driver_id,omitempty"`
	PassengerName string `gorm:"column:passenger_name" json:"passenger_name,omitempty"`
	DriverName    string `gorm:"column:driver_name" json:"driver_name,omitempty"`
	StartTime     string `gorm:"column:start_time" json:"start_time,omitempty"`
	EndTime       string `gorm:"column:end_time" json:"end_time,omitempty"`
	Origin        string `gorm:"column:origin" json:"origin,omitempty"`
	Destination   string `gorm:"column:destination" json:"destination,omitempty"`
	Car           string `gorm:"column:car" json:"car,omitempty"`
	Path          string `gorm:"column:path" json:"path,omitempty"`
}

func (t *Trip) TableName() string {
	return "trip"
}

func GetTrip(tripNum int64) (*Trip, error) {
	trip := &Trip{}
	err := db.Model(&Trip{}).Select([]string{"trip_num", "passenger_id", "driver_id", "passenger_name", "driver_name", "start_time", "end_time", "origin", "destination", "car", "path"}).
		Where("trip_num = ?", tripNum).First(trip).Error
	if err != nil {
		return nil, err
	}
	return trip, nil
}

func AddTrip(trip *Trip) error {
	return db.Create(trip).Error
}

