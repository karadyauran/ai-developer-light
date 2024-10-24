package service

import (
	"errors"
	"eco_track_api/internal/model"
	"gorm.io/gorm"
)

type CarbonService struct {
	db *gorm.DB
}

func NewCarbonService(db *gorm.DB) *CarbonService {
	return &CarbonService{db: db}
}

type CarbonCalculationRequest struct {
	UserID   int     `json:"user_id"`
	Activity string  `json:"activity"`
	Amount   float64 `json:"amount"`
	Unit     string  `json:"unit"`
}

type CarbonCalculationResponse struct {
	TotalEmissions float64 `json:"total_emissions"`
}

func (s *CarbonService) Calculate(input CarbonCalculationRequest) (CarbonCalculationResponse, error) {
	var conversionFactor float64
	switch input.Activity {
	case "driving":
		conversionFactor = 2.31
	case "flying":
		conversionFactor = 0.21
	default:
		return CarbonCalculationResponse{}, errors.New("unsupported activity")
	}

	totalEmissions := input.Amount * conversionFactor

	record := model.CarbonRecord{
		UserID:   input.UserID,
		Activity: input.Activity,
		Amount:   input.Amount,
		Unit:     input.Unit,
	}

	if err := s.db.Create(&record).Error; err != nil {
		return CarbonCalculationResponse{}, err
	}

	return CarbonCalculationResponse{TotalEmissions: totalEmissions}, nil
}