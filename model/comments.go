package model

type Comment struct {
	Id      int    `json:"id"`
	PostId  int    `json:"post_id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}
