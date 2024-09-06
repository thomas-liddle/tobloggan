package markdown

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

type Converter struct {
	buffer    *bytes.Buffer
	converter goldmark.Markdown
}

func NewConverter() *Converter {
	return &Converter{
		buffer: new(bytes.Buffer),
		converter: goldmark.New(
			goldmark.WithRendererOptions(
				html.WithUnsafe(),
			),
			goldmark.WithExtensions(
				extension.GFM,
				extension.DefinitionList,
				extension.Footnote,
			),
		),
	}
}

func (this *Converter) Convert(content string) (string, error) {
	this.buffer.Reset()
	err := this.converter.Convert([]byte(content), this.buffer)
	return this.buffer.String(), err
}
