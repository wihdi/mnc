package usecase

import (
	"github.com/wihdi/mnc/domain"
)

type HistoryUsecase struct {
	historyRepository domain.HistoryRepository
}

func NewHistoryUsecase(historyRepository domain.HistoryRepository) *HistoryUsecase {
	return &HistoryUsecase{
		historyRepository: historyRepository,
	}
}

func (u *HistoryUsecase) AddHistory(history *domain.History) error {
	return u.historyRepository.AddHistory(history)
}

func (u *HistoryUsecase) GetAllHistories() ([]*domain.History, error) {
	return u.historyRepository.GetAllHistories()
}
