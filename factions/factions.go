package factions

import (
	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/charmbracelet/lipgloss"
)

type Faction struct {
	Name  string
	Color lipgloss.Color
}

func All() []Faction {
	return []Faction{
		{"blue", colors.PlayerBlue},
		{"green", colors.PlayerGreen},
		{"orange", colors.PlayerOrange},
		{"purple", colors.PlayerPurple},
		{"yellow", colors.PlayerYellow},
	}
}
