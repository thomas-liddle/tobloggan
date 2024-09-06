package stations

import (
	"errors"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestStackTraceErrorFixture(t *testing.T) {
	gunit.Run(new(StackTraceErrorFixture), t)
}

type StackTraceErrorFixture struct {
	*gunit.Fixture
}

func (this *StackTraceErrorFixture) Test() {
	gopherErr := errors.New("gophers")
	err := StackTraceError(gopherErr)
	if this.So(err, should.Wrap, gopherErr) {
		this.So(err.Error(), should.Contain, "gophers")
		this.So(err.Error(), should.Contain, "stack:")
	}
}

func (this *StackTraceErrorFixture) TestNil() {
	var err error
	err = StackTraceError(err)
	this.So(err, should.BeNil)
}
