package models

type Tag struct {
	Model
	Name    string  `json: "name"`
	Article Article `json: "article" gorm:"many2many:article_tags"`
}
