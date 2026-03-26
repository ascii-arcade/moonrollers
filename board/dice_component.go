package board

import (
	"time"
)

type diceComponent struct {
	model    *Model
	dicePool DicePool
}

const (
	rollFrames   = 7
	rollInterval = 200 * time.Millisecond
)

func newDiceComponent(model *Model, dp DicePool) diceComponent {
	return diceComponent{
		model:    model,
		dicePool: dp,
	}
}

func (d *diceComponent) render() string {
	return d.dicePool.Render(d.model.style)
}
