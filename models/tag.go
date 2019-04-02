package models

type Tag struct {
	Model
	Name     string     `json: "name"`
	Articles []*Article `json: "articles" gorm:"many2many:article_tags"`
}

//func (t *Tag) Articles() []Article {
//	var articles []Article
//	if err := db.GetDB().Model(&t).Related(&articles, "Tags").Error; err != nil {
//		return nil
//	}
//	return articles
//}
