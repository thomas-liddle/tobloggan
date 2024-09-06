package main

import (
	"flag"
	"io"
	"log"
	"os"
	"sync/atomic"

	"github.com/mdwhatcott/pipelines"
	"github.com/mdwhatcott/tobloggan/code/contracts"
	stations2 "github.com/mdwhatcott/tobloggan/code/stations"
)

func main() {
	var sourceDirectory string
	flags := flag.NewFlagSet("tobloggan", flag.ExitOnError)
	flags.StringVar(&sourceDirectory, "src", "", "The directory containing blog source files (*.md).")
	_ = flags.Parse(os.Args[1:])
	ok := GenerateBlog(sourceDirectory, os.Stderr)
	if !ok {
		os.Exit(1)
	}
}

func GenerateBlog(sourceDirectory string, stderr io.Writer) bool {
	var (
		fs       = os.DirFS(sourceDirectory)
		logger   = log.New(stderr, "", log.Lmicroseconds)
		failed   = new(atomic.Bool)
		input    = make(chan any, 1)
		pipeline = pipelines.New(input,
			pipelines.Options.Logger(logger),
			pipelines.Options.StationSingleton(stations2.NewSourceScanner(fs)),
			pipelines.Options.StationSingleton(stations2.NewSourceReader(fs)),
			pipelines.Options.StationSingleton(stations2.NewArticleParser(stations2.NewGoldmarkMarkdownConverter())),
			// TODO: render articles
			// TODO: render home page
			pipelines.Options.StationSingleton(stations2.NewReporter(logger, failed)),
		)
	)
	input <- contracts.SourceDirectory(".")
	close(input)
	pipeline.Listen()
	return !failed.Load()
}
