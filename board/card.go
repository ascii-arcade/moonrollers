package board

import (
	"strings"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/ascii-arcade/moonrollers/deck"
	"github.com/ascii-arcade/moonrollers/dice"
	"github.com/charmbracelet/lipgloss"
)

const (
	emptyPip = "◇"
	fullPip  = "◆"
)

type card struct {
	model       *Model
	Crew        *deck.Crew
	description string
	style       lipgloss.Style
}

func newCard(model *Model, crew *deck.Crew) *card {
	c := &card{
		model:       model,
		Crew:        crew,
		description: crew.Ability.Description,
		style:       model.style,
	}

	for _, die := range dice.All() {
		pluralValue := c.style.Foreground(die.Color).Bold(true).Italic(true).Render(die.Symbol + " " + die.Name + "s")
		c.description = strings.ReplaceAll(c.description, "%"+die.Name+"s%", pluralValue)
		singularValue := c.style.Foreground(die.Color).Bold(true).Italic(true).Render(die.Symbol + " " + die.Name)
		c.description = strings.ReplaceAll(c.description, "%"+die.Name+"%", singularValue)
	}

	return c
}

func (c *card) render() string {
	style := c.style.
		Border(lipgloss.NormalBorder()).
		BorderForeground(c.Crew.Faction.Color).
		Width(11).
		Height(7)

	var sb strings.Builder
	sb.WriteString(c.style.Foreground(c.Crew.Faction.Color).Bold(true).Render(c.Crew.Name))
	sb.WriteString("\n\n")

	for _, objective := range c.Crew.Objectives {
		var line strings.Builder
		line.WriteString(" ")
		line.WriteString(c.style.Foreground(objective.Type.Color).Render(objective.Type.Symbol))
		line.WriteString(" ")
		if objective.Hazard {
			line.WriteString(c.style.Foreground(colors.Hazard).Render("!"))
		} else {
			line.WriteString(" ")
		}
		for range objective.Amount {
			line.WriteString(emptyPip)
		}
		line.WriteString("\n")
		sb.WriteString(line.String())
	}

	return style.Render(sb.String())
}
