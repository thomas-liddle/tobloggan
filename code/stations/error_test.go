package stations

import (
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
	err := SourcedError(boink)
	if this.So(err, should.Wrap, boink) {
		this.So(err.Error(), should.ContainSubstring, "boink")
	}
	this.Println("Example error output:", err)
}

func (this *StackTraceErrorFixture) TestNil() {
	var err error
	err = SourcedError(err)
	this.So(err, should.BeNil)
}
