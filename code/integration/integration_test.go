package integration

import (
	"bytes"
	"log"
	"os"
	"testing"
	"testing/fstest"

	"github.com/mdwhatcott/tobloggan/code/html"
	"github.com/mdwhatcott/tobloggan/code/markdown"
)

func Test(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in -short mode")
	}

	var logBuffer bytes.Buffer
	fs := make(FileSystem)
	// TODO: add source files
	config := Config{
		Logger:            log.New(&logBuffer, "[TEST] ", 0),
		MarkdownConverter: markdown.NewConverter(),
		FileSystemReader:  fstest.MapFS(fs),
		FileSystemWriter:  fs,
		TargetDirectory:   "output",
		ArticleTemplate:   html.ArticleTemplate,
		ListingTemplate:   html.ListingTemplate,
	}
	ok := GenerateBlog(config)
	if !ok {
		t.Error("failed to generate blog")
	}
	// TODO: assert content on fs
}

type FileSystem fstest.MapFS

func (this FileSystem) MkdirAll(path string, perm os.FileMode) error {
	this[path] = &fstest.MapFile{Mode: perm}
	return nil
}
func (this FileSystem) WriteFile(filename string, data []byte, perm os.FileMode) error {
	this[filename] = &fstest.MapFile{Data: data, Mode: perm}
	return nil
}
