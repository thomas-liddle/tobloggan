package stations

import (
	"path/filepath"
	"strings"

	"github.com/mdwhatcott/tobloggan/code/contracts"
)

type FileSystemWriter interface {
	contracts.MkdirAll
	contracts.WriteFile
}
type ArticleWriter struct {
	targetDirectory string
	fs              FileSystemWriter
	template        string
}

func NewArticleWriter(targetDirectory string, fs FileSystemWriter, template string) *ArticleWriter {
	return &ArticleWriter{
		targetDirectory: targetDirectory,
		fs:              fs,
		template:        template,
	}
}
func (this *ArticleWriter) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Article:
		path := filepath.Join(this.targetDirectory, input.Slug, "index.html")
		err := this.fs.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			output(contracts.Error(err))
			return
		}
		content := strings.NewReplacer(
			"{{Title}}", input.Title,
			"{{Slug}}", input.Slug,
			"{{Date}}", input.Date.Format("January 2, 2006"),
			"{{Body}}", input.Body,
		).Replace(this.template)
		err = this.fs.WriteFile(path, []byte(content), 0644)
		if err != nil {
			output(contracts.Error(err))
			return
		}
		output(input)
	default:
		output(input)
	}
}
