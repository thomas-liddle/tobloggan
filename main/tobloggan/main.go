package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/mdwhatcott/tobloggan/code/html"
	"github.com/mdwhatcott/tobloggan/code/integration"
	"github.com/mdwhatcott/tobloggan/code/markdown"
)

type Config struct {
	sourceDirectory string
	targetDirectory string
	baseURL         string
}

func main() {
	started := time.Now()
	log.SetFlags(0)
	log.SetPrefix(">>> ")
	cli := parseCLI(os.Args[1:])
	app := integration.Config{
		Logger:            log.Default(),
		MarkdownConverter: markdown.NewConverter(),
		FileSystemReader:  os.DirFS(cli.sourceDirectory),
		FileSystemWriter:  FSWriter{},
		TargetDirectory:   cli.targetDirectory,
		ArticleTemplate:   html.ArticleTemplate,
		ListingTemplate:   html.ListingTemplate,
		BaseURL:           cli.baseURL,
	}
	log.Printf("source: %s", cli.sourceDirectory)
	log.Printf("target: %s", cli.targetDirectory)
	log.Printf("base-url: %s", cli.baseURL)

	ok := integration.GenerateBlog(app)
	if !ok {
		os.Exit(1)
	}
	log.Printf("finished in %s", time.Since(started))
}

func parseCLI(args []string) (result Config) {
	flags := flag.NewFlagSet("integration", flag.ExitOnError)
	flags.StringVar(&result.sourceDirectory, "source", "content", "The directory containing blog source files (*.md).")
	flags.StringVar(&result.targetDirectory, "target", "generated", "The directory to output rendered blog files (*.html).")
	flags.StringVar(&result.baseURL, "base-url", "", `The base-url of all internal hyperlinks. (default "file://<target>")`)
	_ = flags.Parse(args)
	if result.baseURL == "" {
		result.baseURL = "file://" + result.targetDirectory
	}
	return result
}
