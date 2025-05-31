package factions

import (
	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/charmbracelet/lipgloss"
)

type Faction struct {
	Name  string
	Color lipgloss.Color
}

var (
	Blue   = Faction{Name: "blue", Color: colors.DieReactor}
	Green  = Faction{Name: "green", Color: colors.DieShield}
	Orange = Faction{Name: "orange", Color: colors.DieDamage}
	Purple = Faction{Name: "purple", Color: colors.DieExtra}
	Yellow = Faction{Name: "yellow", Color: colors.DieThruster}
)

func All() []Faction {
	return []Faction{Blue, Green, Orange, Purple, Yellow}
}
