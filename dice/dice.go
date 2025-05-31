package dice

import (
	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/charmbracelet/lipgloss"
)

type Die struct {
	Color  lipgloss.Color
	Name   string
	Symbol string
}

var (
	DieDamage   = Die{Symbol: "X", Color: colors.DieDamage, Name: "damage"}
	DieShield   = Die{Symbol: "#", Color: colors.DieShield, Name: "shield"}
	DieThruster = Die{Symbol: "↟", Color: colors.DieThruster, Name: "thruster"}
	DieReactor  = Die{Symbol: "@", Color: colors.DieReactor, Name: "reactor"}
	DieWild     = Die{Symbol: "%", Color: colors.DieWild, Name: "wild"}
	DieExtra    = Die{Symbol: "+", Color: colors.DieExtra, Name: "extra"}
)

func All() []Die {
	return []Die{DieDamage, DieShield, DieThruster, DieReactor, DieWild, DieExtra}
}
