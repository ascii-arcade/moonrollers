package messages

import (
	"github.com/ascii-arcade/moonrollers/screen"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	SwitchViewMsg   struct{ Model tea.Model }
	SwitchScreenMsg struct{ Screen screen.Screen }
	NewGame         struct{}
	JoinGame        struct{ GameCode string }
	RefreshBoard    struct{}
)
