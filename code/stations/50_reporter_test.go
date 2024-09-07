package stations

import (
	"sync/atomic"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestReporterFixture(t *testing.T) {
	gunit.Run(new(ReporterFixture), t)
}

type ReporterFixture struct {
	*gunit.Fixture
	reporter *Reporter
	failure  *atomic.Bool
}

func (this *ReporterFixture) Setup() {
	this.failure = new(atomic.Bool)
	this.reporter = NewReporter(this, this.failure)
}

func (this *ReporterFixture) TestErrReportsFailure() {
	this.reporter.Do(boink, nil)
	this.So(this.failure.Load(), should.BeTrue)
}
