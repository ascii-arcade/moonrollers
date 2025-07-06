package board

import (
	"sort"

	"github.com/ascii-arcade/moonrollers/deck"
	"github.com/charmbracelet/lipgloss"
)

type playerHandComponent struct {
	model *Model
}

func newPlayerHandComponent(model *Model) playerHandComponent {
	return playerHandComponent{
		model: model,
	}
}

func (ph *playerHandComponent) render() string {
	content := make([]string, 0)

	for _, crew := range ph.crew() {
		card := ph.renderCard(newCard(ph.model, crew))
		content = append(content, card)
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, content...)
}

func (ph *playerHandComponent) crew() []*deck.Crew {
	sortedCrew := make([]*deck.Crew, 0)
	for _, crew := range ph.model.Player.Crew {
		sortedCrew = append(sortedCrew, crew)
	}
	sort.Slice(sortedCrew, func(i, j int) bool {
		return sortedCrew[i].Faction.SortOrder < sortedCrew[j].Faction.SortOrder
	})

	return sortedCrew
}

func (ph *playerHandComponent) renderCard(c *card) string {
	width := 20
	height := 9
	descriptionWidth := width - 2
	factionCount := c.model.Player.CrewCount[c.Crew.Faction.Name]

	style := c.style.
		Border(lipgloss.NormalBorder()).
		BorderForeground(c.Crew.Faction.Color).
		Width(width).
		Height(height)

	name := c.style.Foreground(c.Crew.Faction.Color).Bold(true).Render(c.Crew.Name)

	pips := ""
	for range factionCount {
		pips += " " + scoreboardPip
	}

	header := lipgloss.JoinHorizontal(
		lipgloss.Top,
		c.style.Width(width-7).MarginLeft(1).Render(name),
		c.style.Width(6).Align(lipgloss.Right).Foreground(c.Crew.Faction.Color).Render(pips),
	)

	return style.Render(lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.JoinVertical(
			lipgloss.Top,
			header,
			c.style.MarginLeft(1).MarginTop(1).Width(descriptionWidth).Render(c.description),
		),
	))
}
