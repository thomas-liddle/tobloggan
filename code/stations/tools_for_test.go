package stations

import (
	"errors"
	"time"
)

var boink = errors.New("boink")

type StationFixture struct {
	outputs []any
}

func (this *StationFixture) Output(v any) {
	this.outputs = append(this.outputs, v)
}

func date(YYYY_MM_DD string) time.Time {
	t, err := time.Parse("2006-01-02", YYYY_MM_DD)
	if err != nil {
		panic(err)
	}
	return t
}
