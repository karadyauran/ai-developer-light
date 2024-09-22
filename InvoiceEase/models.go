package models

type Invoice struct {
	ID         string  `json:"id"`
	ClientName string  `json:"client_name"`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
	DueDate    string  `json:"due_date"`
}