package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/openzipkin/zipkin-go"
	"github.com/streadway/amqp"
	"net/http"
	"pkg/dao/models"
	"pkg/dao/mq"
	Err "pkg/error"
	"pkg/pb"
	"time"
)

// PaymentService describes the service.
type PaymentService interface {
	Pay(ctx context.Context, billNum int64, accountNum int64, payPassword string) (msg string, err error)
}

type basicPaymentService struct {

}

func (b *basicPaymentService) Pay(ctx context.Context, billNum int64, accountNum int64, payPassword string) (msg string, err error) {
	fmt.Println("ppppp", ctx.Value("Length"))
	fmt.Println("ttttt", ctx.Value("Trace-ID"))
	account, err := models.GetAccount(accountNum, payPassword)
	if err != nil {
		return "GetAccount fail", err
	}
	///////////////////////////////////////////
	data, _ := json.Marshal(&pb.SetPayedAndGetPriceRequest{BillNum: billNum})
	span := zipkin.SpanOrNoopFromContext(ctx)
	req := mq.AsyncReq{
		Method:      http.MethodPost,
		Application: "billing",
		Service:     "set-payed-and-get-price",
		//URL:         "http://localhost:10000/billing/set-payed-and-get-price",
		TraceID:     span.Context().TraceID,
		Priority:    ctx.Value("Length").(int),
		Header:      nil,
		Data:        data,
	}

	mqData, err := json.Marshal(&req)
	if err != nil {
		return "pay fail", err
	}
	/////////////////////////////////////////////
	err = PayMessageServer.Publish(ctx, PublishQueueName, mqData)
	if err != nil {
		return "pay fail", err
	}
	c := make(chan struct{}, 1)
	d := amqp.Delivery{}
	go func() {
		d, err = PayMessageServer.ReceiveResp(ctx)
		c <- struct{}{}
	}()
	select {
	case <-c:
		if err != nil {
			return "pay fail", err
		}
		rsp := pb.SetPayedAndGetPriceReply{}
		_ = json.Unmarshal(d.Body, &rsp)
		account.Asset -= rsp.Price
		err = models.UpdateAccount(accountNum, account)
		if err != nil {
			return "pay fail", err
		}
		return "pay success", nil
	case <-time.After(time.Second):
		return "pay fail", Err.New(Err.RPCRequestTimeout, "pay request timeout")
	}
}

// NewBasicPaymentService returns a naive, stateless implementation of PaymentService.
func NewBasicPaymentService() PaymentService {
	return &basicPaymentService{
	}
}

// New returns a PaymentService with all of the expected middleware wired in.
func New(middleware []Middleware) PaymentService {
	var svc PaymentService = NewBasicPaymentService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
