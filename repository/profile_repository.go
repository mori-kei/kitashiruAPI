package repository

import (
	"fmt"
	"kitashiruAPI/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IProfileRepository interface {
	CreateProfile(profile *model.Profile) error
	UpdateProfile(profile *model.Profile, userId uint) error
	DeleteProfile(userId uint) error
	GetProfileByUserId(profile *model.Profile, userId uint) error
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

func (pr *profileRepository) UpdateProfile(profile *model.Profile, userId uint) error {
	result := pr.db.Model(profile).Clauses(clause.Returning{}).Where("user_id=?", userId).Updates(model.Profile{
		Beuraucracy: profile.Beuraucracy,
		Family:      profile.Family,
		Innovation:  profile.Innovation,
		Market:      profile.Market,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
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

func (pr *profileRepository) GetProfileByUserId(profile *model.Profile, userId uint) error {
	if err := pr.db.Model(profile).Where("user_id", userId).Last(profile).Error; err != nil {
		return err
	}
	return nil
}
