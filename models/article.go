package models

type Article struct {
	Model
	Title    string    `json: "title"`
	Content  string    `json: "content"`
	User     User      `json: "user"`
	Comments []Comment `json: "comments"`
	Tags     []Tag     `json:"tags"`
}
