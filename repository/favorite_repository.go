package repository

import (
	"kitashiruAPI/model"

	"gorm.io/gorm"
)

type IFavoriteRepository interface {
	ToggleFavorite(favorite *model.Favorite) error
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
