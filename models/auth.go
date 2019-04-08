package models

import "github.com/ShawnRong/bento/db"

type Auth struct {
	ID       int    `gorm:"primary_key" json: "id"`
	Username string `json: "username"`
	Password string `json: "password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.GetDB().Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}

	return false
}
