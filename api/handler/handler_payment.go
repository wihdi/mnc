package handler

import (
	"net/http"
"strconv"
"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wihdi/mnc/domain"
	"github.com/wihdi/mnc/usecase"
)

type PaymentHandler struct {
	paymentUsecase  usecase.PaymentUsecase
	historyRepository domain.HistoryRepository
}

func NewPaymentHandler(paymentUsecase usecase.PaymentUsecase, historyRepository domain.HistoryRepository) *PaymentHandler {
	return &PaymentHandler{
		paymentUsecase:  paymentUsecase,
		historyRepository: historyRepository,
	}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var request struct {
		AccountNumber string `json:"account_number"`
		Amount        string `json:"amount"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	amount, err := strconv.Atoi(request.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	err = h.paymentUsecase.ProcessPayment(request.AccountNumber, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	history := &domain.History{
		Description: fmt.Sprintf("Payment from %s , Amount: %d", request.AccountNumber, amount),
	}
	err = h.historyRepository.AddHistory(history)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully"})
}


