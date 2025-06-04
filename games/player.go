package games

import (
	"context"

	"github.com/ascii-arcade/moonrollers/factions"
	"github.com/ascii-arcade/moonrollers/generaterandom"
	"github.com/ascii-arcade/moonrollers/language"
)

type Player struct {
	Name               string
	Faction            *factions.Faction
	Points             int
	TurnOrder          int
	LanguagePreference *language.LanguagePreference

	UpdateChan chan struct{}

	isHost bool
	ctx    context.Context
}

func NewPlayer(ctx context.Context, lang *language.LanguagePreference) *Player {
	return &Player{
		Name:               generaterandom.Name(lang.Lang),
		Points:             0,
		UpdateChan:         make(chan struct{}),
		LanguagePreference: lang,
		ctx:                ctx,
	}
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
