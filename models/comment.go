package models

type Comment struct {
	Model
	Content string `json: "content"`
	User    User   `json: "user"`
}
