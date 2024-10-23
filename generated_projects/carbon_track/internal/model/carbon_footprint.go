package model

type CarbonFootprint struct {
	ID          int     `json:"id" gorm:"primaryKey"`
	UserID      int     `json:"user_id"`
	Activity    string  `json:"activity"`
	CO2Emitted  float64 `json:"co2_emitted"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
}

func (CarbonFootprint) TableName() string {
	return "carbon_footprints"
}
