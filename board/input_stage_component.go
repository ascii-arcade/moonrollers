package board

import (
	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/charmbracelet/lipgloss"
)

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

func inputComponentStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Width(30).
		Height(14).
		Padding(1).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(colors.InputStageBorder)
}
