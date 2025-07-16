package board

type inputStageComponent interface {
	render() string
}

type inputStageEmptyComponent struct{}

func newInputStageEmptyComponent() inputStageEmptyComponent {
	return inputStageEmptyComponent{}
}

func (c inputStageEmptyComponent) render() string {
	return ""
}
