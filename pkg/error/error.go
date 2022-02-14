package error

import (
	"fmt"
)

type code int

type Error struct {
	code code
	msg string
}

const (
	DiscoverNewClientFail  = iota + 10000
	DiscoverClientFail
	DiscoverRegisterFail
	DiscoverDeregisterFail
	DiscoverInstanceNotFound

	LoadBalancerInitFail
	LoadBalancerSelectFail

	TracingEmptyService
	TracingNewTracerFail

	MysqlConnectFail
	MysqlNoAccountOrWrongPWD
	MysqlNoUserOrWrongPWD
	MysqlNoEnoughAsset

	MQNewConnectionFail
	MQNewMsgsFail
	MQInitChannelFail
	MQDeclareQueueFail
	MQPublishMsgFail
	MQConsumeMsgFail
	MQSendRespFail
	MQGetRespFail

	RedisPushRequestFail
	RedisBRPOPRequestFail
	RedisTypeConvertFail

	ProtoMarshalFail
	ProtoUnmarshalFail

	RPCRequestTimeout

	ProxyURLInvalid

	InterceptorInvalidCtx
	InterceptorInvalidMeta
)

func Errorf(c code, format string, a ...interface{}) error {
	return New(c, fmt.Sprintf(format, a...))
}

func New(c code, msg interface{}) error {
	switch msg.(type) {
	case string:
		return &Error{
			code: c,
			msg:  msg.(string),
		}
	default:
		return &Error{
			code: c,
			msg:  msg.(error).Error(),
		}
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("[CODE]:%d, [MESSAGE]:%s ", e.code, e.msg)
}