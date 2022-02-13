package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pkg/config"
	"pkg/gateway"
	"syscall"
)

func main() {
	//pass, err := mq.NewMessageServer(mq.PassengerQueueName, mq.TripQueueName, "", "")
	//if err != nil {
	//	fmt.Errorf("new mq failed, err: %v", err)
	//}
	//trip, err := mq.NewMessageServer(mq.TripQueueName, mq.PassengerQueueName, "", "")
	//if err != nil {
	//	fmt.Errorf("new mq failed, err: %v", err)
	//}
	//err = pass.Publish(context.Background(), []byte("hello, this is pass"))
	//if err != nil {
	//	fmt.Errorf("pass publish msg failed, err: %v", err)
	//}
	//time.Sleep(time.Second)
	//err = trip.Consume(context.Background(), func(d amqp.Delivery) {
	//	fmt.Println("trip recv", d.Body)
	//	d.Ack(false)
	//})

	//pass, err := mq.NewMessageServer("notification_queue", "passenger_queue", "", "")
	//if err != nil {
	//	fmt.Errorf("new mq failed, err: %v", err)
	//}
	//err = pass.Consume(context.Background(), "passenger_queue", func(d amqp.Delivery) {
	//	resp := &pb.PublishOrderRequest{}
	//	err := proto.Unmarshal(d.Body, resp)
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println(resp.GetPassengerName(), resp.GetPassengerId(), resp.GetOrigin(), resp.GetDestination(), resp.GetStartTime())
	//	d.Ack(false)
	//})
	//if err != nil {
	//	log.Println(err)
	//}

	// 测试trip的接收和生成
	//go func() {
	//	bill, err := mq.NewMessageServer("billing_queue")
	//	if err != nil {
	//		fmt.Errorf("new mq failed, err: %v", err)
	//	}
	//	err = bill.Consume(context.Background(), "trip_queue", func(d amqp.Delivery) {
	//		resp := &pb.TripMsg{}
	//		err := proto.Unmarshal(d.Body, resp)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		fmt.Println(resp.GetPassengerName(), resp.GetPassengerId(), resp.GetOrigin(), resp.GetDestination(), resp.GetStartTime())
	//		d.Ack(false)
	//	})
	//}()
	//
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		redis.Passenger{}.LPush(&pb.PublishOrderRequest{
	//			PassengerId:          int64(i),
	//			StartTime:            int64(i),
	//			Origin:               "Origin" + strconv.Itoa(i),
	//			Destination:          "Destination" + strconv.Itoa(i),
	//			PassengerName:        "PassengerName" + strconv.Itoa(i),
	//		})
	//	}
	//}()
	//
	//for i := 0; i < 100; i++ {
	//	redis.Driver{}.LPush(&pb.TakeOrderRequest{
	//		DriverId:             int64(i),
	//		DriverName:           "DriverName" + strconv.Itoa(i),
	//		Location:             "Location" + strconv.Itoa(i),
	//		Car:                  "Car" + strconv.Itoa(i),
	//	})
	//}
	//time.Sleep(10 * time.Second)

	var (
		consulHost = flag.String("consul.host", "127.0.0.1", "consul server ip address")
		consulPort = flag.Int("consul.port", 8500, "consul server port")
	)
	logger := config.GetKitLogger("ReverseProxy")

	proxy, err := gateway.NewReverseProxy(*consulHost, *consulPort, logger)
	if err != nil {
		logger.Log("err", err)
		os.Exit(-1)
	}

	errc := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("ReverseProxy", "Listening", "addr", "10000")
		errc <- http.ListenAndServe(":10000", proxy)
	}()

	logger.Log("exit", <-errc)
}
