package board

import "github.com/ascii-arcade/moonrollers/dice"

type diceComponent struct {
	model    *Model
	dicePool dice.DicePool
}

func newDiceComponent(model *Model, dp dice.DicePool) diceComponent {
	return diceComponent{
		model:    model,
		dicePool: dp,
	}
}

func (d *diceComponent) render() string {
	return d.dicePool.Render(d.model.style)
}
