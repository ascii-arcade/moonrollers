package games

import (
	"github.com/ascii-arcade/moonrollers/factions"
	"github.com/ascii-arcade/moonrollers/generaterandom"
	"github.com/ascii-arcade/moonrollers/language"
)

type Player struct {
	Name      string
	Faction   *factions.Faction
	Points    int
	TurnOrder int

	isHost bool

	UpdateChan chan struct{}
}

func newPlayer(maxTurnOrder int, host bool, lang *language.Language) *Player {
	return &Player{
		Name:       generaterandom.Name(lang),
		Points:     0,
		TurnOrder:  maxTurnOrder + 1,
		UpdateChan: make(chan struct{}),
		isHost:     host,
	}
}

func (p *Player) IsHost() bool {
	return p.isHost
}

func (p *Player) incrementPoints(amount int) {
	p.Points += amount
}

func (p *Player) HasFaction() bool {
	return p.Faction != nil
}
