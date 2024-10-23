package service

import (
	"gorm.io/gorm"
	"generated_projects/carbon_track/internal/model"
)

type DatabaseService struct {
	db *gorm.DB
}

func NewDatabaseService(db *gorm.DB) *DatabaseService {
	return &DatabaseService{db: db}
}

func (ds *DatabaseService) AutoMigrate() error {
	return ds.db.AutoMigrate(&model.CarbonFootprint{})
}

func (ds *DatabaseService) GetDB() *gorm.DB {
	return ds.db
}