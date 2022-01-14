package pb
//
//type GetPassengerInfoRequest struct {
//	Username string `json:"username,omitempty"`
//	Password string `json:"password,omitempty"`
//}
//
//type GetPassengerInfoReply struct {
//	UserId     int64   `json:"user_id,omitempty"`
//	Name       string  `json:"name,omitempty"`
//	Age        int32   `json:"age,omitempty"`
//	AccountNum int64   `json:"account_num,omitempty"`
//	Asset      float32 `json:"asset,omitempty"`
//}
//
//type PublishOrderRequest struct {
//	PassengerId   int64  `json:"passenger_id,omitempty"`
//	StartTime     int64  `json:"start_time,omitempty"`
//	Origin        string `json:"origin,omitempty"`
//	Destination   string `json:"destination,omitempty"`
//	PassengerName string `json:"passenger_name,omitempty"`
//}
//
//type PublishOrderReply struct {
//	Status     bool   `json:"status,omitempty"`
//	DriverName string `json:"driver_name,omitempty"`
//	Location   string `json:"location,omitempty"`
//	Car        string `json:"car,omitempty"`
//	Path       string `json:"path,omitempty"`
//}