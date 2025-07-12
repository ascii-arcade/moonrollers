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
	Crew               map[string]*deck.Crew
	CrewCount          map[string]int
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

func (p *Player) MakeHost() *Player {
	p.isHost = true
	return p
}

func (p *Player) IsHost() bool {
	return p.isHost
}

func (p *Player) HasFaction() bool {
	return p.Faction != nil
}

func (p *Player) OnDisconnect(fn func()) {
	p.onDisconnect = append(p.onDisconnect, fn)
}

func (p *Player) AddCrew(crew *deck.Crew, active bool) {
	if active {
		p.Crew[crew.Faction.Name] = crew
	}
	p.CrewCount[crew.Faction.Name]++
}

func (p *Player) crewIDs() []string {
	ids := make([]string, 0)
	for _, crew := range p.Crew {
		ids = append(ids, crew.ID)
	}
	return ids
}

func (p *Player) update(code int) {
	select {
	case p.UpdateChan <- code:
	default:
	}
}
