package stations

import (
	"testing"

	"github.com/mdwhatcott/tobloggan/code/contracts"
	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestArticleRendererFixture(t *testing.T) {
	gunit.Run(new(ArticleRendererFixture), t)
}

type ArticleRendererFixture struct {
	*gunit.Fixture
	StationFixture
}

func (this *ArticleRendererFixture) Setup() {
	this.station = NewArticleRenderer("{{Slug}}\n{{Title}}\n{{Date}}\n{{Body}}")
}

func (this *ArticleRendererFixture) TestUnhandledTypeEmitted() {
	this.do("wrong-type")
	this.So(this.outputs, should.Equal, []any{"wrong-type"})
}
func (this *ArticleRendererFixture) TestRendering() {
	input := contracts.Article{
		Slug:  "the/slug",
		Title: "The Title",
		Date:  date("2024-09-25"),
		Body:  "The body.",
	}
	this.do(input)
	this.So(this.outputs, should.Equal, []any{
		contracts.Page{
			Path:    "the/slug",
			Content: "the/slug\nThe Title\nSeptember 25, 2024\nThe body.",
		},
	})
}
