package contracts

import "time"

type (
	SourceDirectory string
	SourceFilePath  string
	SourceFile      []byte
	Article         struct {
		Slug  string    `json:"slug"`
		Title string    `json:"title"`
		Date  time.Time `json:"date"`
		Body  string    `json:"-"`
	}
	Page struct {
		Path    string
		Content string
	}
)
