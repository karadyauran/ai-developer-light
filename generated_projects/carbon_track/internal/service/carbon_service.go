package service

import (
	"errors"
	"generated_projects/carbon_track/internal/model"
	"generated_projects/carbon_track/internal/repository"
)

type CarbonInput struct {
	UserID      int     `json:"user_id"`
	Activity    string  `json:"activity"`
	CO2Emitted  float64 `json:"co2_emitted"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
}

type CarbonService struct {
	repo *repository.CarbonRepository
}

func NewCarbonService(repo *repository.CarbonRepository) *CarbonService {
	return &CarbonService{repo: repo}
}

func (cs *CarbonService) CalculateFootprint(input CarbonInput) (model.CarbonFootprint, error) {
	if input.CO2Emitted <= 0 {
		return model.CarbonFootprint{}, errors.New("invalid CO2 emission value")
	}
	footprint := model.CarbonFootprint{
		UserID:      input.UserID,
		Activity:    input.Activity,
		CO2Emitted:  input.CO2Emitted,
		Date:        input.Date,
		Description: input.Description,
	}
	err := cs.repo.SaveFootprint(&footprint)
	if err != nil {
		return model.CarbonFootprint{}, err
	}
	return footprint, nil
}

func (cs *CarbonService) GetFootprintRecords() ([]model.CarbonFootprint, error) {
	return cs.repo.FetchFootprints()
}