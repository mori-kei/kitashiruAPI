package controller

import (
	"kitashiruAPI/model"
	"kitashiruAPI/usecase"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IProfileController interface {
	CreateProfile(c echo.Context) error
	DeleteProfile(c echo.Context) error
	GetProfileByUserId(c echo.Context)error
}

type profileController struct {
	pu usecase.IProfileUsecase
}

func NewProfileController(pu usecase.IProfileUsecase) IProfileController {
	return &profileController{pu}
}

func (pc *profileController) CreateProfile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	profile := model.Profile{}
	if err := c.Bind(&profile); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	profile.UserID = uint(userId.(float64))
	profileRes, err := pc.pu.CreateProfile(profile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, profileRes)
}

func (pc *profileController) DeleteProfile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	err := pc.pu.DeleteProfile(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
func(pc *profileController)GetProfileByUserId(c echo.Context)error{
	user:= c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	profileRes,err := pc.pu.GetProfileByUserId(uint(userId.(float64)))
	if err!= nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK,profileRes)
}