package games

import (
	"github.com/ascii-arcade/moonrollers/dice"
	"github.com/ascii-arcade/moonrollers/messages"
)

func (s *Game) NextTurn() {
	s.withLock(func() {
		if len(s.players) > s.CurrentTurnIndex+1 {
			s.CurrentTurnIndex++
		} else {
			s.CurrentTurnIndex = 0
		}

		if s.isEndGame() {
			for _, player := range s.players {
				player.update(messages.WinnerScreen)
			}
			return
		}

		s.initRollingPools()
	})
}

func (s *Game) initRollingPools() {
	s.RollingPool = dice.NewDicePool(5)
	s.SupplyPool = dice.NewDicePool(7)
}

func (s *Game) isEndGame() bool {
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
