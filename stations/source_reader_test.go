package stations

import (
	"os"
	"testing"
	"testing/fstest"

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
	fs     fstest.MapFS
	reader *SourceReader
}

func (this *SourceReaderFixture) Setup() {
	this.fs = make(fstest.MapFS)
	this.reader = NewSourceReader(this.fs)
}

func (this *SourceReaderFixture) TestUnhandledTypeEmitted() {
	this.reader.Do("wrong-type", this.Output)
	this.So(this.outputs, should.Equal, []any{"wrong-type"})
}
func (this *SourceReaderFixture) TestSourceFileContentReadFromDiskAndEmitted() {
	content := "article 1 content"
	this.fs["src/article-1.md"] = &fstest.MapFile{Data: []byte(content)}
	this.reader.Do(contracts.SourceFilePath("src/article-1.md"), this.Output)
	this.So(this.outputs, should.Equal, []any{
		contracts.SourceFile(content),
	})
}
func (this *SourceReaderFixture) TestIOError() {
	delete(this.fs, "src/article-1.md")
	this.reader.Do(contracts.SourceFilePath("src/article-1.md"), this.Output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, os.ErrNotExist)
	}
}
