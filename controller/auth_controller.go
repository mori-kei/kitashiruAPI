package controller

import (
	"kitashiruAPI/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IAuthController interface {
	GetAuthByJwt(c echo.Context) error
}

type authController struct {
	au usecase.IAuthUsecase
}

func NewAuthController(au usecase.IAuthUsecase) IAuthController {
	return &authController{au}
}

func (ac *authController) GetAuthByJwt(c echo.Context) error {
	cookie, err := c.Cookie("token")
	if err != nil {
		return err
	}
	tokenString := cookie.Value

	authResp, err := ac.au.GetAuthByJwt(tokenString)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, authResp)
}
