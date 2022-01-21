package redis

import (
	"github.com/golang/protobuf/proto"
	"github.com/gomodule/redigo/redis"
	Err "pkg/error"
	"pkg/pb"
	"time"
)

var (
	client *redis.Pool
)

const (
	passengerListName     = "passenger_list"
	driverListName        = "driver_list"
	tripListName          = "trip_list"
	billListName          = "bill_list"
	defaultReadTimeout    = 0 * time.Second
	defaultWriteTimeout   = 3 * time.Second
	defaultConnectTimeout = 3 * time.Second
)

type Passenger struct {}
type Driver struct {}
type Trip struct {}
type Billing struct {}

func init() {
	client = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "127.0.0.1:6379",
				redis.DialConnectTimeout(defaultConnectTimeout),
				redis.DialReadTimeout(defaultReadTimeout),
				redis.DialWriteTimeout(defaultWriteTimeout))
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
		MaxIdle:         500,
		MaxActive:       500,
		IdleTimeout:     3 * time.Second,
		Wait:            true,
	}
}

func (b Billing) LPUSH(request *pb.BillMsg) error {
	data, err := proto.Marshal(request)
	if err != nil {
		return Err.New(Err.ProtoMarshalFail, err.Error())
	}
	conn := client.Get()
	defer conn.Close()
	_, err = conn.Do("LPUSH", billListName, data)
	if err != nil {
		return Err.New(Err.RedisPushRequestFail, err.Error())
	}
	return nil
}

func (b Billing) BRPOPData() ([]byte, error) {
	conn := client.Get()
	defer conn.Close()
	r, err := redis.Values(conn.Do("BRPOP", billListName, 0))
	if err != nil {
		return nil, Err.New(Err.RedisBRPOPRequestFail, err.Error())
	}
	return r[1].([]byte), nil
}

func (t Trip) LPUSH(request *pb.TripMsg) error {
	data, err := proto.Marshal(request)
	if err != nil {
		return Err.New(Err.ProtoMarshalFail, err.Error())
	}
	conn := client.Get()
	defer conn.Close()
	_, err = conn.Do("LPUSH", tripListName, data)
	if err != nil {
		return Err.New(Err.RedisPushRequestFail, err.Error())
	}
	return nil
}

func (t Trip) BRPOPData() ([]byte, error) {
	conn := client.Get()
	defer conn.Close()
	r, err := redis.Values(conn.Do("BRPOP", tripListName, 0))
	if err != nil {
		return nil, Err.New(Err.RedisBRPOPRequestFail, err.Error())
	}
	return r[1].([]byte), nil
	//reply := &pb.TripMsg{}
	//err = proto.Unmarshal(r[1].([]byte), reply)
	//if err != nil {
	//	return nil, Err.New(Err.ProtoUnmarshalFail, err.Error())
	//}
	//return reply, nil
}

func (p Passenger) LPush(request *pb.PublishOrderRequest) error {
	if request == nil {
		return Err.New(Err.RedisPushRequestFail, "nil LPUSH PublishOrderRequest")
	}
	data, err := proto.Marshal(request)
	if err != nil {
		return Err.New(Err.ProtoMarshalFail, err.Error())
	}
	conn := client.Get()
	defer conn.Close()
	_, err = conn.Do("LPUSH", passengerListName, data)
	if err != nil {
		return Err.New(Err.RedisPushRequestFail, err.Error())
	}
	return nil
}

func (p Passenger) BRPOP() (*pb.PublishOrderRequest, error) {
	conn := client.Get()
	defer conn.Close()
	r, err := redis.Values(conn.Do("BRPOP", passengerListName, 0))
	if err != nil {
		return nil, Err.New(Err.RedisBRPOPRequestFail, err.Error())
	}
	reply := &pb.PublishOrderRequest{}
	err = proto.Unmarshal(r[1].([]byte), reply)
	if err != nil {
		return nil, Err.New(Err.ProtoUnmarshalFail, err.Error())
	}
	return reply, nil
}

func (d Driver) LPush(request *pb.TakeOrderRequest) error {
	if request == nil {
		return Err.New(Err.RedisPushRequestFail, "nil LPUSH TakeOrderRequest")
	}
	data, err := proto.Marshal(request)
	if err != nil {
		return Err.New(Err.ProtoMarshalFail, err.Error())
	}
	conn := client.Get()
	defer conn.Close()
	_, err = conn.Do("LPUSH", driverListName, data)
	if err != nil {
		return Err.New(Err.RedisPushRequestFail, err.Error())
	}
	return nil
}

func (d Driver) BRPOP() (*pb.TakeOrderRequest, error) {
	conn := client.Get()
	defer conn.Close()
	r, err := redis.Values(conn.Do("BRPOP", driverListName, 0))
	if err != nil {
		return nil, Err.New(Err.RedisBRPOPRequestFail, err.Error())
	}
	reply := &pb.TakeOrderRequest{}
	err = proto.Unmarshal(r[1].([]byte), reply)
	if err != nil {
		return nil, Err.New(Err.ProtoUnmarshalFail, err.Error())
	}
	return reply, nil
}