package contracts

import "time"

// SourceDirectory is just a path representing a directory.
type SourceDirectory string

// SourceFilePath is just a path representing a content file
type SourceFilePath string

// SourceFile is the raw contents of a content file.
type SourceFile []byte

// Article represents the parsed data from a SourceFile.
type Article struct {
	Draft bool      `json:"draft"` // If true, don't publish this article.
	Slug  string    `json:"slug"`  // The URL path of the article.
	Title string    `json:"title"` // The <h1> of the article.
	Date  time.Time `json:"date"`  // Affects sorting of articles.
	Body  string    `json:"-"`     // Markdown, will become HTML.
}

// Page represents a rendered article, ready to write to disk.
type Page struct {
	Path    string
	Content string
}
