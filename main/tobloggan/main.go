package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/mdwhatcott/tobloggan/code/integration"
	"github.com/mdwhatcott/tobloggan/code/markdown"
)

func main() {
	started := time.Now()
	var (
		sourceDirectory string
		targetDirectory string
	)
	flags := flag.NewFlagSet("integration", flag.ExitOnError)
	flags.StringVar(&sourceDirectory, "source", "", "The directory containing blog source files (*.md).")
	flags.StringVar(&targetDirectory, "target", "", "The directory to output rendered blog files (*.html).")
	_ = flags.Parse(os.Args[1:])

	config := integration.Config{
		Logger:            log.New(os.Stderr, ">>> ", 0),
		MarkdownConverter: markdown.NewConverter(),
		FileSystemReader:  os.DirFS(sourceDirectory),
		FileSystemWriter:  FSWriter{},
		TargetDirectory:   targetDirectory,
	}
	ok := integration.GenerateBlog(config)
	if !ok {
		os.Exit(1)
	}
	config.Logger.Printf("integration finished in %s", time.Since(started))
}
