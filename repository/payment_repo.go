package repository

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/wihdi/mnc/domain"
)

type CustomerRepository interface {
	GetCustomerByAccountNumber(accountNumber string) (*domain.User, error)
	UpdateCustomerBalance(customer *domain.User) error
}

type customerRepository struct {
	filePath string
}

func NewCustomerRepository(filePath string) *customerRepository {
	return &customerRepository{
		filePath: filePath,
	}
}

func (r *customerRepository) GetCustomerByAccountNumber(accountNumber string) (*domain.User, error) {
	customers, err := r.readCustomers()
	if err != nil {
		return nil, err
	}

	for _, customer := range customers {
		if strings.EqualFold(customer.AccountNumber, accountNumber) {
			return &customer, nil
		}
	}

	return nil, nil
}

func (r *customerRepository) UpdateCustomerBalance(user *domain.User) error {
	customers, err := r.readCustomers()
	if err != nil {
		return err
	}

	for i, c := range customers {
		if c.ID == user.ID {
			customers[i] = *user
			break
		}
	}

	err = r.writeCustomers(customers)
	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) readCustomers() ([]domain.User, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var customers []domain.User
	err = json.NewDecoder(file).Decode(&customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}


func (r *customerRepository) writeCustomers(customers []domain.User) error {
	file, err := os.Create(r.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(&customers)
	if err != nil {
		return err
	}

	return nil
}
