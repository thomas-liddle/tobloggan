package stations

import (
	"os"
	"strings"
	"testing"
	"testing/fstest"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"

	"github.com/mdwhatcott/tobloggan/code/contracts"
)

func TestListingWriterFixture(t *testing.T) {
	gunit.Run(new(ListingWriterFixture), t)
}

type ListingWriterFixture struct {
	*gunit.Fixture
	StationFixture
	fs           fstest.MapFS
	writer       *ListingWriter
	writeFileErr error
}

func (this *ListingWriterFixture) WriteFile(filename string, data []byte, perm os.FileMode) error {
	this.fs[filename] = &fstest.MapFile{Data: data, Mode: perm}
	return this.writeFileErr
}

func (this *ListingWriterFixture) Setup() {
	this.fs = make(fstest.MapFS)
	this.writer = NewListingWriter("target/directory", this)
}

func (this *ListingWriterFixture) TestUnhandledTypeEmitted() {
	this.writer.Do("wrong-type", this.Output)
	this.So(this.outputs, should.Equal, []any{"wrong-type"})
}
func (this *ListingWriterFixture) TestArticlesWrittenToListing() {
	this.writer.Do(contracts.Article{Slug: "s1", Title: "t1", Date: date("2024-09-01"), Body: "b1"}, this.Output)
	this.writer.Do(contracts.Article{Slug: "s2", Title: "t2", Date: date("2024-09-02"), Body: "b2"}, this.Output)
	this.writer.Do(contracts.Article{Slug: "s3", Title: "t3", Date: date("2024-09-03"), Body: "b3"}, this.Output)

	this.writer.Finalize(this.Output)

	raw, err := this.fs.ReadFile("target/directory/index.html")
	content := string(raw)
	this.So(err, should.BeNil)
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

func (this *ListingWriterFixture) TestWriteFileError() {
	this.writeFileErr = boink
	this.writer.Do(contracts.Article{Slug: "s1", Title: "t1", Date: date("2024-09-01"), Body: "b1"}, this.Output)
	this.writer.Do(contracts.Article{Slug: "s2", Title: "t2", Date: date("2024-09-02"), Body: "b2"}, this.Output)
	this.writer.Do(contracts.Article{Slug: "s3", Title: "t3", Date: date("2024-09-03"), Body: "b3"}, this.Output)

	this.writer.Finalize(this.Output)

	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, boink)
	}
}
