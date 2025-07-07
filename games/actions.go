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

func (s *Game) Roll(isRolling bool) {
	s.withLock(func() {
		s.IsRolled = true
		s.RollingPool.Roll()
	})
}
