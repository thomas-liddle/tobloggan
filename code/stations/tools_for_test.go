package stations

import (
	"errors"
	"time"

	"github.com/mdwhatcott/pipelines"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

var boink = errors.New("boink")

type StationFixture struct {
	*gunit.Fixture
	station pipelines.Station
	outputs []any
}

func (this *StationFixture) TestUnhandledTypeEmitted() {
	this.do("wrong-type")
	this.So(this.outputs, should.Equal, []any{"wrong-type"})
}
func (this *StationFixture) do(input any) {
	this.station.Do(input, this.output)
}
func (this *StationFixture) finalize() {
	this.station.(pipelines.Finalizer).Finalize(this.output)
}
func (this *StationFixture) output(v any) {
	this.outputs = append(this.outputs, v)
}
func (this *StationFixture) assertOutputs(expected ...any) {
	this.So(this.outputs, should.Equal, expected)
}

////////////////////////////////////////////////////////////

func date(YYYY_MM_DD string) time.Time {
	t, err := time.Parse("2006-01-02", YYYY_MM_DD)
	if err != nil {
		panic(err)
	}
	return t
}
