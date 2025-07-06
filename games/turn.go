package games

func (s *Game) IsEndGame() bool {
	for _, player := range s.players {
		oneOfEach := true
		for _, count := range player.CrewCount {
			if count == 0 {
				oneOfEach = false
			}
			if count >= s.Settings.CardsOfAFactionToWin {
				return true
			}
		}
		if oneOfEach {
			return true
		}
	}
	return false
}
