package stations

import (
	"errors"
	"time"

	"github.com/mdwhatcott/pipelines"
)

var boink = errors.New("boink")

type StationFixture struct {
	station pipelines.Station // TODO: migrate all fixtures to use this field
	outputs []any
}

func (this *StationFixture) do(input any) { // TODO: migrate all fixtures to use this method?
	this.station.Do(input, this.Output)
}
func (this *StationFixture) Output(v any) { // TODO: rename to output (lowercase), or inline completely
	this.outputs = append(this.outputs, v)
}

////////////////////////////////////////////////////////////

func date(YYYY_MM_DD string) time.Time {
	t, err := time.Parse("2006-01-02", YYYY_MM_DD)
	if err != nil {
		panic(err)
	}
	return t
}
