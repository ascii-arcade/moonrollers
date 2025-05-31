package games

import "github.com/ascii-arcade/moonrollers/factions"

func (s *Game) SetFaction(player *Player, faction *factions.Faction) {
	s.withLock(func() {
		if faction == nil {
			return
		}
		player.Faction = faction
	})
}

func (s *Game) AddPoints(pName string, amount int) {
	s.withLock(func() {
		if player, exists := s.getPlayer(pName); exists {
			player.incrementPoints(amount)
		}
	})
}
