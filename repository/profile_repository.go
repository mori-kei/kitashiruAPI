package repository

import (
	"kitashiruAPI/model"

	"gorm.io/gorm"
)

type IProfileRepository interface {
	CreateProfile(profile *model.Profile) error
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) IProfileRepository {
	return &profileRepository{db}
}

func (pr *profileRepository) CreateProfile(profile *model.Profile) error {
	if err := pr.db.Create(profile).Error; err != nil {
		return err
	}
	return nil
}
