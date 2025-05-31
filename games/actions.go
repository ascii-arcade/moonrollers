package games

func (s *Game) AddPoints(pName string, amount int) {
	s.withLock(func() {
		if player, exists := s.getPlayer(pName); exists {
			player.incrementPoints(amount)
		}
	})
}
