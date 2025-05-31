package games

import (
	"github.com/ascii-arcade/moonrollers/factions"
	"github.com/ascii-arcade/moonrollers/generaterandom"
	"github.com/charmbracelet/lipgloss"
)

const defaultPlayerColor = lipgloss.Color("#FFFFFF")

type Player struct {
	Name      string
	Faction   *factions.Faction
	Points    int
	TurnOrder int

	isHost bool

	UpdateChan chan struct{}
}

func newPlayer(maxTurnOrder int, host bool) *Player {
	return &Player{
		Name:       generaterandom.Name(),
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
