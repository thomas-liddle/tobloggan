package stations

import (
	"testing"

	"github.com/mdwhatcott/tobloggan/code/contracts"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestBaseURLRewriterFixture(t *testing.T) {
	gunit.Run(new(BaseURLRewriterFixture), t)
}

type BaseURLRewriterFixture struct {
	*gunit.Fixture
	StationFixture
	rewriter *BaseURLRewriter
}

func (this *BaseURLRewriterFixture) Setup() {
	this.rewriter = NewBaseURLRewriter("https://base-url.com/blog")
}

func (this *BaseURLRewriterFixture) TestUnhandledTypeEmitted() {
	this.rewriter.Do("wrong-type", this.Output)
	this.So(this.outputs, should.Equal, []any{"wrong-type"})
}
func (this *BaseURLRewriterFixture) TestRewriteWithBaseURL() {
	input := contracts.Page{
		Content: `<a href="/some-other-page">click me</a>` +
			`<a href="/still-another-path">click me</a>`,
	}
	this.rewriter.Do(input, this.Output)
	this.So(this.outputs, should.Equal, []any{contracts.Page{
		Content: `<a href="https://base-url.com/blog/some-other-page">click me</a>` +
			`<a href="https://base-url.com/blog/still-another-path">click me</a>`,
	}})
}
