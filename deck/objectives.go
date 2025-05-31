package deck

import "github.com/ascii-arcade/moonrollers/dice"

type objective struct {
	Type   dice.Die
	Amount int
	Hazard bool
}
