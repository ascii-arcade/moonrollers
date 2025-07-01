package games

import (
	"context"

	"github.com/ascii-arcade/moonrollers/deck"
	"github.com/ascii-arcade/moonrollers/factions"
	"github.com/ascii-arcade/moonrollers/language"
	"github.com/charmbracelet/ssh"
)

type Player struct {
	Name               string
	Faction            *factions.Faction
	Points             int
	Crew               []*deck.Crew
	TurnOrder          int
	LanguagePreference *language.LanguagePreference

	UpdateChan chan int

	isHost    bool
	connected bool

	Sess ssh.Session

	onDisconnect []func()
	ctx          context.Context
}

func (p *Player) SetName(name string) *Player {
	p.Name = name
	return p
}

func (p *Player) SetTurnOrder(order int) *Player {
	p.TurnOrder = order
	return p
}

func (p *Player) MakeHost() *Player {
	p.isHost = true
	return p
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

func (p *Player) OnDisconnect(fn func()) {
	p.onDisconnect = append(p.onDisconnect, fn)
}

func (p *Player) update(code int) {
	select {
	case p.UpdateChan <- code:
	default:
	}
}
