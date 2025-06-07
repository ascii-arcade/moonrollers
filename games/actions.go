package games

import (
	"errors"

	"github.com/ascii-arcade/moonrollers/factions"
)

func (s *Game) SetFaction(player *Player, faction *factions.Faction) error {
	return s.withErrLock(func() error {
		if faction == nil {
			return errors.New("faction_cannot_be_nil")
		}

		player.Faction = faction
		return nil
	})
}

func (s *Game) AddPoints(player *Player, amount int) error {
	return s.withErrLock(func() error {
		player, exists := s.getPlayer(player.Sess)
		if !exists {
			return errors.New("player_not_found")
		}
		player.incrementPoints(amount)
		return nil
	})
}
