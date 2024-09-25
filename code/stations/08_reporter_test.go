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
	StationFixture
	log     bytes.Buffer
	failure *atomic.Bool
}

func (this *ReporterFixture) Setup() {
	this.failure = new(atomic.Bool)
	this.station = NewReporter(log.New(io.MultiWriter(this, &this.log), "", 0), this.failure)
}

func (this *ReporterFixture) TestErrReportsFailure() {
	this.do(boink)
	this.So(this.failure.Load(), should.BeTrue)
	this.So(this.log.String(), should.ContainSubstring, boink.Error())
}
func (this *ReporterFixture) TestReportsPage() {
	this.do(contracts.Page{Path: "/output/a-tale-of-two-cities/index.html"})
	this.So(this.log.String(), should.ContainSubstring, "/output/a-tale-of-two-cities/index.html")
}
