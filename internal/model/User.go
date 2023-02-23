package model

import "gorm.io/gorm"

const ()

// 用户
type User struct {
	gorm.Model
	Username  string `gorm:"type:varchar(30);comment:用户名" json:"username"`
	Gender    string `gorm:"type:varchar(10);comment:性别" json:"gender"`
	Password  string `gorm:"type:varchar(20);comment:密码" json:"password"`
	Email     string `gorm:"type:varchar(25);comment:邮箱" json:"email"`
	Avatar    string `gorm:"type:varchar(200);comment:头像" json:"avatar"`
	Info      string `gorm:"type:varchar(200);comment:个人简介" json:"info"`
	IsDisable bool   `gorm:"type:bool;default:false;comment:是否被禁(0:否 1:是)" json:"is_disable"`
}
