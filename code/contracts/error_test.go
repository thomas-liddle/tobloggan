package contracts

import (
	"errors"
	"runtime"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestErrorFixture(t *testing.T) {
	gunit.Run(new(ErrorFixture), t)
}

type ErrorFixture struct {
	*gunit.Fixture
}

func (this *ErrorFixture) TestError() {
	boink := errors.New("boink")
	err := Error(boink)
	_, file, _, _ := runtime.Caller(0)
	if this.So(err, should.Wrap, boink) {
		this.So(err.Error(), should.ContainSubstring, "boink")
		this.So(err.Error(), should.ContainSubstring, file)
	}
	this.Println("Example error output:", err)
}
func (this *ErrorFixture) TestErrorf() {
	boink := errors.New("boink")
	err := Errorf("%w", boink)
	_, file, _, _ := runtime.Caller(0)
	if this.So(err, should.Wrap, boink) {
		this.So(err.Error(), should.ContainSubstring, "boink")
		this.So(err.Error(), should.ContainSubstring, file)
	}
	this.Println("Example error output:", err)
}
