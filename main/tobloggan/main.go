package main

import (
	"log"
	"os"
	"time"

	"tobloggan/code/html"
	"tobloggan/code/integration"
	"tobloggan/code/markdown"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix(">>> ")
}

func main() {
	cli := parseFlags(os.Args[1:])
	ok := integration.GenerateBlog(integration.Config{
		Clock:             time.Now,
		Logger:            log.Default(),
		MarkdownConverter: markdown.NewConverter(),
		FileSystemReader:  os.DirFS(cli.sourceDirectory),
		FileSystemWriter:  FSWriter{},
		TargetDirectory:   cli.targetDirectory,
		ArticleTemplate:   html.ArticleTemplate,
		ListingTemplate:   html.ListingTemplate,
		BaseURL:           cli.baseURL,
	})
	if !ok {
		log.Fatal("WARNING: The blog failed to generate successfully!")
	}
}
