package model

import "time"

type Article struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Name            string    `json:"name" `
	Url             string    `json:"url"`
	OverView        string    `json:"overview"`
	Message         string    `json:"message"`
	Appeal          string    `json:"appeal"`
	CapitalAmount   int       `json:"capital_amount"`
	EarningAmount   int       `json:"earning_amount"`
	CompanySize     int       `json:"company_size"`
	Address         string    `json:"address"`
	IsPublished     bool      `json:"is_published" gorm:"default:true"`
	FamilyPoint     int       `json:"family_point"`
	InnovationPoint int       `json:"innovation_point"`
	MarketPoint     int       `json:"market_point"`
	BurePoint       int       `json:"bure_point"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
