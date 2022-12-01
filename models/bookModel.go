package models

type Book struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	ISBN     string `json:"isbn"`
	AuthorID int    `json:"authorId"`
}
