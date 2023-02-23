package model

import "gorm.io/gorm"

// 分类
type Category struct {
	gorm.Model
	Name   string `gorm:"type:varchar(20);not null;comment:分类名称" json:"name"`
	Status int8   `gorm:"type:tinyint;comment:标签状态(0:删除 1:正常)" json:"status"`
	// Articles []Article `gorm:"foreignKey:CategoryId"` // 重写外键
}
