package stations

import (
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestDraftRemovalFixture(t *testing.T) {
	gunit.Run(new(DraftRemovalFixture), t)
}

type DraftRemovalFixture struct {
	StationFixture
}

func (this *DraftRemovalFixture) Setup() {
	this.station = NewDraftRemoval()
}

func (this *DraftRemovalFixture) TestDraftDropped() {
	this.do(draftArticle)
	this.So(this.outputs, should.HaveLength, 0)
}
func (this *DraftRemovalFixture) TestNonDraftRetained() {
	this.do(goodArticle)
	this.So(this.outputs, should.HaveLength, 1)
	this.assertOutputs(goodArticle)
}

var draftArticle = contracts.Article{
	Date:  time.Date(2024, 9, 4, 0, 0, 0, 0, time.UTC),
	Slug:  "/article/1",
	Title: "Article 1",
	Draft: true,
	Body:  "The contents of article 1.",
}
