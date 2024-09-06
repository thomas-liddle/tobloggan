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
		failure  = new(atomic.Bool)
		input    = make(chan any, 1)
		pipeline = pipelines.New(input,
			pipelines.Options.Logger(config.Logger),
			pipelines.Options.StationSingleton(stations.NewSourceScanner(config.FileSystemReader)),
			pipelines.Options.StationSingleton(stations.NewSourceReader(config.FileSystemReader)),
			pipelines.Options.StationSingleton(stations.NewArticleParser(config.MarkdownConverter)),
			pipelines.Options.StationSingleton(stations.NewArticleWriter(config.TargetDirectory, config.FileSystemWriter)),
			pipelines.Options.StationSingleton(stations.NewListingWriter(config.TargetDirectory, config.FileSystemWriter)),
			pipelines.Options.StationSingleton(NewReporter(config.Logger, failure)),
		)
	)
	input <- contracts.SourceDirectory(".")
	close(input)
	pipeline.Listen()
	return !failure.Load()
}
