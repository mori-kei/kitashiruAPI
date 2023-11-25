package usecase

import (
	"kitashiruAPI/model"
	"kitashiruAPI/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IAdminUsecase interface {
	SignUp(admin model.Admin) (model.AdminResponse, error)
	Login(admin model.Admin) (model.AuthResponse, string, error)
}

type adminUsecase struct {
	ar repository.IAdminRepository
}

func NewAdminUsecase(ar repository.IAdminRepository) IAdminUsecase {
	return &adminUsecase{ar}
}

func (au *adminUsecase) SignUp(admin model.Admin) (model.AdminResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return model.AdminResponse{}, err
	}
	newAdmin := model.Admin{Email: admin.Email, Password: string(hash)}
	if err := au.ar.CreateAdmin(&newAdmin); err != nil {
		return model.AdminResponse{}, err
	}
	resAdmin := model.AdminResponse{
		ID:    newAdmin.ID,
		Email: newAdmin.Email,
	}

	return resAdmin, nil
}

func (au *adminUsecase) Login(admin model.Admin) (model.AuthResponse, string, error) {
	storedAdmin := model.Admin{}
	if err := au.ar.GetUserByEmail(&storedAdmin, admin.Email); err != nil {
		return model.AuthResponse{}, "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(
		storedAdmin.Password), []byte(admin.Password))
	if err != nil {
		return model.AuthResponse{}, "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   storedAdmin.ID,
		"email":     storedAdmin.Email,
		"user_type": "admin",
		"exp":       time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return model.AuthResponse{}, "", err
	}
	resAdmin := model.AuthResponse{
		ID:        storedAdmin.ID,
		Email:     storedAdmin.Email,
		User_type: "admin",
	}
	return resAdmin, tokenString, nil
}
