package usecase

import (
	"kitashiruAPI/model"
	"kitashiruAPI/repository"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type IAuthUsecase interface {
	GetAuthByJwt(tokenString string) (model.AuthResponse, error)
}

type authUsecase struct {
	ar repository.IAuthRepository
}

func NewAuthUsecase(ar repository.IAuthRepository) IAuthUsecase {
	return &authUsecase{ar}
}
func (au *authUsecase) GetAuthByJwt(tokenString string) (model.AuthResponse, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 適切な署名キーを返すためのコードをここに記述する
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return model.AuthResponse{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := uint(claims["user_id"].(float64))
		user_type := string(claims["user_type"].(string))
		response := model.AuthResponse{
			ID:        id,
			Email:     claims["email"].(string),
			User_type: user_type,
		}
		return response, nil
	} else {
		return model.AuthResponse{}, err
	}
}
