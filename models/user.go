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
