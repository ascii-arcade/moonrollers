package board

import tea "github.com/charmbracelet/bubbletea"

type inputStageComponent interface {
	render() string
	update(msg tea.Msg) (any, tea.Cmd)
}

type inputStageEmptyComponent struct {
	model *Model
}

func newInputStageEmptyComponent(model *Model) inputStageEmptyComponent {
	return inputStageEmptyComponent{
		model: model,
	}
}

func (c inputStageEmptyComponent) render() string {
	return ""
}

func (s inputStageEmptyComponent) update(msg tea.Msg) (any, tea.Cmd) {
	return s.model, nil
}
