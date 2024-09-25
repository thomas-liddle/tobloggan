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
		listing  = stations.NewListingRenderer(config.ListingTemplate)
		renderer = stations.NewArticleRenderer(config.ArticleTemplate)
		baseURL  = stations.NewBaseURLRewriter(config.BaseURL)
		writer   = stations.NewPageWriter(config.TargetDirectory, config.FileSystemWriter)
		reporter = stations.NewReporter(config.Logger, failure)

		pipeline = pipelines.New(input,
			pipelines.Options.Logger(config.Logger),
			pipelines.Options.StationSingleton(scanner),
			pipelines.Options.StationSingleton(reader), pipelines.Options.FanOut(5),
			pipelines.Options.StationSingleton(parser),
			pipelines.Options.StationSingleton(markdown),
			pipelines.Options.StationSingleton(listing),
			pipelines.Options.StationSingleton(renderer),
			pipelines.Options.StationSingleton(baseURL),
			pipelines.Options.StationSingleton(writer), pipelines.Options.FanOut(5),
			pipelines.Options.StationSingleton(reporter),
		)
	)
	input <- contracts.SourceDirectory(".")
	close(input)
	pipeline.Listen()
	return !failure.Load()
}
