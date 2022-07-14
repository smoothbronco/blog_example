package model

type Article struct {
	ID      int    `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateInput struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdateInput struct {
	ID      int    `json:"id"`
	Author  string `json:"author"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
