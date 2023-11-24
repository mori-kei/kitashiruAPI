package repository

import (
	"kitashiruAPI/model"

	"gorm.io/gorm"
)

type IArticleRepository interface {
	CreateArticle(article *model.Article) error
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
