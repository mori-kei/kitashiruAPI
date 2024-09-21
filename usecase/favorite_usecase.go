package usecase

import (
	"kitashiruAPI/model"
	"kitashiruAPI/repository"
)

type IFavoriteUsecase interface {
	ToggleFavorite(userId uint, articleJd uint) error
	GetAllFavoriteArticles(userId uint) ([]model.ArticleWithLikedStatus, error)
}

type favoriteUsecase struct {
	fr repository.IFavoriteRepository
	ar repository.IArticleRepository
}

func NewFavoriteUsecase(fr repository.IFavoriteRepository, ar repository.IArticleRepository) IFavoriteUsecase {
	return &favoriteUsecase{fr, ar}
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
func (fu *favoriteUsecase) GetAllFavoriteArticles(userId uint) ([]model.ArticleWithLikedStatus, error) {

	favorites, err := fu.fr.GetFavoritesByUserID(userId)
	if err != nil {
		return nil, err
	}
	// 記事IDのスライスを作成
	var articleIDs []uint
	for _, favorite := range favorites {
		articleIDs = append(articleIDs, favorite.ArticleID)
	}
	favoriteArticles, err := fu.ar.GetArticlesByArticleId(articleIDs)
	if err != nil {
		return nil, err
	}
	var favoriteArticlesWithStatus []model.ArticleWithLikedStatus
	for _, favoriteArticle := range favoriteArticles {
		favoriteArticleWithStatus := model.ArticleWithLikedStatus{
			Article: favoriteArticle,
			Liked:   true,
		}

		favoriteArticlesWithStatus = append(favoriteArticlesWithStatus, favoriteArticleWithStatus)

	}
	return favoriteArticlesWithStatus, nil
}
