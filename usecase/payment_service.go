package usecase

import (
	"errors"

	//"github.com/wihdi/mnc/domain"
	"github.com/wihdi/mnc/repository"
)

type PaymentUsecase interface {
	ProcessPayment(accountNumber string, amount int) error
}

type paymentUsecase struct {
	customerRepository repository.CustomerRepository
}

func NewPaymentUsecase(customerRepository repository.CustomerRepository) PaymentUsecase {
	return &paymentUsecase{
		customerRepository: customerRepository,
	}
}

func (u *paymentUsecase) ProcessPayment(accountNumber string, amount int) error {
	customer, err := u.customerRepository.GetCustomerByAccountNumber(accountNumber)
	if err != nil {
		return err
	}

	if customer == nil {
		return errors.New("customer not found")
	}

	if customer.Balance < amount {
		return errors.New("insufficient balance")
	}

	customer.Balance -= amount

	err = u.customerRepository.UpdateCustomerBalance(customer)
	if err != nil {
		return err
	}

	return nil
}
