package models

type Article struct {
	Model
	Title    string    `json: "title"`
	Content  string    `json: "content"`
	UserID   uint      `json: "user_id"`
	Comments []Comment `json: "comments"`
	Tags     []Tag     `json:"tags" gorm:"many2many:article_tags"`
}
