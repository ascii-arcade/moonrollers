package games

import (
	"github.com/ascii-arcade/moonrollers/dice"
	"github.com/ascii-arcade/moonrollers/messages"
	"github.com/ascii-arcade/moonrollers/rules"
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

		startTurn := rules.NewStartTurn(s.players[s.CurrentTurnIndex].crewIDs())
		s.InputState = InputStateRoll
		s.RollCount = 0
		s.InputCrew = nil
		s.InputObjective = nil

		s.initRollingPools(startTurn.RollingPoolSize)
	})
}

func (s *Game) initRollingPools(rollingPoolSize int) {
	s.RollingPool = dice.NewDicePool(rollingPoolSize)
	s.SupplyPool = dice.NewDicePool(12 - rollingPoolSize)
}

func (s *Game) isEndGame() bool {
	if len(s.Deck) == 0 {
		return true
	}

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
