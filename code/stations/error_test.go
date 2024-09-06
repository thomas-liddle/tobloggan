package stations

import (
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestSourcedErrorFixture(t *testing.T) {
	gunit.Run(new(SourcedErrorFixture), t)
}

type SourcedErrorFixture struct {
	*gunit.Fixture
}

func (this *SourcedErrorFixture) Test() {
	err := SourcedError(boink)
	if this.So(err, should.Wrap, boink) {
		this.So(err.Error(), should.ContainSubstring, "boink")
	}
	this.Println("Example error output:", err)
}

func (this *SourcedErrorFixture) TestNil() {
	var err error
	err = SourcedError(err)
	this.So(err, should.BeNil)
}
