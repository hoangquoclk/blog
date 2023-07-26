package model

type Post struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	CategoryId string `json:"categoryId"`
	UserId     string `json:"userId"`
	Content    string `json:"content"`
}
