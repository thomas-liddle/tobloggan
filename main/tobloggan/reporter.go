package main

import (
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
	switch input.(type) {
	case error:
		this.failed.Store(true)
	}
	this.logger.Printf("%v", input)
}
