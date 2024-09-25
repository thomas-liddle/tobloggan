package main

import (
	"log"
	"os"
	"time"

	"github.com/mdwhatcott/tobloggan/code/html"
	"github.com/mdwhatcott/tobloggan/code/integration"
	"github.com/mdwhatcott/tobloggan/code/markdown"
)

func main() {
	started := time.Now()
	log.SetFlags(0)
	log.SetPrefix(">>> ")
	cli := parseFlags(os.Args[1:])
	ok := integration.GenerateBlog(integration.Config{
		Logger:           log.Default(),
		Markdown:         markdown.NewConverter(),
		FileSystemReader: os.DirFS(cli.sourceDirectory),
		FileSystemWriter: FSWriter{},
		TargetDirectory:  cli.targetDirectory,
		ArticleTemplate:  html.ArticleTemplate,
		ListingTemplate:  html.ListingTemplate,
		BaseURL:          cli.baseURL,
	})
	log.Printf("finished in %s", time.Since(started))
	if !ok {
		os.Exit(1)
	}
}
