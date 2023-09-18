package usecase

import (
	"kitashiruAPI/model"
	"kitashiruAPI/repository"

	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface{
	Signup(user model.User)(model.UserResponse,error)
}

type userUsecase struct {
	ur repository.IUserRepository
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &userUsecase{ur}
}

func(uu *userUsecase) Signup(user model.User)(model.UserResponse,error) {
	// 平文のパスワードをハッシュ化
	hash,err := bcrypt.GenerateFromPassword([]byte(user.Password),10)
	if err != nil{
	return model.UserResponse{},err
	}
	newUser := model.User{Email: user.Email,Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil{
		return model.UserResponse{},err
	}
	resUser:= model.UserResponse{
		ID: newUser.ID,
		Email: newUser.Email,
	}
	return resUser, nil
}