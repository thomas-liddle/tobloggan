package stations

import (
	"time"

	"tobloggan/code/contracts"
)

type FutureRemoval struct {
	started time.Time
}

func NewFutureRemoval(started time.Time) *FutureRemoval {
	return &FutureRemoval{started: started}
}

func (this *FutureRemoval) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Article:
		if !input.Date.After(this.started) {
			output(input)
		}
	default:
		output(input)
	}
}
