package contracts

import "time"

type SourceDirectory string
type SourceFilePath string
type SourceFile []byte
type Article struct {
	Slug  string    `json:"slug"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
	Body  string    `json:"-"`
}
type Page struct {
	Path    string
	Content string
}
