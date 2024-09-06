package stations

import (
	"os"
	"testing"
	"testing/fstest"
	"time"

	"github.com/mdwhatcott/tobloggan/code/contracts"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestArticleWriterFixture(t *testing.T) {
	gunit.Run(new(ArticleWriterFixture), t)
}

type ArticleWriterFixture struct {
	*gunit.Fixture
	StationFixture
	fs           fstest.MapFS
	writer       *ArticleWriter
	writeFileErr error
	mkdirAllErr  error
}

func (this *ArticleWriterFixture) MkdirAll(path string, perm os.FileMode) error {
	this.fs[path] = &fstest.MapFile{Mode: perm}
	return this.mkdirAllErr
}
func (this *ArticleWriterFixture) WriteFile(filename string, data []byte, perm os.FileMode) error {
	this.fs[filename] = &fstest.MapFile{Data: data, Mode: perm}
	return this.writeFileErr
}
func (this *ArticleWriterFixture) Setup() {
	this.fs = fstest.MapFS{}
	this.writer = NewArticleWriter("target/directory", this)
}
func (this *ArticleWriterFixture) TestUnhandledTypeEmitted() {
	this.writer.Do("wrong-type", this.Output)
	this.So(this.outputs, should.Equal, []any{"wrong-type"})
}
func (this *ArticleWriterFixture) TestMkdirErr() {
	this.mkdirAllErr = boink
	input := contracts.Article{Slug: "s", Title: "t", Body: "b"}
	this.writer.Do(input, this.Output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, boink)
	}
	this.So(this.fs, should.NotContainKey, "target/directory/the/slug/index.html")
}
func (this *ArticleWriterFixture) TestWriteFileErr() {
	this.writeFileErr = boink
	input := contracts.Article{Slug: "s", Title: "t", Body: "b"}
	this.writer.Do(input, this.Output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, boink)
	}
}
func (this *ArticleWriterFixture) TestHappyPath() {
	input := contracts.Article{
		Slug:  "/the/slug",
		Title: "The Title",
		Date:  time.Date(2024, time.September, 6, 0, 0, 0, 0, time.UTC),
		Body:  "The Body",
	}
	this.writer.Do(input, this.Output)

	this.So(this.outputs, should.Equal, []any{input})
	content, err := this.fs.ReadFile("target/directory/the/slug/index.html")
	this.So(err, should.BeNil)
	this.So(string(content), should.ContainSubstring, input.Slug)
	this.So(string(content), should.ContainSubstring, input.Title)
	this.So(string(content), should.ContainSubstring, "September 6, 2024")
	this.So(string(content), should.ContainSubstring, input.Body)
}
