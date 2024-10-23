package repository

import (
	"generated_projects/carbon_track/internal/model"
	"gorm.io/gorm"
)

type CarbonRepository struct {
	db *gorm.DB
}

func NewCarbonRepository(db *gorm.DB) *CarbonRepository {
	return &CarbonRepository{db: db}
}

func (r *CarbonRepository) SaveFootprint(footprint *model.CarbonFootprint) error {
	return r.db.Create(footprint).Error
}

func (r *CarbonRepository) FetchFootprints() ([]model.CarbonFootprint, error) {
	var footprints []model.CarbonFootprint
	err := r.db.Find(&footprints).Error
	return footprints, err
}
