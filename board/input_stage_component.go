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
	return inputComponentStyle(false).Render("")
}

func inputComponentStyle(isCenter bool) lipgloss.Style {
	align := lipgloss.Left
	if isCenter {
		align = lipgloss.Center
	}

	return lipgloss.NewStyle().
		Width(30).
		Height(13).
		PaddingLeft(1).PaddingRight(1).
		Align(align).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(colors.InputStageBorder)
}
