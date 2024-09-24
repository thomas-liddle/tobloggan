package stations

import (
	"bytes"
	"io"
	"log"
	"sync/atomic"
	"testing"

	"github.com/mdwhatcott/tobloggan/code/contracts"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestReporterFixture(t *testing.T) {
	gunit.Run(new(ReporterFixture), t)
}

type ReporterFixture struct {
	*gunit.Fixture
	StationFixture
	log      bytes.Buffer
	reporter *Reporter
	failure  *atomic.Bool
}

func (this *ReporterFixture) Setup() {
	this.failure = new(atomic.Bool)
	this.reporter = NewReporter(log.New(io.MultiWriter(this, &this.log), "", 0), this.failure)
}
func (this *ReporterFixture) Teardown() {
	this.So(this.outputs, should.BeEmpty)
}

func (this *ReporterFixture) TestErrReportsFailure() {
	this.reporter.Do(boink, nil)
	this.So(this.failure.Load(), should.BeTrue)
	this.So(this.log.String(), should.ContainSubstring, boink.Error())
}
func (this *ReporterFixture) TestReportsArticle() {
	this.reporter.Do(contracts.Article{Title: "A Tale of Two Cities"}, this.Output)
	this.So(this.log.String(), should.ContainSubstring, "A Tale of Two Cities")
}
func (this *ReporterFixture) TestReportsWhatever() {
	this.reporter.Do(42, nil)
	this.So(this.log.String(), should.ContainSubstring, "unexpected type: int")
}
