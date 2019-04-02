package models

type User struct {
	Model
	Email    string    `gorm:"type:varchar(100);unique_index" json:"email" binding:"required,email"`
	Name     string    `json:"name"`
	Role     string    `gorm:"size:255" json:"role"`
	Active   bool      `gorm:"default: false" json:"active"`
	Password string    `json:"password"`
	Articles []Article `json: "articles"`
	Comments []Comment `json: "comments"`
}

// @todo custom validate

//func (u *User) Articles() []Article {
//	var articles []Article
//	if err := db.GetDB().Where("user_id = ?", u.ID).Find(&articles).Error; err != nil {
//		return nil
//	}
//	return articles
//}
//
//func (u *User) Comments() []Comment {
//	var comments []Comment
//	if err := db.GetDB().Where("user_id = ?", u.ID).Find(&comments).Error; err != nil {
//		return nil
//	}
//	return comments
//}
