package stations

import (
	"reflect"
	"sync/atomic"

	"github.com/mdwhatcott/tobloggan/code/contracts"
)

type Reporter struct {
	logger contracts.Logger
	failed *atomic.Bool
}

func NewReporter(logger contracts.Logger, failed *atomic.Bool) *Reporter {
	return &Reporter{logger: logger, failed: failed}
}

func (this *Reporter) Do(input any, _ func(any)) {
	switch input := input.(type) {
	case error:
		this.failed.Store(true)
		this.logger.Printf("err: %v", input)
	case contracts.Article: // TODO: contracts.Page
		this.logger.Printf("article: %s", input.Title)
	default:
		this.logger.Printf("unexpected type: %s", reflect.TypeOf(input))
	}
}
