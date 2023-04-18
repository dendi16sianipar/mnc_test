package usecase

import (
	"mnc_test/models"
	"mnc_test/repositories"
)

type customerUsecase struct {
	customerRepo repositories.CustomerRepo
}

type CustomerUsecase interface {
	Create(newCustomer *models.Customer) (*models.Customer, error)
	Login(email string, password string) (string, error)
}

func NewCustomerUsecase(customerRepo repositories.CustomerRepo) CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
	}
}

func (r *customerUsecase) Create(newCustomer *models.Customer) (*models.Customer, error) {
	return r.customerRepo.Create(newCustomer)
}

func (r *customerUsecase) Login(email string, password string) (string, error) {
	return r.customerRepo.Login(email, password)
}
