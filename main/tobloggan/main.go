package main

import (
	"log"
	"os"
	"time"

	"tobloggan/code/html"
	"tobloggan/code/integration"
	"tobloggan/code/markdown"
)

func main() {
	logger := log.New(os.Stderr, ">>> ", 0)
	cli := parseFlags(os.Args[1:])
	ok := integration.GenerateBlog(integration.Config{
		Clock:             time.Now,
		Logger:            logger,
		MarkdownConverter: markdown.NewConverter(),
		FileSystemReader:  os.DirFS(cli.sourceDirectory),
		FileSystemWriter:  FSWriter{},
		TargetDirectory:   cli.targetDirectory,
		ArticleTemplate:   html.ArticleTemplate,
		ListingTemplate:   html.ListingTemplate,
		BaseURL:           cli.baseURL,
	})
	if !ok {
		logger.Fatal("WARNING: The blog failed to generate successfully!")
	}
}
