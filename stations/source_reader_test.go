package stations

import (
	"os"
	"testing"
	"testing/fstest"
	"time"

	"github.com/mdwhatcott/tobloggan/contracts"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestSourceReaderFixture(t *testing.T) {
	gunit.Run(new(SourceReaderFixture), t)
}

type SourceReaderFixture struct {
	*gunit.Fixture
	StationFixture
	fs          fstest.MapFS
	reader      *SourceReader
	markdownErr error
}

func (this *SourceReaderFixture) Setup() {
	this.fs = make(fstest.MapFS)
	this.reader = NewSourceReader(this.fs, this)
}

func (this *SourceReaderFixture) Convert(content string) (string, error) {
	return content + " CONVERTED", this.markdownErr
}
func (this *SourceReaderFixture) TestArticleMetaAndContentReadFromDiskAndEmitted() {
	this.fs["src/article-1.md"] = &fstest.MapFile{Data: []byte(article1Content)}
	this.reader.Do(contracts.SourceFilePath("src/article-1.md"), this.Output)
	this.So(this.outputs, should.Equal, []any{
		contracts.Source{
			Slug:  "/article/1",
			Title: "Article 1",
			Date:  time.Date(2024, time.September, 4, 0, 0, 0, 0, time.UTC),
			Body:  "The contents of article 1. CONVERTED",
		},
	})
}
func (this *SourceReaderFixture) TestIOError() {
	delete(this.fs, "src/article-1.md")
	this.reader.Do(contracts.SourceFilePath("src/article-1.md"), this.Output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, os.ErrNotExist)
	}
}
func (this *SourceReaderFixture) TestMissingDivider() {
	this.fs["src/article-1.md"] = &fstest.MapFile{Data: []byte("{} Content without separator")}
	this.reader.Do(contracts.SourceFilePath("src/article-1.md"), this.Output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, contracts.ErrMalformedSource)
	}
}
func (this *SourceReaderFixture) TestMalformedMetadata() {
	this.fs["src/article-1.md"] = &fstest.MapFile{Data: []byte("{bad-json}\n+++\nContent")}
	this.reader.Do(contracts.SourceFilePath("src/article-1.md"), this.Output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, contracts.ErrMalformedSource)
	}
}
func (this *SourceReaderFixture) TestInvalidMarkdown() {
	this.markdownErr = boink
	this.fs["src/article-1.md"] = &fstest.MapFile{Data: []byte(article1Content)}
	this.reader.Do(contracts.SourceFilePath("src/article-1.md"), this.Output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, contracts.ErrMalformedSource)
	}
}

const article1Content = `{
	"date": "2024-09-04T00:00:00Z",
	"slug": "/article/1",
	"title": "Article 1"
}

+++

The contents of article 1.`
