package model

import "time"

type Article struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Url           string    `json:"url"`
	OverView      string    `json:"overview"`
	Message       string    `json:"message"`
	Appeal        string    `json:"appeal"`
	CapitalAmount int       `json:"capital_amount"`
	EarningAmount int       `json:"earning_amount"`
	CompanySize   int       `json:"company_size"`
	Address       string    `json:"address"`
	IsPublished   bool      `json:"is_published" gorm:"default:true"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
