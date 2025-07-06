package board

import (
	"time"

	"github.com/ascii-arcade/moonrollers/dice"
)

type diceComponent struct {
	model    *Model
	dicePool dice.DicePool
}

const (
	rollFrames   = 15
	rollInterval = 200 * time.Millisecond
)

func newDiceComponent(model *Model, dp dice.DicePool) diceComponent {
	return diceComponent{
		model:    model,
		dicePool: dp,
	}
}

func (d *diceComponent) render() string {
	return d.dicePool.Render(d.model.style)
}
