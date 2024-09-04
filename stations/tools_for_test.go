package stations

import "errors"

var boink = errors.New("boink")

type StationFixture struct {
	outputs []any
}

func (this *StationFixture) Output(v any) {
	this.outputs = append(this.outputs, v)
}
