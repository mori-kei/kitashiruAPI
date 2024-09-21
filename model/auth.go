package model

type AuthResponse struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Email     string `json:"email" gorm:"unique"`
	User_type string `json:"user_type"`
}


