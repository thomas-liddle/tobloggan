package stations

import (
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

var pastArticle = contracts.Article{
	Date: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
}

var currentArticle = contracts.Article{
	Date: time.Now().Add(-time.Hour),
}

var futureArticle = contracts.Article{
	Date: time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC),
}

func TestFutureRemovalFixture(t *testing.T) {
	gunit.Run(new(FutureRemovalFixture), t)
}

type FutureRemovalFixture struct {
	StationFixture
}

func (this *FutureRemovalFixture) Setup() {
	this.station = NewFutureRemoval(time.Now())
}

func (this *FutureRemovalFixture) TestPastArticleKept() {
	this.do(pastArticle)
	this.assertOutputs(pastArticle)
}

func (this *FutureRemovalFixture) TestCurrentArticleKept() {
	this.do(currentArticle)
	this.assertOutputs(currentArticle)
}

func (this *FutureRemovalFixture) TestFutureArticleDropped() {
	this.do(futureArticle)
	this.So(this.outputs, should.HaveLength, 0)
}
