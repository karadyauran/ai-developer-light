package model

import "time"

type CarbonRecord struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"user_id"`
	Activity  string    `json:"activity"`
	Amount    float64   `json:"amount"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
}

func (CarbonRecord) TableName() string {
	return "carbon_records"
}

type CarbonCalculationInput struct {
	Activity string  `json:"activity"`
	Amount   float64 `json:"amount"`
	Unit     string  `json:"unit"`
}