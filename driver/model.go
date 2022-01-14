package pb
//
//type DriverInfoRequest struct {
//	Username string `json:"username,omitempty"`
//	Password string `json:"password,omitempty"`
//}
//
//type DriverInfoReply struct {
//	UserId     int64   `json:"user_id,omitempty"`
//	Name       string  `json:"name,omitempty"`
//	Age        int32   `json:"age,omitempty"`
//	AccountNum int64   `json:"account_num,omitempty"`
//	Asset      float32 `json:"asset,omitempty"`
//}
//
//type TakeOrderRequest struct {
//	DriverId   int64  `json:"driver_id,omitempty"`
//	DriverName string `json:"driver_name,omitempty"`
//	Location   string `json:"location,omitempty"`
//	Car        string `json:"car,omitempty"`
//}
//
//type TakeOrderReply struct {
//	PassengerName string `json:"passenger_name,omitempty"`
//	StartTime     int64  `json:"start_time,omitempty"`
//	Origin        string `json:"origin,omitempty"`
//	Destination   string `json:"destination,omitempty"`
//	Path          string `json:"path,omitempty"`
//}