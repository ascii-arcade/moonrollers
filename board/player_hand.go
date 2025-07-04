package board

import (
	"sort"
	"strconv"

	"github.com/ascii-arcade/moonrollers/deck"
	"github.com/charmbracelet/lipgloss"
)

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

	for _, crew := range ph.crew() {
		card := ph.renderCard(newCard(ph.model, crew))
		content = append(content, card)
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, content...)
}

func (ph *playerHand) crew() []*deck.Crew {
	sortedCrew := make([]*deck.Crew, 0)
	for _, crew := range ph.model.Player.Crew {
		sortedCrew = append(sortedCrew, crew)
	}
	sort.Slice(sortedCrew, func(i, j int) bool {
		return sortedCrew[i].Faction.SortOrder < sortedCrew[j].Faction.SortOrder
	})

	return sortedCrew
}

func (ph *playerHand) renderCard(c *card) string {
	width := 20
	height := 9
	descriptionWidth := width - 2

	style := c.style.
		Border(lipgloss.NormalBorder()).
		BorderForeground(c.Crew.Faction.Color).
		Width(width).
		Height(height)

	name := c.style.Foreground(c.Crew.Faction.Color).Bold(true).Render(c.Crew.Name)

	return style.Render(lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.JoinVertical(
			lipgloss.Top,
			c.style.MarginLeft(1).Render(name),
			c.style.MarginLeft(1).Render(strconv.Itoa(c.model.Player.CrewCount[c.Crew.Faction.Name])),
			c.style.MarginLeft(1).MarginTop(1).Width(descriptionWidth).Render(c.description),
		),
	))
}
