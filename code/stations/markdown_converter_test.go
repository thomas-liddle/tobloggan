package stations

import (
	"testing"

	"github.com/mdwhatcott/tobloggan/code/contracts"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestMarkdownConverterFixture(t *testing.T) {
	gunit.Run(new(MarkdownConverterFixture), t)
}

type MarkdownConverterFixture struct {
	*gunit.Fixture
	StationFixture
	markdownErr error
}

func (this *MarkdownConverterFixture) Convert(content string) (string, error) {
	return content + " CONVERTED", this.markdownErr
}

func (this *MarkdownConverterFixture) Setup() {
	this.station = NewMarkdownConverter(this)
}

func (this *MarkdownConverterFixture) TestUnhandledTypeEmitted() {
	this.station.Do("wrong-type", this.Output)
	this.So(this.outputs, should.Equal, []any{"wrong-type"})
}
func (this *MarkdownConverterFixture) TestBodyConverted() {
	input := contracts.Article{Body: article1Content}
	this.station.Do(input, this.Output)
	this.So(this.outputs, should.Equal, []any{
		contracts.Article{Body: article1Content + " CONVERTED"},
	})
}
func (this *MarkdownConverterFixture) TestInvalidMarkdown() {
	this.markdownErr = boink
	this.station.Do(contracts.Article{Body: article1Content}, this.Output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, errMalformedSource)
	}
}
