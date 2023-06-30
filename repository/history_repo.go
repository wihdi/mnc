package repository

import (
	"encoding/json"
	"io"
	"os"

	"github.com/wihdi/mnc/domain"
)

type HistoryRepository struct {
	filePath string
}

func NewHistoryRepository(filePath string) *HistoryRepository {
	return &HistoryRepository{
		filePath: filePath,
	}
}

func (r *HistoryRepository) AddHistory(history *domain.History) error {
	histories, err := r.getAllHistoriesFromFile()
	if err != nil {
		return err
	}

	history.ID = len(histories) + 1
	histories = append(histories, history)

	err = r.saveHistoriesToFile(histories)
	if err != nil {
		return err
	}

	return nil
}

func (r *HistoryRepository) GetAllHistories() ([]*domain.History, error) {
	return r.getAllHistoriesFromFile()
}

func (r *HistoryRepository) getAllHistoriesFromFile() ([]*domain.History, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		// Jika file tidak ditemukan, return slice kosong
		if os.IsNotExist(err) {
			return []*domain.History{}, nil
		}
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var histories []*domain.History
	err = json.Unmarshal(data, &histories)
	if err != nil {
		return nil, err
	}

	return histories, nil
}

func (r *HistoryRepository) saveHistoriesToFile(histories []*domain.History) error {
	data, err := json.Marshal(histories)
	if err != nil {
		return err
	}

	err = os.WriteFile(r.filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
