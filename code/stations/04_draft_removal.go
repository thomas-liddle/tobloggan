package stations

type DraftRemoval struct{}

func (this *DraftRemoval) Do(input any, output func(any)) {
	// TODO: given a contracts.Article, only output it if !input.Draft.
}
