package model

type Siteinfo struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(20);comment:站长姓名" json:"name"`
	Desc   string `gorm:"type:varchar(200);comment:站长简介" json:"desc"`
	Wechat string `gorm:"type:varchar(100);comment:站长微信" json:"wechat"`
	Weibo  string `gorm:"type:varchar(200);comment:站长微博" json:"weibo"`
	Bili   string `gorm:"type:varchar(200);comment:站长哔哩哔哩" json:"bili"`
	Avatar string `gorm:"type:varchar(200);comment:站长头像" json:"avatar"`
}
