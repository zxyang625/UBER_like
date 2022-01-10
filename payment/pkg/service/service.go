package service

import (
	"context"
	"fmt"
)

// PaymentService describes the service.
type PaymentService interface {
	Pay(ctx context.Context, billNum int64, accountNum int64, payPassword string) (msg string, err error)
}

type basicPaymentService struct {

}

func (b *basicPaymentService) Pay(ctx context.Context, billNum int64, accountNum int64, payPassword string) (msg string, err error) {
	return fmt.Sprintf("%d_%d_%s", billNum, accountNum, payPassword), nil
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
