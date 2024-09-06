package main

import (
	"flag"
	"log"
	"os"

	"github.com/mdwhatcott/tobloggan/code/markdown"
	"github.com/mdwhatcott/tobloggan/code/tobloggan"
)

func main() {
	var (
		sourceDirectory string
		targetDirectory string
	)
	flags := flag.NewFlagSet("tobloggan", flag.ExitOnError)
	flags.StringVar(&sourceDirectory, "source", "", "The directory containing blog source files (*.md).")
	flags.StringVar(&targetDirectory, "target", "", "The directory to output rendered blog files (*.html).")
	_ = flags.Parse(os.Args[1:])

	config := tobloggan.Config{
		Logger:            log.New(os.Stderr, "", log.Ltime),
		MarkdownConverter: markdown.NewConverter(),
		FileSystemReader:  os.DirFS(sourceDirectory),
		FileSystemWriter:  FSWriter{},
		TargetDirectory:   targetDirectory,
	}
	ok := tobloggan.GenerateBlog(config)
	if !ok {
		os.Exit(1)
	}
}
