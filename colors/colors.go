package colors

import "github.com/charmbracelet/lipgloss"

const (
	Border = lipgloss.Color("#404040")
	Error  = lipgloss.Color("#ff5555")
	Hazard = lipgloss.Color("#ff0000")
	Logo   = lipgloss.Color("#eb5761")

	InputStageBorder = lipgloss.Color("#7f8c8d")

	DieUnrolled           = lipgloss.Color("#009688")
	DieReactor            = lipgloss.Color("#2677fe")
	DieReactorUnselected  = lipgloss.Color("#133a7d")
	DieShield             = lipgloss.Color("#23741e")
	DieShieldUnselected   = lipgloss.Color("#10340e")
	DieDamage             = lipgloss.Color("#CC5500")
	DieDamageUnselected   = lipgloss.Color("#481f01")
	DieExtra              = lipgloss.Color("#ab19a6")
	DieExtraUnselected    = lipgloss.Color("#3f093d")
	DieThruster           = lipgloss.Color("#fefe26")
	DieThrusterUnselected = lipgloss.Color("#69690f")
	DieWild               = lipgloss.Color("#f0f0f0")
	DieWildUnselected     = lipgloss.Color("#777777")
)

func GetDark(dieType string) lipgloss.Color {
	switch dieType {
	case "reactor":
		return DieReactorUnselected
	case "shield":
		return DieShieldUnselected
	case "damage":
		return DieDamageUnselected
	case "extra":
		return DieExtraUnselected
	case "thruster":
		return DieThrusterUnselected
	case "wild":
		return DieWildUnselected
	default:
		return DieUnrolled
	}
}
