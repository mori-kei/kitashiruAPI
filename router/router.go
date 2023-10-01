package router

import (
	"kitashiruAPI/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.IUserController) *echo.Echo{
	e := echo.New()
	e.POST("/signup",uc.SignUp)
	e.POST("/login",uc.LogIn)
	return e
}