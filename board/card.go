package board

import (
	"strings"

	"github.com/ascii-arcade/moonrollers/deck"
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
		model:       model,
		Crew:        crew,
		description: crew.Ability.Description,
		style:       model.style,
	}

	for _, die := range deck.AllDice() {
		pluralValue := c.style.Foreground(die.Color).Bold(true).Italic(true).Render(die.Symbol + " " + die.Name + "s")
		c.description = strings.ReplaceAll(c.description, "%"+die.Name+"s%", pluralValue)
		singularValue := c.style.Foreground(die.Color).Bold(true).Italic(true).Render(die.Symbol + " " + die.Name)
		c.description = strings.ReplaceAll(c.description, "%"+die.Name+"%", singularValue)
	}

	return c
}

func (c *card) render() string {
	return c.style.Render(c.Crew.Name) + "\n" + c.style.Render(c.description)
}
