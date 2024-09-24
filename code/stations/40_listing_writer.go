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
	fs              contracts.WriteFile
	articles        []contracts.Article
	template        string
}

func NewListingWriter(targetDirectory string, fs contracts.WriteFile, template string) *ListingWriter {
	return &ListingWriter{targetDirectory: targetDirectory, fs: fs, template: template}
}

func (this *ListingWriter) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Article:
		this.articles = append(this.articles, input)
	}
	output(input)
}
func (this *ListingWriter) Finalize(output func(any)) {
	sort.SliceStable(this.articles, func(i, j int) bool {
		return this.articles[i].Date.Before(this.articles[j].Date)
	})
	var builder strings.Builder
	for _, article := range this.articles {
		builder.WriteString("\t\t\t")
		_, _ = fmt.Fprintf(&builder, `<li><a href="%s">%s</a></li>`, article.Slug, article.Title)
		builder.WriteString("\n")
	}
	replacer := strings.NewReplacer("{{Listing}}", builder.String())
	content := replacer.Replace(this.template)
	path := filepath.Join(this.targetDirectory, "index.html")
	err := this.fs.WriteFile(path, []byte(content), 0644)
	if err != nil {
		output(contracts.Error(err))
	}
}
