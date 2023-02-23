package model

import "gorm.io/gorm"

// 标签
type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);comment:标签名" json:"name"`
}
