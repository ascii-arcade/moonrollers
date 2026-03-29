package dice

import (
	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/charmbracelet/lipgloss"
)

type Die struct {
	Color    lipgloss.Color
	ID       string
	Symbol   string
	Selected bool
	Value    int
	Mimics   string
}

var (
	DieUnrolled = Die{Symbol: "?", Color: colors.DieUnrolled, ID: "unrolled", Value: 1}

	DieDamage   = Die{Symbol: "X", Color: colors.DieDamage, ID: "damage", Value: 1}
	DieShield   = Die{Symbol: "#", Color: colors.DieShield, ID: "shield", Value: 1}
	DieThruster = Die{Symbol: "↟", Color: colors.DieThruster, ID: "thruster", Value: 1}
	DieReactor  = Die{Symbol: "@", Color: colors.DieReactor, ID: "reactor", Value: 1}
	DieWild     = Die{Symbol: "%", Color: colors.DieWild, ID: "wild", Value: 1}
	DieExtra    = Die{Symbol: "+", Color: colors.DieExtra, ID: "extra", Value: 1}
)

func All() []Die {
	return []Die{
		DieDamage,
		DieShield,
		DieThruster,
		DieReactor,
		DieWild,
		DieExtra,
	}
}

func (d *Die) Render(style lipgloss.Style, dark bool) string {
	s := style.
		Height(1).
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(d.Color).
		Foreground(d.Color)

	if dark {
		s = s.Foreground(colors.GetDark(d.ID)).BorderForeground(colors.GetDark(d.ID))
	}

	return s.Render(d.Symbol)
}

func (d *Die) MimicsType(dieId string) bool {
	return d.Mimics == dieId
}
