package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/ascii-arcade/moonrollers/deck"
	"github.com/ascii-arcade/moonrollers/dice"
	"github.com/ascii-arcade/moonrollers/language"
	"github.com/charmbracelet/lipgloss"
)

const (
	emptyPip = "◇"
	fullPip  = "◆"
	hazard   = "!"
)

type card struct {
	model       *Model
	Crew        *deck.Crew
	description string
	style       lipgloss.Style
}

func newCard(model *Model, crew *deck.Crew) *card {
	c := &card{
		model: model,
		Crew:  crew,
		style: model.style,
	}

	c.description = c.model.lang().Get("crew_abilities", crew.ID)
	for _, die := range dice.All() {
		symbolStyle := c.style.Foreground(die.Color).Bold(true).Italic(true)
		findSingular := language.Languages["EN"].Get("dice", die.ID)
		findPlural := language.Languages["EN"].Get("dice_plural", die.ID)

		singular := c.model.lang().Get("dice", die.ID)
		plural := c.model.lang().Get("dice_plural", die.ID)

		singularValue := symbolStyle.Render(fmt.Sprintf("%s %s", die.Symbol, singular))
		pluralValue := symbolStyle.Render(fmt.Sprintf("%s %s", die.Symbol, plural))

		c.description = strings.ReplaceAll(c.description, "%"+findSingular+"%", singularValue)
		c.description = strings.ReplaceAll(c.description, "%"+findPlural+"%", pluralValue)
	}
	hazardStyle := c.style.Foreground(colors.Hazard).Bold(true).Italic(true)
	hazardValue := hazardStyle.Render(fmt.Sprintf("%s %s", hazard, "hazard"))
	c.description = strings.ReplaceAll(c.description, "%hazard%", hazardValue)

	return c
}

func (c *card) renderForHand() string {
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
			c.style.MarginLeft(1).MarginTop(1).Width(descriptionWidth).Render(c.description),
		),
	))
}

func (c *card) renderForHire() string {
	width := 20
	height := 14
	iconWidth := 8
	objectivesWidth := 7
	descriptionWidth := width - 2

	style := c.style.
		Border(lipgloss.NormalBorder()).
		BorderForeground(c.Crew.Faction.Color).
		Width(width).
		Height(height)

	name := c.style.Foreground(c.Crew.Faction.Color).Bold(true).Render(c.Crew.Name)

	var objectives strings.Builder
	for i, objective := range c.Crew.Objectives {
		var line strings.Builder
		if i > 0 {
			line.WriteString("\n")
		}
		line.WriteString(c.style.Foreground(objective.Type.Color).Render(objective.Type.Symbol))
		line.WriteString(" ")
		if objective.Hazard {
			line.WriteString(c.style.Foreground(colors.Hazard).Render(hazard))
		} else {
			line.WriteString(" ")
		}
		for range objective.Amount {
			line.WriteString(emptyPip)
		}
		objectives.WriteString(line.String())
	}

	return style.Render(lipgloss.JoinHorizontal(
		lipgloss.Left,
		lipgloss.JoinVertical(
			lipgloss.Top,
			c.style.MarginLeft(1).Render(name),
			lipgloss.JoinHorizontal(
				lipgloss.Left,
				c.style.MarginLeft(1).MarginTop(1).Width(objectivesWidth).Render(objectives.String()),
				c.style.MarginLeft(3).Width(iconWidth).Foreground(c.Crew.Faction.Color).Render(c.Crew.Faction.Icon),
			),
			c.style.MarginLeft(1).MarginTop(1).Width(descriptionWidth).Render(c.description),
		),
	))
}
