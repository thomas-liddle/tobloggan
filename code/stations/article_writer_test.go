package stations

import (
	"os"
	"testing"
	"testing/fstest"
	"time"

	"github.com/mdwhatcott/tobloggan/code/contracts"
	"github.com/smarty/gunit"
)

func TestArticleWriterFixture(t *testing.T) {
	gunit.Run(new(ArticleWriterFixture), t)
}

type ArticleWriterFixture struct {
	*gunit.Fixture
	*StationFixture
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
	this.writer = NewArticleWriter("/target/directory", this)
}

func (this *ArticleWriterFixture) Test() {
	input := contracts.Article{
		Slug:  "/the/slug",
		Title: "The Title",
		Date:  time.Date(2024, time.September, 6, 0, 0, 0, 0, time.UTC),
		Body:  "The Body",
	}
	this.writer.Do(input, this.Output)
	
}
