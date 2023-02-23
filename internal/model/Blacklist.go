package model

import (
	"gorm.io/gorm"
)

// 黑名单
type BlackList struct {
	gorm.Model
	Adder   int `gorm:"type:int;comment:添加者" json:"adder"`
	BeAdded int `gorm:"type:int;comment:被添加者" json:"be_added"`
}
