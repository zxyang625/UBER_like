package trip

type GenTripRequest struct {
	PassengerReq *PublishOrderRequest `json:"passenger_req"`
	DriverReq    *TakeOrderRequest    `json:"driver_req"`
}

type GenTripReply struct {
	Status bool     `json:"status,omitempty"`
	Trip   *TripMsg `json:"trip_msg,omitempty"`
}

type PublishOrderRequest struct {
	PassengerId   int64  `json:"passenger_id,omitempty"`
	StartTime     int64  `json:"start_time,omitempty"`
	Origin        string `json:"origin,omitempty"`
	Destination   string `json:"destination,omitempty"`
	PassengerName string `json:"passenger_name,omitempty"`
}

type TakeOrderRequest struct {
	DriverId   int64  `json:"driver_id,omitempty"`
	DriverName string `json:"driver_name,omitempty"`
	Location   string `json:"location,omitempty"`
	Car        string `json:"car,omitempty"`
}

type TripMsg struct {
	TripNum       int64  `json:"trip_num,omitempty"`
	PassengerId   int64  `json:"passenger_id,omitempty"`
	DriverId      int64  `json:"driver_id,omitempty"`
	PassengerName string `json:"passenger_name,omitempty"`
	DriverName    string `json:"driver_name,omitempty"`
	StartTime     int64  `json:"start_time,omitempty"`
	EndTime       int64  `json:"end_time,omitempty"`
	Origin        string `json:"origin,omitempty"`
	Destination   string `json:"destination,omitempty"`
	Car           string `json:"car,omitempty"`
	Path          string `json:"path,omitempty"`
}

