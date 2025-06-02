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

	c.description = c.model.lang().Get("crew_abilities." + crew.ID)
	for _, die := range dice.All() {
		symbolStyle := c.style.Foreground(die.Color).Bold(true).Italic(true)
		findSingular := language.Languages["EN"].Get("dice." + die.ID)
		findPlural := language.Languages["EN"].Get("dice_plural." + die.ID)

		singular := c.model.lang().Get("dice." + die.ID)
		plural := c.model.lang().Get("dice_plural." + die.ID)

		singularValue := symbolStyle.Render(fmt.Sprintf("%s %s", die.Symbol, singular))
		pluralValue := symbolStyle.Render(fmt.Sprintf("%s %s", die.Symbol, plural))

		c.description = strings.ReplaceAll(c.description, "%"+findSingular+"%", singularValue)
		c.description = strings.ReplaceAll(c.description, "%"+findPlural+"%", pluralValue)
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
			line.WriteString(c.style.Foreground(colors.Hazard).Render(hazard))
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
