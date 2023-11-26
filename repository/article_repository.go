package repository

import (
	"kitashiruAPI/model"

	"gorm.io/gorm"
)

type IArticleRepository interface {
	CreateArticle(article *model.Article) error
	GetAllArticles() ([]model.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) IArticleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) CreateArticle(article *model.Article) error {
	if err := ar.db.Create(article).Error; err != nil {
		return err
	}
	return nil
}

func (ar *articleRepository) GetAllArticles() ([]model.Article, error) {
	var articles []model.Article
	if err := ar.db.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
