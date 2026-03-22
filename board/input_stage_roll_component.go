package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
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
