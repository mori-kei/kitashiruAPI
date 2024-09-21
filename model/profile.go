package model

import (
	"time"

	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"not null"`
	Beuraucracy float64 `json:"beuraucracy" gorm:"not null"`
	Family float64 `json:"family" gorm:"not null"`
	Innovation float64 `json:"Innovation" gorm:"not null"`
	Market float64 `json:"market" gorm:"not null"`
}

type ProfileResponse struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	Beuraucracy float64 `json:"beuraucracy" gorm:"not null"`
	Family float64 `json:"family" gorm:"not null"`
	Innovation float64 `json:"Innovation" gorm:"not null"`
	Market float64 `json:"market" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}