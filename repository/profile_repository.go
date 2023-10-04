package repository

import (
	"fmt"
	"kitashiruAPI/model"

	"gorm.io/gorm"
)

type IProfileRepository interface {
	CreateProfile(profile *model.Profile) error
	DeleteProfile(userId uint) error
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

func (pr *profileRepository) DeleteProfile(userId uint) error {
	result := pr.db.Where("user_id=?", userId).Delete(&model.Profile{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
