package models

type Article struct {
	Model
	Title    string    `json: "title"`
	Content  string    `json: "content"`
	UserID   uint      `json: "user_id"`
	User     User      `json: "user"`
	Comments []Comment `json: "comments"`
	Tags     []*Tag    `json:"tags" gorm:"many2many:article_tags"`
}

//func (a *Article) User() *User {
//	var user User
//	db.GetDB().Where("id = ?", a.UserID).Find(&user)
//	return &user
//}
//
//func (a *Article) Comments() []Comment {
//	var comments []Comment
//	if err := db.GetDB().Where("user_id = ?", a.ID).Find(&comments).Error; err != nil {
//		return nil
//	}
//	return comments
//}
//
//func (a *Article) Tags() []Tag {
//	var tags []Tag
//	if err := db.GetDB().Model(&a).Related(&tags, "articles").Error; err != nil {
//		return nil
//	}
//	return tags
//}
