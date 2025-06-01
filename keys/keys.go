package keys

import (
	"slices"

	"github.com/charmbracelet/lipgloss"
)

type Keys []string

func (k Keys) TriggeredBy(msg string) bool {
	return slices.Contains(k, msg)
}

func (k Keys) String(style lipgloss.Style) string {
	if len(k) == 0 {
		return ""
	}
	return style.Bold(true).Italic(true).Render("'" + k[0] + "'")
}

var (
	ExitApplication = Keys{"ctrl+c"}

	MenuJoinGame     = Keys{"j"}
	MenuStartNewGame = Keys{"n"}

	PreviousScreen = Keys{"esc"}
	Submit         = Keys{"enter"}

	LobbyStartGame   = Keys{"s"}
	LobbyJoinFaction = Keys{"1", "2", "3", "4", "5", "6"}

	GameIncrementPoint = Keys{"a"}
)
