package board

import (
	"fmt"

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
	return inputComponentStyle().Render(
		fmt.Sprintf("Press %s to roll!", keys.GameRollDice.String(c.model.style)),
	)
}
