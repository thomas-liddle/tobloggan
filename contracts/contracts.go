package contracts

import (
	"errors"
	"time"
)

type SourceDirectory string
type SourceFilePath string
type SourceFile []byte
type Article struct {
	Slug  string    `json:"slug"`
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
	Body  string    `json:"-"`
}

var ErrMalformedSource = errors.New("malformed source")
