package config

import (
	kitlog "github.com/go-kit/kit/log"
	"os"
)

const (
	MethodGenTrip = "GenTrip"

	MethodGetDriverInfo = "GetDriverInfo"
	MethodTakeOrder     = "TakeOrder"

	MethodGetPassengerInfo = "GetPassengerInfo"
	MethodPublishOrder     = "PublishOrder"

	MethodPay = "Pay"

	MethodGenBill             = "GenBill"
	MethodGetBillList         = "GetBillList"
	MethodGetBill             = "GetBill"
	MethodSetPayedAndGetPrice = "SetPayedAndGetPrice"
)

var (
	SystemPayment   = "payment"
	SystemPassenger = "passenger"
	SystemDriver    = "driver"
	SystemTrip      = "trip"
	SystemBilling   = "billing"
)

//var KitLogger kitlog.Logger

//func init() {
//	KitLogger = kitlog.NewLogfmtLogger(os.Stderr)
//	KitLogger = kitlog.With(KitLogger, "ts", kitlog.DefaultTimestampUTC)
//	KitLogger = kitlog.With(KitLogger, "caller", kitlog.DefaultCaller)
//}

func GetKitLogger(system string) kitlog.Logger {
	KitLogger := kitlog.NewLogfmtLogger(os.Stderr)
	KitLogger = kitlog.With(KitLogger, "ts", kitlog.DefaultTimestampUTC)
	KitLogger = kitlog.With(KitLogger, "caller", kitlog.DefaultCaller)
	KitLogger = kitlog.WithPrefix(KitLogger, "system", system)
	return KitLogger
}
