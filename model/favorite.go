package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserID    uint `gorm:"index"` // users テーブルとの外部キー
	ArticleID uint `gorm:"index"` // articles テーブルとの外部キー

}
