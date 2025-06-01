package games

import (
	"errors"

	"github.com/ascii-arcade/moonrollers/factions"
)

func (s *Game) SetFaction(player *Player, faction *factions.Faction) error {
	return s.withLock(func() error {
		if faction == nil {
			return errors.New("Faction cannot be nil")
		}

		player.Faction = faction
		return nil
	})
}

func (s *Game) AddPoints(pName string, amount int) error {
	return s.withLock(func() error {
		player, exists := s.getPlayer(pName)
		if !exists {
			return errors.New("Player not found")
		}
		player.incrementPoints(amount)
		return nil
	})
}
