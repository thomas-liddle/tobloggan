package contracts

import (
	"errors"
	"time"
)

type SourceDirectory string
type SourceFilePath string
type Article struct {
	Slug  string    `json:"slug"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
	Body  string    `json:"-"`
}

type Markdown interface {
	Convert(content string) (string, error)
}

var ErrMalformedSource = errors.New("malformed source")

const ScanCompleted = "** scan completed **"
