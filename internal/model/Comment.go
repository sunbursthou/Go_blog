package model

import "gorm.io/gorm"

// 评论
type Comment struct {
	gorm.Model
	UserID  int    `gorm:"type:int;not null;comment:评论用户id" json:"user_id"`
	Content string `gorm:"type:varchar(300);not null;comment:评论内容" json:"content"`
	Status  int8   `gorm:"type:tinyint;default:1;comment:评论状态(1:正常 2:删除 3:审核中)" json:"status"`
	Like    int    `gorm:"type:int;default:0;comment:赞" json:"like"`
	Dislike int    `gorm:"type:int;default:0;comment:踩" json:"dislike"`
	// Istop   bool
}

// 一级评论
type ArticleComment struct {
	Comment
	BelongToArticleID int  `gorm:"type:int;comment:所属文章id" json:"belong_to_article_id"`
	Istop             bool `gorm:"type:bool;default:false;comment:是否置顶(0:否 1:是)" json:"is_top"`
}

// 子评论和回复
type ChildComment struct {
	Comment
	BelongToParentID int `gorm:"type:int;comment:所属父评论ID" json:"belog_to_parent_id"`
	BelongToReplyID  int `gorm:"type:int;comment:所属回复ID" json:"belong_to_reply_id"`
}
