package model

import (
	"gorm.io/gorm"
)

// 留言
type Message struct {
	gorm.Model
	Nickname string `gorm:"type:varchar(15);not null;comment:昵称" json:"nick_name"`
	Content  string `gorm:"type:varchar(200);not null;comment:留言内容" json:"content"`
}
