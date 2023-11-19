package repository

import (
	"kitashiruAPI/model"

	"gorm.io/gorm"
)




type IAdminRepository interface {
	CreateAdmin(admin *model.Admin )error
}

type adminRepository struct{
	db *gorm.DB
}

func NewAdminRepository (db *gorm.DB) IAdminRepository{
		return &adminRepository{db}
}

func (ar *adminRepository) CreateAdmin(admin *model.Admin)error{
	if err:= ar.db.Create(admin).Error; err!= nil {
		return err
	}
	return nil
}