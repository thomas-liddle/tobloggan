package stations

type Markdown interface {
	Convert(content string) (string, error)
}

//type MarkdownConverter struct{}

//func (this *MarkdownConverter) Do(input any, output func(any)) {
//    TODO: given a contracts.Article, use the provided Markdown interface to convert and re-assign the Body field.
//}
