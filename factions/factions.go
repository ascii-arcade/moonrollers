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
	Blue   = Faction{Name: "Komek", Color: colors.DieReactor}
	Green  = Faction{Name: "Henko", Color: colors.DieShield}
	Orange = Faction{Name: "Magnomi", Color: colors.DieDamage}
	Purple = Faction{Name: "Sorelia", Color: colors.DieExtra}
	Yellow = Faction{Name: "Ventus", Color: colors.DieThruster}
)

func All() []Faction {
	return []Faction{Blue, Green, Orange, Purple, Yellow}
}
