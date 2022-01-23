package service

import (
	"context"
	"encoding/binary"
	"github.com/streadway/amqp"
	"math"
	"pkg/dao/models"
	Err "pkg/error"
	"time"
)

// PaymentService describes the service.
type PaymentService interface {
	Pay(ctx context.Context, billNum int64, accountNum int64, payPassword string) (msg string, err error)
}

type basicPaymentService struct {

}

func (b *basicPaymentService) Pay(ctx context.Context, billNum int64, accountNum int64, payPassword string) (msg string, err error) {
	account, err := models.GetAccount(accountNum, payPassword)
	if err != nil {
		return "GetAccount fail", err
	}
	data := make([]byte, binary.MaxVarintLen64)
	binary.BigEndian.PutUint64(data, uint64(billNum))
	err = PayMessageServer.Publish(ctx, PublishQueueName, data)
	if err != nil {
		return "pay fail", err
	}
	c := make(chan struct{}, 1)
	d := amqp.Delivery{}
	var price float32
	go func() {
		d, err = PayMessageServer.ReceiveResp(ctx)
		c <- struct{}{}
	}()
	select {
	case <-c:
		if err != nil {
			return "pay fail", err
		}
		bits := binary.LittleEndian.Uint32(d.Body)
		price = math.Float32frombits(bits)
		account.Asset -= price
		err = models.UpdateAccount(accountNum, account)
		if err != nil {
			return "pay fail", err
		}
		return "pay success", nil
	case <-time.After(time.Second):
		return "pay fail", Err.New(Err.RPCRequestTimeout, "PublishOrder timeout")
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

func TxPay(ctx context.Context, account *models.Account, bill *models.Bill) (msg string, err error) {
	tx := models.GetDB().Begin().WithContext(ctx)
	if account.Asset < bill.Price {
		tx.Rollback()
		return "pay fail", Err.New(Err.MysqlNoEnoughAsset, "no enough asset")
	}
	account.Asset -= bill.Price
	err = models.UpdateAccount(account.AccountNum, account)
	if err != nil {
		tx.Rollback()
		return "UpdateAccount fail", err
	}
	bill.Payed = true
	err = models.UpdateBill(bill.BillNum, bill)
	if err != nil {
		tx.Rollback()
		return "UpdateBill fail", err
	}
	tx.Commit()
	return "pay success", nil
}
