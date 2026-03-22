package board

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type forHireComponent struct {
	model *Model
}

func newForHireComponent(model *Model) forHireComponent {
	return forHireComponent{
		model: model,
	}
}

func (fh *forHireComponent) render() string {
	content := make([]string, 0)
	var rows []string

	for _, card := range fh.model.Game.CrewForHire {
		if fh.model.Game.InputCrew != nil && card.ID == fh.model.Game.InputCrew.ID {
			content = append(content, fh.renderCard(newCard(fh.model, fh.model.Game.InputCrew)))
			continue
		}
		content = append(content, fh.renderCard(newCard(fh.model, card)))
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

func (fh *forHireComponent) renderCard(c *card) string {
	width := 20
	height := 14
	iconWidth := 8
	objectivesWidth := 9
	descriptionWidth := width - 2

	style := c.style.
		Border(lipgloss.NormalBorder()).
		BorderForeground(c.Crew.Faction.Color).
		Width(width).
		Height(height)

	name := c.style.Foreground(c.Crew.Faction.Color).Bold(true).Render(c.Crew.Name)
	nameWidth, _ := lipgloss.Size(name)

	var objectivesString strings.Builder
	objectives := c.Crew.Objectives
	if fh.model.Game.InputCrew != nil && fh.model.Game.InputCrew.ID == c.Crew.ID && fh.model.Game.InputCrew.Objectives != nil {
		objectives = fh.model.Game.InputCrew.Objectives
	}
	for i, objective := range objectives {
		var line strings.Builder
		if i > 0 {
			line.WriteString("\n")
		}
		line.WriteString(objective.Render(c.style, fh.model.Game.InputObjective != nil && fh.model.Game.InputObjective == objective))
		objectivesString.WriteString(line.String())
	}

	nameBar := lipgloss.JoinHorizontal(
		lipgloss.Left,
		c.style.MarginLeft(1).Render(name),
		func() string {
			if c.model.Game.InputCrew != c.Crew {
				return ""
			}
			return c.style.AlignHorizontal(lipgloss.Right).Width(width - nameWidth - 2).Render("*")
		}(),
	)

	return style.Render(lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.JoinVertical(
			lipgloss.Top,
			nameBar,
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				c.style.MarginLeft(1).MarginTop(1).Width(objectivesWidth).Render(objectivesString.String()),
				c.style.MarginLeft(1).Width(iconWidth).Foreground(c.Crew.Faction.Color).Render(c.Crew.Faction.Icon),
			),
			c.style.MarginLeft(1).MarginTop(1).Width(descriptionWidth).Render(c.description),
		),
	))
}
