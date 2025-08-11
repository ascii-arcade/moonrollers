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
	hazardValue := hazardStyle.Render(fmt.Sprintf("%s %s", deck.Hazard, "hazard"))
	c.description = strings.ReplaceAll(c.description, "%hazard%", hazardValue)

	return c
}
