package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wihdi/mnc/domain"
	"github.com/wihdi/mnc/usecase"
)

type HistoryHandler struct {
	historyUsecase usecase.HistoryUsecase
}

func NewHistoryHandler(historyUsecase usecase.HistoryUsecase) *HistoryHandler {
	return &HistoryHandler{
		historyUsecase: historyUsecase,
	}
}

func (h *HistoryHandler) LogActivity(c *gin.Context) {
	action := c.Param("action")

	history := &domain.History{
		Description: "Performed action: " + action,
	}

	err := h.historyUsecase.AddHistory(history)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Activity logged successfully"})
}

func (h *HistoryHandler) GetAllActivities(c *gin.Context) {
	histories, err := h.historyUsecase.GetAllHistories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, histories)
}
