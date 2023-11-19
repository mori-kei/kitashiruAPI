package repository

import (
	"kitashiruAPI/model"

	"gorm.io/gorm"
)

type IAdminRepository interface {
	CreateAdmin(admin *model.Admin) error
	GetUserByEmail(admin *model.Admin, email string) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) IAdminRepository {
	return &adminRepository{db}
}

func (ar *adminRepository) CreateAdmin(admin *model.Admin) error {
	if err := ar.db.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func (ar *adminRepository) GetUserByEmail(admin *model.Admin, email string) error {
	if err := ar.db.Where("email=?", email).First(admin).Error; err != nil {
		return err
	}
	return nil
}
