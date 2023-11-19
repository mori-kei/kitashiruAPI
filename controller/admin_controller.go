package controller

import (
	"kitashiruAPI/model"
	"kitashiruAPI/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IAdminController interface {
	SignUp(c echo.Context) error
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
