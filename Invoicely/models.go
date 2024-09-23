package models

type Invoice struct {
	ID          int     `json:"id"`
	ClientName  string  `json:"client_name"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
}

type Report struct {
	TotalIncome float64 `json:"total_income"`
	Pending     int     `json:"pending"`
	Paid        int     `json:"paid"`
}