package games

type PlayerState struct {
	turnOrder int
}

func (p *PlayerState) setTurnOrder(order int) {
	p.turnOrder = order
}
