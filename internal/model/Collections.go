package model

import "gorm.io/gorm"

// 收藏
type Collections struct {
	gorm.Model
	ArticleID int `gorm:"type:int;comment:收藏文章id" json:"article_id"`
	Collector int `gorm:"type:int;comment:收藏者" json:"collector"`
}
