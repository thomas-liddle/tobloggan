package main

import (
	"log"
	"os"
	"time"

	"github.com/mdwhatcott/tobloggan/code/html"
	"github.com/mdwhatcott/tobloggan/code/integration"
	"github.com/mdwhatcott/tobloggan/code/markdown"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix(">>> ")
}

func main() {
	cli := parseFlags(os.Args[1:])
	ok := integration.GenerateBlog(integration.Config{
		Clock:            time.Now,
		Logger:           log.Default(),
		Markdown:         markdown.NewConverter(),
		FileSystemReader: os.DirFS(cli.sourceDirectory),
		FileSystemWriter: FSWriter{},
		TargetDirectory:  cli.targetDirectory,
		ArticleTemplate:  html.ArticleTemplate,
		ListingTemplate:  html.ListingTemplate,
		BaseURL:          cli.baseURL,
	})
	if !ok {
		log.Fatal("WARNING: The blog failed to generate successfully!")
	}
}
