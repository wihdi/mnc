package domain

type Customer struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
	Balance       int    `json:"balance"`
}
