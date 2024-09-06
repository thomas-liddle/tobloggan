package stations

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	"github.com/mdwhatcott/tobloggan/code/contracts"
)

type ListingWriter struct {
	targetDirectory string
	fs              contracts.FSWriter
	articles        []contracts.Article
}

func NewListingWriter(targetDirectory string, fs contracts.FSWriter) *ListingWriter {
	return &ListingWriter{targetDirectory: targetDirectory, fs: fs}
}

func (this *ListingWriter) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Article:
		this.articles = append(this.articles, input)
	default:
		output(input)
	}
}
func (this *ListingWriter) Finalize(output func(any)) {
	sort.Slice(this.articles, func(i, j int) bool {
		return this.articles[i].Date.Before(this.articles[j].Date)
	})
	var builder strings.Builder
	for _, article := range this.articles {
		_, _ = fmt.Fprintf(&builder, "\t\t\t"+`<li><a href="%s">%s</a></li>`+"\n", article.Slug, article.Title)
	}
	replacer := strings.NewReplacer(
		"{{CSS}}", css,
		"{{Listing}}", builder.String(),
	)
	content := replacer.Replace(listingTemplate)
	path := filepath.Join(this.targetDirectory, "index.html")
	err := this.fs.WriteFile(path, []byte(content), 0644)
	if err != nil {
		output(contracts.Error(err))
	}
}

const listingTemplate = `<!doctype html>
<html lang="en">
    <head>
        <title>Your Name Here</title>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta name="author" content="Your Name Here">
        <meta name="description" content="description">
        <link rel="canonical" href="https://your-domain-here.com">
        {{CSS}}
    </head>

    <body>
        <nav>
            <a href="/about/">About</a>
        </nav>
        <h1>Your Name Here</h1>
        <p>Something about yourself and this website here.</p>
		<ul>
		{{Listing}}
		</ul>
        <br>
        <br>
    </body>
</html>
`
