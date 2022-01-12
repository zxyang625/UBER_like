package config

import (
	kitlog "github.com/go-kit/kit/log"
	"log"
	"os"
)

const (
	System                 = "Driver"
	MethodGetDriverInfo    = "GetDriverInfo"
	MethodTakeOrder 	   = "TakeOrder"
)

var Logger *log.Logger
var KitLogger kitlog.Logger

func init() {
	Logger = log.New(os.Stderr, System, log.LstdFlags)

	KitLogger = kitlog.NewLogfmtLogger(os.Stderr)
	KitLogger = kitlog.WithPrefix(KitLogger, "system", System)
	KitLogger = kitlog.With(KitLogger, "ts", kitlog.DefaultTimestampUTC)
	KitLogger = kitlog.With(KitLogger, "caller", kitlog.DefaultCaller)
}

func GetFmtLogger() *log.Logger {
	logger := Logger
	return logger
}

func GetKitLogger() kitlog.Logger {
	logger := KitLogger
	return logger
}
