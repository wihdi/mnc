package usecase

import (
	"github.com/wihdi/mnc/domain"
	"github.com/wihdi/mnc/repository"
	"errors"
)

type TransferUsecase interface {
	GetAllTransfers() ([]domain.Transfer, error)
	TransferFunds(senderAccount, receiverAccount domain.Transfers, amount int) error
}

type transferUsecase struct {
	transferRepository repository.TransferRepository
}

var (
	ErrInvalidAmount = errors.New("invalid amount")
)

func NewTransferUsecase(transferRepository repository.TransferRepository) TransferUsecase {
	return &transferUsecase{
		transferRepository: transferRepository,
	}
}

func (u *transferUsecase) GetAllTransfers() ([]domain.Transfer, error) {
	transfers, err := u.transferRepository.GetAllTransfers()
	if err != nil {
		return nil, err
	}
	return transfers, nil
}

func (u *transferUsecase) TransferFunds(senderAccount, receiverAccount domain.Transfers, amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	senderUser, err := u.transferRepository.FindByAccountNumber(senderAccount.AccountNumber)
    if err != nil {
        return err
    }
    if senderUser == nil {
        return errors.New("sender account not registered")
    }

    // Cek apakah penerima ada dalam database
    receiverUser, err := u.transferRepository.FindByAccountNumber(receiverAccount.AccountNumber)
    if err != nil {
        return err
    }
    if receiverUser == nil {
        return errors.New("receiver account not registered")
    }

	return nil
}

