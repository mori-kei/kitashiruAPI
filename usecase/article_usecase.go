package usecase

import (
	"kitashiruAPI/model"
	"kitashiruAPI/repository"
)

type IArticleUsecase interface {
	CreateArticle(article model.Article) (model.Article, error)
}

type articleUsecase struct {
	ar repository.IArticleRepository
}

func NewArticleUsecase(ar repository.IArticleRepository) IArticleUsecase {
	return &articleUsecase{ar}
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
