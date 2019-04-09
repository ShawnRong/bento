package models

type Comment struct {
	Model
	Content   string  `json: "content"`
	UserID    uint    `json: "user_id" gorm: "index"`
	User      User    `json: "user"`
	ArticleID uint    `json: "article_id" gorm: "index"`
	Article   Article `json: "article"`
}
