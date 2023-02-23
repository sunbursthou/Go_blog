package model

import "gorm.io/gorm"

type Announcement struct {
	gorm.Model
	Content string `gorm:"type:varchar(200);comment:网站公告" json:"content"`
}
