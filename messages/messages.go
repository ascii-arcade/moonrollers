package messages

import (
	"github.com/ascii-arcade/moonrollers/games"
	"github.com/ascii-arcade/moonrollers/screen"
)

type (
	SwitchToMenuMsg  struct{}
	SwitchToBoardMsg struct{ Game *games.Game }
	SwitchScreenMsg  struct{ Screen screen.Screen }
	RefreshBoard     struct{}
)
