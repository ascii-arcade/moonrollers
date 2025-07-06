package dice

import (
	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/charmbracelet/lipgloss"
)

type Die struct {
	Color  lipgloss.Color
	ID     string
	Symbol string
}

var (
	DieUnrolled = Die{Symbol: "?", Color: colors.DieUnrolled, ID: "unrolled"}

	DieDamage   = Die{Symbol: "X", Color: colors.DieDamage, ID: "damage"}
	DieShield   = Die{Symbol: "#", Color: colors.DieShield, ID: "shield"}
	DieThruster = Die{Symbol: "↟", Color: colors.DieThruster, ID: "thruster"}
	DieReactor  = Die{Symbol: "@", Color: colors.DieReactor, ID: "reactor"}
	DieWild     = Die{Symbol: "%", Color: colors.DieWild, ID: "wild"}
	DieExtra    = Die{Symbol: "+", Color: colors.DieExtra, ID: "extra"}
)

func All() []Die {
	return []Die{DieDamage, DieShield, DieThruster, DieReactor, DieWild, DieExtra}
}

func (d *Die) Render(style lipgloss.Style) string {
	return style.
		Height(1).
		Border(lipgloss.RoundedBorder()).
		Padding(0, 1).
		Align(lipgloss.Center).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(d.Color).
		Foreground(d.Color).
		Render(d.Symbol)
}
