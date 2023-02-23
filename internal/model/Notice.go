package model

import "gorm.io/gorm"

// 通知
type Notice struct {
	gorm.Model
	SenderID      int    `gorm:"type:int;comment:发送者id" json:"sender_id"`
	ReceiverID    int    `gorm:"type:int;comment:接收者id" json:"receiver_id"`
	Content       string `gorm:"type:varchar(300);comment:通知内容" json:"content"`
	ReplyParentID int    `gorm:"type:int;comment:回复的父评论id" json:"reply_parent_id"`
	ReplyChildID  int    `gorm:"type:int;comment:回复的子评论id" json:"reply_child_id"`
}
