package factions

import "github.com/charmbracelet/lipgloss"

type Faction struct {
	Name  string
	Color lipgloss.Color
}

func All() []Faction {
	return []Faction{
		{"blue", lipgloss.Color("#2677fe")},
		{"yellow", lipgloss.Color("#fefe26")},
		{"orange", lipgloss.Color("#CC5500")},
		{"green", lipgloss.Color("#23741e")},
		{"purple", lipgloss.Color("#ab19a6")},
	}
}
