package models

type Comment struct {
	Model
	Content   string  `json: "content"`
	UserID    uint    `json: "user_id"`
	User      User    `json: "user"`
	ArticleID uint    `json: "article_id"`
	Article   Article `json: "article"`
}

//func (c *Comment) User() *User {
//	var user User
//	db.GetDB().Where("id = ?", c.UserID).Find(&user)
//	return &user
//}
//
//func (c *Comment) Article() *Article {
//	var article Article
//	db.GetDB().Where("id = ?", c.ArticleID).Find(&article)
//	return &article
//}
