package games

import (
	"errors"

	"github.com/ascii-arcade/moonrollers/factions"
)

func (s *Game) SetFaction(player *Player, faction *factions.Faction) error {
	return s.withLock(func() error {
		if faction == nil {
			return errors.New("error.faction_cannot_be_nil")
		}

		player.Faction = faction
		return nil
	})
}

func (s *Game) AddPoints(pName string, amount int) error {
	return s.withLock(func() error {
		player, exists := s.getPlayer(pName)
		if !exists {
			return errors.New("error.player_not_found")
		}
		player.incrementPoints(amount)
		return nil
	})
}
