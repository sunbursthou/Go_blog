package model

import "gorm.io/gorm"

const (
	PUBLIC = iota + 1
	SECRET
	DRAFT
)

// 文章
type Article struct {
	gorm.Model
	CategoryID int    `gorm:"type:int;not null;comment:分类id" json:"category_id"`
	UserID     int    `gorm:"type:int;not null;comment:作者ID" json:"user_id"`
	Title      string `gorm:"type:varchar(50);not null;comment:文章标题" json:"title"`
	Desc       string `gorm:"type:varchar(150);comment:文章简述" json:"desc"`
	Img        string `gorm:"type:varchar(200);comment:文章配图" json:"img"`
	Views      int    `gorm:"type:int;not null;default:0;comment:浏览量" json:"views"`
	Like       int    `gorm:"type:int;not null; default:0;comment:赞" json:"like"`
	// Tag        string `gorm:"type:varchar(20);comment:标签" json:"tag"`
	Content  string `gorm:"type:longtext;comment:文章内容" json:"content"`
	Status   int8   `gorm:"type:tinyint;comment:状态(1:公开 2:私密 3:草稿)" json:"status"`
	Istop    bool   `gorm:"type:bool;default:false;comment:是否置顶(1:置顶 0:不置顶)" json:"is_top"`
	IsDelete bool   `gorm:"type:bool;default:false;comment:是否软删除(1:是 0:否)" json:"is_delete"`
}

// Tag与Article关联表
type ArticleToTag struct {
	ArticleID int `gorm:"type:int;comment:文章id;index:article_to_tag" json:"article_id"`
	TagID     int `gorm:"type:int;comment:标签id;index:article_to_tag" json:"tag_id"`
}
