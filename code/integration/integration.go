package integration

import (
	"io/fs"
	"sync/atomic"

	"github.com/mdwhatcott/pipelines"
	"github.com/mdwhatcott/tobloggan/code/contracts"
	"github.com/mdwhatcott/tobloggan/code/stations"
)

type Config struct {
	Logger           contracts.Logger
	Markdown         stations.Markdown
	FileSystemReader fs.FS
	FileSystemWriter stations.FileSystemWriter
	TargetDirectory  string
	ArticleTemplate  string
	ListingTemplate  string
	BaseURL          string
}

func GenerateBlog(config Config) bool {
	var (
		failure = new(atomic.Bool)
		input   = make(chan any, 1)

		scanner  = stations.NewSourceScanner(config.FileSystemReader)
		reader   = stations.NewSourceReader(config.FileSystemReader)
		parser   = stations.NewArticleParser()
		markdown = stations.NewMarkdownConverter(config.Markdown)
		//renderer = stations.NewArticleRenderer(config.ArticleTemplate)
		articles = stations.NewArticleWriter(config.TargetDirectory, config.FileSystemWriter, config.ArticleTemplate)
		listing  = stations.NewListingWriter(config.TargetDirectory, config.FileSystemWriter, config.ListingTemplate)
		reporter = stations.NewReporter(config.Logger, failure)

		// TODO: incorporate baseurl writer

		pipeline = pipelines.New(input,
			pipelines.Options.Logger(config.Logger),
			pipelines.Options.StationSingleton(scanner),
			pipelines.Options.StationSingleton(reader), pipelines.Options.FanOut(5),
			pipelines.Options.StationSingleton(parser),
			pipelines.Options.StationSingleton(markdown),
			//pipelines.Options.StationSingleton(renderer),
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
