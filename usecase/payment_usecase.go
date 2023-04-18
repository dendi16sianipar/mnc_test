package usecase

import (
	"mnc_test/models"
	"mnc_test/repositories"
)

type transactionUsecase struct {
	paymentRepo repositories.PaymentRepo
}

type PaymentUsecase interface {
	Create(transaction *models.Payment) (*models.Payment, error)
}

func NewPaymentUsecase(paymentRepo repositories.PaymentRepo) PaymentUsecase {
	return &transactionUsecase{
		paymentRepo: paymentRepo,
	}
}

func (r *transactionUsecase) Create(transaction *models.Payment) (*models.Payment, error) {
	return r.paymentRepo.Create(transaction)
}
