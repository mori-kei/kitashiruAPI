package repository

import (
	"kitashiruAPI/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IArticleRepository interface {
	CreateArticle(article *model.Article) error
	GetAllArticles() ([]model.Article, error)
	GetArticle(article *model.Article, articleId uint) error
	UpdateArticle(article *model.Article, articleId uint) error
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

func (ar *articleRepository) GetArticle(article *model.Article, articleId uint) error {
	if err := ar.db.Model(article).Where("id", articleId).First(article).Error; err != nil {
		return err
	}
	return nil
}
func (ar *articleRepository) UpdateArticle(article *model.Article, articleId uint) error {
	result := ar.db.Model(article).Clauses(clause.Returning{}).Where("id", articleId).Updates(article)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
