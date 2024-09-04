package stations

type StationFixture struct {
	outputs []any
}

func (this *StationFixture) Output(v any) {
	this.outputs = append(this.outputs, v)
}
