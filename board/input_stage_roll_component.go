package board

import (
	"fmt"
	"strings"
	"time"

	"github.com/ascii-arcade/moonrollers/keys"
	tea "github.com/charmbracelet/bubbletea"
)

type inputStageRollComponent struct {
	model *Model
}

func newInputStageRollComponent(model *Model) inputStageRollComponent {
	return inputStageRollComponent{
		model: model,
	}
}

func (c inputStageRollComponent) render() string {
	if c.model.Game.GetCurrentPlayer().Name != c.model.Player.Name {
		return inputComponentStyle(false).Render(fmt.Sprintf("Waiting for %s to roll...\n", c.model.Game.GetCurrentPlayer().Name))
	}

	var output strings.Builder
	if c.model.Game.RollCount > 0 {
		fmt.Fprintf(&output, "Press %s to end turn\n", keys.GameEndTurn.String(c.model.style))
	}
	return inputComponentStyle(true).Render(
		output.String(),
		fmt.Sprintf("Press %s to roll!", keys.GameRollDice.String(c.model.style)),
	)
}

func inputStageRollHandler(s *tableScreen, msg tea.KeyMsg) (*Model, tea.Cmd) {
	game := s.model.Game

	switch {
	case keys.GameRollDice.TriggeredBy(msg.String()):
		if game.InputObjective != nil && game.InputObjective.IsCompleted() {
			game.InputObjective = nil
		}
		if !s.isRolling {
			s.rollTickCount = 0
			s.isRolling = true
			return s.model, tea.Tick(rollInterval, func(time.Time) tea.Msg {
				return rollMsg{}
			})
		}
	case keys.GameEndTurn.TriggeredBy(msg.String()):
		game.NextTurn(false)
	}

	return s.model, nil
}
