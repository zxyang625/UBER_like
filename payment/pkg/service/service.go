package service

import (
	"context"
	"pkg/dao/models"
	Err "pkg/error"
)

// PaymentService describes the service.
type PaymentService interface {
	Pay(ctx context.Context, billNum int64, accountNum int64, payPassword string) (msg string, err error)
}

type basicPaymentService struct {

}

func (b *basicPaymentService) Pay(ctx context.Context, billNum int64, accountNum int64, payPassword string) (msg string, err error) {
	// 1. 根据订单号查找订单金额
	// 2. 获取根据accountNum查找对应账户
	// 3. 验证密码和表中是否一致
	// 4. 检测账上余额是否大于订单金额
	// 5. 如果是则执行事务，并将订单的payed改为true
	account, err := models.GetAccount(accountNum, payPassword)
	if err != nil {
		return "GetAccount fail", err
	}
	bill, err := models.GetBill(billNum)
	if err != nil {
		return "GetBill fail", err
	}
	//parentSpan := opentracing.SpanFromContext(ctx)
	//childSpan := parentSpan.Tracer().StartSpan("mysql", opentracing.ChildOf(parentSpan.Context()))
	//defer childSpan.Finish()
	resp, err := TxPay(ctx, account, bill)
	if err != nil {
		return "fail", err
	}
	return resp, err
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
