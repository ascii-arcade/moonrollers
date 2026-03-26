package board

import (
	"github.com/ascii-arcade/moonrollers/messages"
	"github.com/ascii-arcade/moonrollers/rules"
)

func (s *Game) NextTurn(busted bool) {
	s.withLock(func() {
		s.CurrentTurnIndex++
		if len(s.players) <= s.CurrentTurnIndex {
			s.CurrentTurnIndex = 0
		}

		if s.isEndGame() {
			for _, player := range s.players {
				player.update(messages.WinnerScreen)
			}
			return
		}

		if !busted {
			for _, crew := range s.CrewForHire {
				if crew.ID == s.InputCrew.ID {
					for _, inputObjective := range s.InputCrew.Objectives {
						if !inputObjective.IsCompleted() {
							inputObjective.StartedBy = ""
							inputObjective.StartedByColor = ""
							inputObjective.CompletedAmount = 0
						}
					}
					crew.Objectives = s.InputCrew.Objectives
					break
				}
			}
		}

		startTurn := rules.NewStartTurn(s.players[s.CurrentTurnIndex].CrewIDs())
		s.InputState = InputStateRoll
		s.RollCount = 0
		s.InputCrew = nil
		s.InputObjective = nil

		s.initRollingPools(startTurn.RollingPoolSize)
	})
}

func (s *Game) initRollingPools(rollingPoolSize int) {
	s.RollingPool = NewDicePool(rollingPoolSize)
	s.SupplyPool = NewDicePool(12 - rollingPoolSize)
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
