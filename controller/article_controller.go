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
	GetAllArticles(c echo.Context) error
	GetAllPublicArticleRandom(c echo.Context) error
	GetArticle(c echo.Context) error
	UpdateArticle(c echo.Context) error
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

func (ac *articleController) handleUserToken(c echo.Context) (uint, error) {
	user := c.Get("user")
	if user == nil {
		return 0, c.JSON(http.StatusBadRequest, "User token not found")
	}

	claims, ok := user.(*jwt.Token).Claims.(jwt.MapClaims)
	if !ok {
		return 0, c.JSON(http.StatusBadRequest, "Invalid user token")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, c.JSON(http.StatusBadRequest, "Invalid user ID in token")
	}

	return uint(userIDFloat), nil
}

func (ac *articleController) GetMatchArticles(c echo.Context) error {
	userID, err := ac.handleUserToken(c)
	if err != nil {
		return err
	}

	articles, err := ac.au.GetMatchArticles(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, articles)
}

func (ac *articleController) GetAllArticles(c echo.Context) error {
	articles, err := ac.au.GetAllArticles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, articles)
}

func (ac *articleController) GetAllPublicArticleRandom(c echo.Context) error {
	userID, err := ac.handleUserToken(c)
	if err != nil {
		return err
	}

	articles, err := ac.au.GetAllArticlesRandom(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, articles)
}

func (ac *articleController) GetArticle(c echo.Context) error {
	id := c.Param("articleId")
	articleID, _ := strconv.Atoi(id)
	userID, err := ac.handleUserToken(c)
	articleRes, err := ac.au.GetArticle(userID,uint(articleID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, articleRes)
}

func (ac *articleController) UpdateArticle(c echo.Context) error {
	article := model.Article{}
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := c.Param("articleId")
	articleID, _ := strconv.Atoi(id)
	articleRes, err := ac.au.UpdateArticle(article, uint(articleID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, articleRes)
}
