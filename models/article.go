package models

type Article struct {
	Model
	Title    string    `json: "title"`
	Content  string    `json: "content"`
	State    bool      `json: "state"`
	UserID   uint      `json: "user_id" gorm: "index"`
	User     User      `json: "user"`
	Comments []Comment `json: "comments"`
	Tags     []*Tag    `json:"tags" gorm:"many2many:article_tags"`
}

//@TODO add gorm callback transform time into timestamp
