package tobloggan

import (
	"os"
	"sync/atomic"

	"github.com/mdwhatcott/pipelines"
	"github.com/mdwhatcott/tobloggan/code/contracts"
	"github.com/mdwhatcott/tobloggan/code/markdown"
	"github.com/mdwhatcott/tobloggan/code/stations"
)

type Config struct {
	SourceDirectory string
	Logger          contracts.Logger
}

func GenerateBlog(config Config) bool {
	var (
		fs       = os.DirFS(config.SourceDirectory)
		failed   = new(atomic.Bool)
		input    = make(chan any, 1)
		pipeline = pipelines.New(input,
			pipelines.Options.Logger(config.Logger),
			pipelines.Options.StationSingleton(stations.NewSourceScanner(fs)),
			pipelines.Options.StationSingleton(stations.NewSourceReader(fs)),
			pipelines.Options.StationSingleton(stations.NewArticleParser(markdown.NewConverter())),
			// TODO: render articles
			// TODO: render home page
			pipelines.Options.StationSingleton(NewReporter(config.Logger, failed)),
		)
	)
	input <- contracts.SourceDirectory(".")
	close(input)
	pipeline.Listen()
	return !failed.Load()
}
