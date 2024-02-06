package usecase

import (
	"kitashiruAPI/model"
	"kitashiruAPI/repository"
)

type IFavoriteUsecase interface {
	ToggleFavorite(userId uint, articleJd uint) error
}

type favoriteUsecase struct {
	fr repository.IFavoriteRepository
}

func NewFavoriteUsecase(fr repository.IFavoriteRepository) IFavoriteUsecase {
	return &favoriteUsecase{fr}
}

func (fu *favoriteUsecase) ToggleFavorite(userId uint, articleId uint) error {
	favorite := model.Favorite{
		UserID:    uint(userId),
		ArticleID: uint(articleId),
	}
	err := fu.fr.ToggleFavorite(&favorite)
	if err != nil {
		return err
	}
	return nil

}
