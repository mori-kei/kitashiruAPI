package usecase

import (
	"kitashiruAPI/model"
	"kitashiruAPI/repository"
	"math/rand"
	"sort"
	"time"
)

type IArticleUsecase interface {
	CreateArticle(article model.Article) (model.Article, error)
	GetMatchArticles(userId uint) ([]model.ArticleWithLikedStatus, error)
	GetAllArticles() ([]model.Article, error)
	GetAllArticlesRandom(userId uint) ([]model.ArticleWithLikedStatus, error)
	GetArticle(userId uint, articleId uint) (model.ArticleWithLikedStatus, error)
	UpdateArticle(article model.Article, articleId uint) (model.Article, error)
}

type articleUsecase struct {
	ar repository.IArticleRepository
	pr repository.IProfileRepository
	fr repository.IFavoriteRepository
}

func NewArticleUsecase(ar repository.IArticleRepository, pr repository.IProfileRepository, fr repository.IFavoriteRepository) IArticleUsecase {
	return &articleUsecase{ar, pr, fr}
}

func (au *articleUsecase) CreateArticle(article model.Article) (model.Article, error) {
	if err := au.ar.CreateArticle(&article); err != nil {
		return model.Article{}, err
	}
	resArticle := model.Article{
		ID:              article.ID,
		Url:             article.Url,
		OverView:        article.OverView,
		Message:         article.Message,
		Appeal:          article.Appeal,
		CapitalAmount:   article.CapitalAmount,
		EarningAmount:   article.EarningAmount,
		CompanySize:     article.CompanySize,
		Address:         article.Address,
		IsPublished:     article.IsPublished,
		FamilyPoint:     article.FamilyPoint,
		InnovationPoint: article.InnovationPoint,
		MarketPoint:     article.MarketPoint,
		BurePoint:       article.BurePoint,
		CreatedAt:       article.CreatedAt,
		UpdatedAt:       article.UpdatedAt,
	}
	return resArticle, nil
}

func (au *articleUsecase) GetMatchArticles(userId uint) ([]model.ArticleWithLikedStatus, error) {
	profile := model.Profile{}
	favorites, err := au.fr.GetFavoritesByUserID(userId)
	if err := au.pr.GetProfileByUserId(&profile, userId); err != nil {
		return nil, err
	}
	articles, err := au.ar.GetAllPublicArticles()
	if err != nil {
		return nil, err
	}
	// 計算した絶対値の合計を格納するためのマップ
	articleDiff := make(map[uint]int)

	// ユーザープロフィールとArticleのポイントの差分を計算し、合計を算出
	for _, article := range articles {
		diff := Abs(int(profile.Beuraucracy) - article.BurePoint)
		diff += Abs(int(profile.Family) - article.FamilyPoint)
		diff += Abs(int(profile.Innovation) - article.InnovationPoint)
		diff += Abs(int(profile.Market) - article.MarketPoint)

		articleDiff[article.ID] = diff
	}

	// 計算結果を基にArticleをソートする
	sort.Slice(articles, func(i, j int) bool {
		return articleDiff[articles[i].ID] < articleDiff[articles[j].ID]
	})

	var articlesWithStatus []model.ArticleWithLikedStatus
	for _, article := range articles {
		// 記事ごとにいいねの状態を判定
		var liked bool
		for _, favorite := range favorites {
			if favorite.ArticleID == article.ID {
				liked = true
				break
			}
		}

		// ArticleWithLikedStatus 構造体に詰める
		articleWithStatus := model.ArticleWithLikedStatus{
			Article: article,
			Liked:   liked,
		}

		articlesWithStatus = append(articlesWithStatus, articleWithStatus)
	}
	if err != nil {
		return nil, err
	}
	shuffleArticles(articles)
	return articlesWithStatus, nil
}
func (au *articleUsecase) GetAllArticles() ([]model.Article, error) {
	articles, err := au.ar.GetAllArticles()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (au *articleUsecase) GetAllArticlesRandom(userId uint) ([]model.ArticleWithLikedStatus, error) {
	articles, err := au.ar.GetAllPublicArticles()
	favorites, err := au.fr.GetFavoritesByUserID(userId)
	var articlesWithStatus []model.ArticleWithLikedStatus
	for _, article := range articles {
		// 記事ごとにいいねの状態を判定
		var liked bool
		for _, favorite := range favorites {
			if favorite.ArticleID == article.ID {
				liked = true
				break
			}
		}

		// ArticleWithLikedStatus 構造体に詰める
		articleWithStatus := model.ArticleWithLikedStatus{
			Article: article,
			Liked:   liked,
		}

		articlesWithStatus = append(articlesWithStatus, articleWithStatus)
	}
	if err != nil {
		return nil, err
	}
	shuffleArticles(articles)
	return articlesWithStatus, nil
}

func (au *articleUsecase) GetArticle(userId uint, articleId uint) (model.ArticleWithLikedStatus, error) {
	article := model.Article{}

	if err := au.ar.GetArticle(&article, articleId); err != nil {
		return model.ArticleWithLikedStatus{}, err
	}

	favorite := model.Favorite{
		UserID:    userId,
		ArticleID: articleId,
	}
	isLiked, err := au.fr.IsFavoriteExists(&favorite)
	if err != nil {
		return model.ArticleWithLikedStatus{}, err
	}
	response := model.ArticleWithLikedStatus{
		Article: article,
		Liked:   isLiked,
	}
	return response, nil
}
func (au *articleUsecase) UpdateArticle(article model.Article, articleId uint) (model.Article, error) {
	if err := au.ar.UpdateArticle(&article, articleId); err != nil {
		return model.Article{}, err
	}
	return article, nil
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func shuffleArticles(articles []model.Article) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(articles), func(i, j int) {
		articles[i], articles[j] = articles[j], articles[i]
	})
}
