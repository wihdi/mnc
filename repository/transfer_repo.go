package repository

import (
	"encoding/json"
	"os"
"log"
	"github.com/wihdi/mnc/domain"
)

type TransferRepository interface {
	GetAllTransfers() ([]domain.Transfer, error)
	FindByAccountNumber(account_number string) (*domain.User, error)
	// SaveTransfer(transfer domain.Transfer) error
}

type transferRepository struct {
	filePath string
}

func NewTransferRepository(filePath string) *transferRepository {
	return &transferRepository{filePath: filePath}
}

func (r *transferRepository) GetAllTransfers() ([]domain.Transfer, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var transfers []domain.Transfer
	err = json.NewDecoder(file).Decode(&transfers)
	if err != nil {
		return nil, err
	}

	return transfers, nil
}
func (r *transferRepository) FindByAccountNumber(account_number string) (*domain.User, error) {
    users, err := r.readUsers()
    if err != nil {
        log.Println("Failed to read users from JSON file:", err)
        return nil, err
    }

    for _, user := range users {
        if user.AccountNumber == account_number {
            return &user, nil
        }
    }

    return nil, nil
}


func (r *transferRepository) readUsers() ([]domain.User, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []domain.User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

