package stations

import (
	"os"
	"testing"
	"testing/fstest"

	"github.com/mdwhatcott/tobloggan/contracts"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestSourceScannerFixture(t *testing.T) {
	gunit.Run(new(SourceScannerFixture), t)
}

type SourceScannerFixture struct {
	*gunit.Fixture
	StationFixture
	fs      fstest.MapFS
	scanner *SourceScanner
}

func (this *SourceScannerFixture) Setup() {
	this.fs = make(fstest.MapFS)
	this.fs["src/article-1.md"] = &fstest.MapFile{Data: []byte("article 1 source")}
	this.fs["src/article-2.txt"] = &fstest.MapFile{Data: []byte("article 2 source")}
	this.fs["src/article-3.md"] = &fstest.MapFile{Data: []byte("article 3 source")}
	this.fs["src/inner/article-4.md"] = &fstest.MapFile{Data: []byte("article 4 source")}
	this.scanner = NewSourceScanner(this.fs)
}

func (this *SourceScannerFixture) TestUnhandledTypeEmitted() {
	this.scanner.Do("wrong-type", this.Output)
	this.So(this.outputs, should.Equal, []any{"wrong-type"})
}
func (this *SourceScannerFixture) TestGivenASourceDirectoryThatDoesNotExist_EmitError() {
	clear(this.fs)
	this.scanner.Do(contracts.SourceDirectory("NOT-THERE"), this.Output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, os.ErrNotExist)
	}
}
func (this *SourceScannerFixture) TestGivenASourceDirectoryWithBlogSourceFiles_EmitAllContainingBlogSourceFilePaths() {
	this.scanner.Do(contracts.SourceDirectory("src"), this.Output)
	this.So(this.outputs, should.Equal, []any{
		contracts.SourceFilePath("src/article-1.md"),
		contracts.SourceFilePath("src/article-3.md"),
		contracts.SourceFilePath("src/inner/article-4.md"),
	})
}
