package handler

import (
	"errors"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wihdi/mnc/domain"
	"github.com/wihdi/mnc/repository"
	"github.com/wihdi/mnc/usecase"
)

type TransferHandler struct {
	transferUsecase   usecase.TransferUsecase
	historyRepository domain.HistoryRepository
	userRepository    repository.TransferRepository
}

func NewTransferHandler(
	transferUsecase usecase.TransferUsecase,
	historyRepository domain.HistoryRepository,
	userRepository repository.TransferRepository,
) *TransferHandler {
	return &TransferHandler{
		transferUsecase:   transferUsecase,
		historyRepository: historyRepository,
		userRepository:    userRepository,
	}
}

func (h *TransferHandler) GetAllTransfers(c *gin.Context) {
	transfers, err := h.transferUsecase.GetAllTransfers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transfers)
}

func (h *TransferHandler) CreateTransfer(c *gin.Context) {
	var request struct {
		SenderAccount   domain.Transfers `json:"sender_account"`
		ReceiverAccount domain.Transfers `json:"receiver_account"`
		Amount          int               `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah pengirim ada dalam database
	senderUser, err := h.userRepository.FindByAccountNumber(request.SenderAccount.AccountNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if senderUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sender account not registered"})
		return
	}

	// Cek apakah penerima ada dalam database
	receiverUser, err := h.userRepository.FindByAccountNumber(request.ReceiverAccount.AccountNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if receiverUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Receiver account not registered"})
		return
	}

	err = h.transferUsecase.TransferFunds(request.SenderAccount, request.ReceiverAccount, request.Amount)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidAmount) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	history := &domain.History{
		Description: fmt.Sprintf("TF from %s to %s, Amount: %d", request.SenderAccount.AccountNumber, request.ReceiverAccount.AccountNumber, request.Amount),
	}
	err = h.historyRepository.AddHistory(history)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
}
