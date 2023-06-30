package domain

type Transfer struct {
	
	SenderAccount   string`json:"sender_account"`
	ReceiverAccount string `json:"receiver_account"`
	Amount          int   `json:"amount"`
}
