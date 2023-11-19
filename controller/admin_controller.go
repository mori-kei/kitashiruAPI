package controller

import (
	"kitashiruAPI/model"
	"kitashiruAPI/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IAdminController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
}

type adminController struct {
	au usecase.IAdminUsecase
}

func NewAdminController(au usecase.IAdminUsecase) IAdminController {
	return &adminController{au}
}

func (ac *adminController) SignUp(c echo.Context) error {
	admin := model.Admin{}
	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	adminRes, err := ac.au.SignUp(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, adminRes)
}
func (ac *adminController) LogIn(c echo.Context) error {
	admin := model.Admin{}
	if err := c.Bind(&admin); err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}
	adminRes, tokenString, err := ac.au.Login(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	//↓postmanで動作確認する時はコメントアウトを行い一旦falseにする
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, adminRes)
}
