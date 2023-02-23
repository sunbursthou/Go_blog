package dao

import "github.com/sunbursthou/Go_blog/internal/model"

// 创建articletotag关联
func CreateArticleToTag(articleID, tagID int) error {
	data := &model.ArticleToTag{
		ArticleID: articleID,
		TagID:     tagID,
	}
	err := DB.Create(data).Error
	return err
}
