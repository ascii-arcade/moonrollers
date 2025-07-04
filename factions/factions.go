package factions

import (
	"strings"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/charmbracelet/lipgloss"
)

type Faction struct {
	Name      string
	Icon      string
	Color     lipgloss.Color
	SortOrder int
}

var icons = map[string]string{
	"magnomi": strings.Join([]string{
		"█▀▀▀▀▀█",
		"█ ███ █",
		"█ ▀▀▀ █",
		"█ ███ █",
		"▀▀▀▀▀▀▀",
	}, "\n"),

	"ventus": strings.Join([]string{
		"▓▓▓▓▓▓▓▓",
		"▓ ██ ▓▓",
		"▓▓▓▓▓▓▓▓",
		"▓ ▓▓ ██",
		"▓▓▓▓▓▓▓▓",
	}, "\n"),

	"komek": strings.Join([]string{
		"████████",
		"█ ██ ███",
		"█ ██ ███",
		"█ ▓▓ ▓▓█",
		"████████",
	}, "\n"),

	"henko": strings.Join([]string{
		"▓▓▓▓▓▓▓▓",
		"▓ ██ ██▓",
		"▓ ▓▓▓▓ ▓",
		"▓ ██ ██▓",
		"▓▓▓▓▓▓▓▓",
	}, "\n"),

	"sorelia": strings.Join([]string{
		"▓▀▀░░▀▀▓",
		"█ █░░█ █",
		"█░████░█",
		"█ █░░█ █",
		"▓▄▄▄▄▄▄▓",
	}, "\n"),
}

var (
	Blue   = Faction{Name: "Komek", Color: colors.DieReactor, Icon: icons["komek"], SortOrder: 1}
	Green  = Faction{Name: "Henko", Color: colors.DieShield, Icon: icons["henko"], SortOrder: 2}
	Orange = Faction{Name: "Magnomi", Color: colors.DieDamage, Icon: icons["magnomi"], SortOrder: 3}
	Purple = Faction{Name: "Sorelia", Color: colors.DieExtra, Icon: icons["sorelia"], SortOrder: 4}
	Yellow = Faction{Name: "Ventus", Color: colors.DieThruster, Icon: icons["ventus"], SortOrder: 5}
)

func All() []Faction {
	return []Faction{Blue, Green, Orange, Purple, Yellow}
}
