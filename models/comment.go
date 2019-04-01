package models

type Comment struct {
	Model
	Content   string `json: "content"`
	UserID    uint   `json: "user_id"`
	ArticleID uint   `json: "article_id"`
}
