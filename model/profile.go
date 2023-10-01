package model

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	ID     uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"not null"`
	Beuraucracy float64 `json:"beuraucracy" gorm:"not null"`
	Family float64 `json:"family" gorm:"not null"`
	Innovation float64 `json:"Innovation" gorm:"not null"`
	Market float64 `json:"market" gorm:"not null"`
}

