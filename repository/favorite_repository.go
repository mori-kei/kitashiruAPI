package repository

import (
	"kitashiruAPI/model"

	"gorm.io/gorm"
)

type IFavoriteRepository interface {
	ToggleFavorite(favorite *model.Favorite) error
	GetFavoritesByUserID(userID uint) ([]model.Favorite, error)
	IsFavoriteExists(favorite *model.Favorite) (bool, error)
}

type favoriteRepository struct {
	db *gorm.DB
}

func NewFavoriteRepository(db *gorm.DB) IFavoriteRepository {
	return &favoriteRepository{db}
}

func (fr *favoriteRepository) ToggleFavorite(favorite *model.Favorite) error {
	result := fr.db.Where(&favorite).First(favorite)

	if result.Error == nil {
		fr.db.Delete(favorite)
	}
	fr.db.Create(favorite)
	return nil
}

func (fr *favoriteRepository) GetFavoritesByUserID(userID uint) ([]model.Favorite, error) {
	var favorites []model.Favorite
	if err := fr.db.Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
		return nil, err
	}
	return favorites, nil
}

func (fr *favoriteRepository) IsFavoriteExists(favorite *model.Favorite) (bool, error) {
	result := fr.db.Where(favorite).First(&model.Favorite{})
	if result.Error == gorm.ErrRecordNotFound {
		return false, nil
	} else if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
