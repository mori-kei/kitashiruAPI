package controller

import (
	"kitashiruAPI/model"
	"kitashiruAPI/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IArticleController interface {
	CreateArticle(c echo.Context) error
}

type articleController struct {
	au usecase.IArticleUsecase
}

func NewArticleController(au usecase.IArticleUsecase) IArticleController {
	return &articleController{au}
}

func (ac *articleController) CreateArticle(c echo.Context) error {
	article := model.Article{}

	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	articleRes, err := ac.au.CreateArticle(article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, articleRes)
}
