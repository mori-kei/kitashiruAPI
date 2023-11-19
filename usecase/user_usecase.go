package usecase

import (
	"kitashiruAPI/model"
	"kitashiruAPI/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Signup(user model.User) (model.UserResponse, error)
	Login(user model.User) (model.UserResponse, string, error)
	GetUserByJwt(tokenString string) (model.UserResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) Signup(user model.User) (model.UserResponse, error) {
	// 平文のパスワードをハッシュ化
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user model.User) (model.UserResponse, string, error) {
	storedUser := model.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return model.UserResponse{}, "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return model.UserResponse{}, "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   storedUser.ID,
		"email":     storedUser.Email,
		"user_type": "user",
		"exp":       time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return model.UserResponse{}, "", err
	}
	resUser := model.UserResponse{
		ID:    storedUser.ID,
		Email: storedUser.Email,
	}
	return resUser, tokenString, nil
}

func (uu *userUsecase) GetUserByJwt(tokenString string) (model.UserResponse, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 適切な署名キーを返すためのコードをここに記述する
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return model.UserResponse{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := uint(claims["user_id"].(float64))
		response := model.UserResponse{
			ID:    id,
			Email: claims["email"].(string),
		}
		return response, nil
	} else {
		return model.UserResponse{}, err
	}
}
