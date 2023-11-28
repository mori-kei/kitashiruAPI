package controller

import (
	"kitashiruAPI/model"
	"kitashiruAPI/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IArticleController interface {
	CreateArticle(c echo.Context) error
	GetMatchArticles(c echo.Context) error
	GetAllArticleRandom(c echo.Context) error
	GetArticle(c echo.Context) error
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

func (ac *articleController) GetMatchArticles(c echo.Context) error {
	user := c.Get("user")
	if user == nil {
		return c.JSON(http.StatusBadRequest, "User token not found")
	}

	claims, ok := user.(*jwt.Token).Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusBadRequest, "Invalid user token")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return c.JSON(http.StatusBadRequest, "Invalid user ID in token")
	}

	articles, err := ac.au.GetMatchArticles(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, articles)
}

func (ac *articleController) GetAllArticleRandom(c echo.Context) error {
	articles, err := ac.au.GetAllArticlesRandom()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, articles)
}
func (ac *articleController) GetArticle(c echo.Context) error {
	id := c.Param("articleId")
	articleId, _ := strconv.Atoi(id)
	articleRes, err := ac.au.GetArticle(uint(articleId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, articleRes)
}
