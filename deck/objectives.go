package deck

import "github.com/ascii-arcade/moonrollers/dice"

type objective struct {
	Type   dice.Die
	Amount int
	Hazard bool
}

func (o *objective) Points() int {
	if o.Type == dice.DieWild {
		return 2 * o.Amount
	}

	return o.Amount
}
