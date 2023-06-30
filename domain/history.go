package domain

type History struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

type HistoryRepository interface {
	AddHistory(history *History) error
	GetAllHistories() ([]*History, error)
}

