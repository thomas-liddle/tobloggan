package stations

import (
	"tobloggan/code/contracts"
)

type DraftRemoval struct{}

func NewDraftRemoval() DraftRemoval {
	return DraftRemoval{}
}

func (this DraftRemoval) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Article:
		if input.Draft {
			return
		}
		output(input)
	default:
		output(input)
	}
}
