package domain

type User struct {
	ID             int    `json:"id"`
	Username string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	AccountNumber string `json:"account_number"`
	Balance       int    `json:"balance"`
	Amount          int    `json:"amount"`

}

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Transfers struct{
	Username string `json:"username"`
	AccountNumber string `json:"account_number"`
}
