package board

import (
	"sort"
	"strings"

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

	for i := range 5 {
		if len(ph.crew()) > i {
			card := ph.renderCard(newCard(ph.model, ph.crew()[i]))
			content = append(content, card)
			continue
		}
		emptyCard := lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Width(20).
			Height(9).
			Render("")
		content = append(content, emptyCard)
	}

	cards := lipgloss.JoinHorizontal(lipgloss.Left, content...)

	hazards := make([]string, 0)
	for _, hazard := range ph.model.Player.Hazards {
		hazards = append(hazards, hazard.Render(ph.model.style))
	}

	return lipgloss.JoinHorizontal(lipgloss.Left,
		cards,
		lipgloss.JoinVertical(lipgloss.Top, hazards...),
	)
}

func (ph *playerHandComponent) crew() []*Crew {
	sortedCrew := make([]*Crew, 0)
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

	var pips strings.Builder
	for range factionCount {
		pips.WriteString(" " + scoreboardPip)
	}

	header := lipgloss.JoinHorizontal(
		lipgloss.Top,
		c.style.Width(width-7).MarginLeft(1).Render(name),
		c.style.Width(6).Align(lipgloss.Right).Foreground(c.Crew.Faction.Color).Render(pips.String()),
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
