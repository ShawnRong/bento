package models

type Tag struct {
	Model
	Name     string     `json: "name"`
	Articles []*Article `json: "articles" gorm:"many2many:article_tags"`
}
