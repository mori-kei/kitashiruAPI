package controller

import (
	"kitashiruAPI/usecase"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IFavoriteController interface {
	ToggleFavorite(c echo.Context) error
}

type favoriteContoroller struct {
	fu usecase.IFavoriteUsecase
}

func NewFavoriteController(fu usecase.IFavoriteUsecase) IFavoriteController {
	return &favoriteContoroller{fu}
}

func (fc *favoriteContoroller) ToggleFavorite(c echo.Context) error {
	id := c.Param("articleId")
	articleId, _ := strconv.Atoi(id)
	cookie, err := c.Cookie("token")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Token cookie not found"})
	}

	// tokenをパースしてclaimsを取得
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid token"})
	}

	// claimsからuser_idを取得
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get claims"})
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get user_id from claims"})
	}

	userId := uint(userIDFloat)

	fc.fu.ToggleFavorite(uint(userId), uint(articleId))
	return c.JSON(http.StatusOK, nil)
}
