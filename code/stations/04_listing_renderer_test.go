package stations

import (
	"os"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"

	"github.com/mdwhatcott/tobloggan/code/contracts"
	"github.com/mdwhatcott/tobloggan/code/html"
)

func TestListingWriterFixture(t *testing.T) {
	gunit.Run(new(ListingWriterFixture), t)
}

type ListingWriterFixture struct {
	StationFixture
	fs           fstest.MapFS
	writeFileErr error
}

func (this *ListingWriterFixture) WriteFile(filename string, data []byte, perm os.FileMode) error {
	this.fs[filename] = &fstest.MapFile{Data: data, Mode: perm}
	return this.writeFileErr
}

func (this *ListingWriterFixture) Setup() {
	this.fs = make(fstest.MapFS)
	this.station = NewListingRenderer(html.ListingTemplate)
}

func (this *ListingWriterFixture) TestArticlesWrittenToListing() {
	article1 := contracts.Article{Slug: "s1", Title: "t1", Date: date("2024-09-01"), Body: "b1"}
	article2 := contracts.Article{Slug: "s2", Title: "t2", Date: date("2024-09-02"), Body: "b2"}
	article3 := contracts.Article{Slug: "s3", Title: "t3", Date: date("2024-09-03"), Body: "b3"}
	this.do(article1)
	this.do(article2)
	this.do(article3)

	this.finalize()

	this.So(this.outputs[:3], should.Equal, []any{
		article1,
		article2,
		article3,
	})
	page := this.outputs[3].(contracts.Page)
	content := page.Content
	this.So(page.Path, should.Equal, "/")
	this.So(content, should.ContainSubstring, `href="s1"`)
	this.So(content, should.ContainSubstring, `href="s2"`)
	this.So(content, should.ContainSubstring, `href="s3"`)
	this.So(content, should.ContainSubstring, `>t1<`)
	this.So(content, should.ContainSubstring, `>t2<`)
	this.So(content, should.ContainSubstring, `>t3<`)
	d1 := strings.Index(content, ">t1<")
	d2 := strings.Index(content, ">t2<")
	d3 := strings.Index(content, ">t3<")
	this.So(d1, should.BeLessThan, d2)
	this.So(d2, should.BeLessThan, d3)
}
