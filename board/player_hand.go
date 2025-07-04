package board

import "github.com/charmbracelet/lipgloss"

type playerHand struct {
	model *Model
}

func newPlayerHand(model *Model) playerHand {
	return playerHand{
		model: model,
	}
}

func (ph *playerHand) render() string {
	content := make([]string, 0)

	for _, crew := range ph.model.Player.Crew {
		card := newCard(ph.model, crew).renderForHand()
		content = append(content, card)
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, content...)
}
