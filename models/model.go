package models

import (
	"time"

	"github.com/ShawnRong/bento/db"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json: id`
	CreatedAt time.Time  `json: created_at`
	UpdatedAt time.Time  `json: updated_at`
	DeletedAt *time.Time `sql:"index" json: deleted_at`
}

func AutoMigrate() {
	db.GetDB().AutoMigrate(&User{})
}
