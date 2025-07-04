package board

import "github.com/charmbracelet/lipgloss"

type forHire struct {
	model *Model
}

func newForHire(model *Model) forHire {
	return forHire{
		model: model,
	}
}

func (fh *forHire) render() string {
	content := make([]string, 0)
	var rows []string

	for _, card := range fh.model.Game.CrewForHire {
		content = append(content, newCard(fh.model, card).renderForHire())
	}

	if len(content) > 0 {
		row1 := lipgloss.JoinHorizontal(lipgloss.Left, content[:min(3, len(content))]...)
		rows = append(rows, row1)
	}

	if len(content) > 3 {
		row2 := lipgloss.JoinHorizontal(lipgloss.Left, content[3:min(6, len(content))]...)
		rows = append(rows, row2)
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}
