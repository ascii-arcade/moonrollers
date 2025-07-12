package games

type PlayerState struct {
	IsHost    bool
	TurnOrder int
}

func (p *PlayerState) setTurnOrder(order int) {
	p.TurnOrder = order
}
