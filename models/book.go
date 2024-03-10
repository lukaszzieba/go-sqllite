package models

type Book struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Isbn   string `json:"isbn"`
}
