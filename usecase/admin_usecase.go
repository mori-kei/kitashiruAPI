package usecase

import (
	"kitashiruAPI/model"
	"kitashiruAPI/repository"

	"golang.org/x/crypto/bcrypt"
)

type IAdminUsecase interface {
	SignUp(admin model.Admin) (model.AdminResponse, error)
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
