package main

import (
	"flag"
	"log"
)

type Config struct {
	sourceDirectory string
	targetDirectory string
	baseURL         string
}

func parseFlags(args []string) (result Config) {
	flags := flag.NewFlagSet("integration", flag.ExitOnError)
	flags.StringVar(&result.sourceDirectory, "source", "content", "The directory containing blog source files (*.md).")
	flags.StringVar(&result.targetDirectory, "target", "docs", "The directory to output rendered blog files (*.html).")
	flags.StringVar(&result.baseURL, "base-url", "http://localhost:8000/", `The base-url of all internal hyperlinks.`)
	_ = flags.Parse(args)
	log.Printf("source: %s", result.sourceDirectory)
	log.Printf("target: %s", result.targetDirectory)
	log.Printf("base-url: %s", result.baseURL)
	return result
}
