package tobloggan

import (
	"io/fs"
	"sync/atomic"

	"github.com/mdwhatcott/pipelines"
	"github.com/mdwhatcott/tobloggan/code/contracts"
	"github.com/mdwhatcott/tobloggan/code/stations"
)

type Config struct {
	Logger            contracts.Logger
	MarkdownConverter stations.MarkdownConverter
	FileSystemReader  fs.FS
	FileSystemWriter  stations.FileSystemWriter
	TargetDirectory   string
}

func GenerateBlog(config Config) bool {
	var (
		failure = new(atomic.Bool)
		input   = make(chan any, 1)

		scanner  = stations.NewSourceScanner(config.FileSystemReader)
		reader   = stations.NewSourceReader(config.FileSystemReader)
		parser   = stations.NewArticleParser(config.MarkdownConverter)
		articles = stations.NewArticleWriter(config.TargetDirectory, config.FileSystemWriter)
		listing  = stations.NewListingWriter(config.TargetDirectory, config.FileSystemWriter)
		reporter = stations.NewReporter(config.Logger, failure)

		pipeline = pipelines.New(input,
			pipelines.Options.Logger(config.Logger),
			pipelines.Options.StationSingleton(scanner),
			pipelines.Options.StationSingleton(reader), pipelines.Options.FanOut(5),
			pipelines.Options.StationSingleton(parser),
			pipelines.Options.StationSingleton(articles), pipelines.Options.FanOut(5),
			pipelines.Options.StationSingleton(listing),
			pipelines.Options.StationSingleton(reporter),
		)
	)
	input <- contracts.SourceDirectory(".")
	close(input)
	pipeline.Listen()
	return !failure.Load()
}
