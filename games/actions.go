package games

import (
	"errors"
	"slices"

	"github.com/ascii-arcade/moonrollers/dice"
	"github.com/ascii-arcade/moonrollers/factions"
)

func (s *Game) SetFaction(player *Player, faction *factions.Faction) error {
	return s.withErrLock(func() error {
		player.Faction = faction
		return nil
	})
}

func (s *Game) Roll(isRolling bool) {
	s.withLock(func() {
		s.RollCount++
		s.RollingPool.Roll()
		if !isRolling {
			switch {
			case s.InputCrew != nil && !s.InputCrew.CanCommit(s.RollingPool, s.GetCurrentPlayer().Sess.User()),
				s.InputObjective != nil && s.RollingPool.ValueOf(s.InputObjective.Type.ID) == 0:
				s.InputState = Busted
			case s.InputObjective != nil && s.InputObjective.CompletedAmount < s.InputObjective.Amount:
				s.InputState = InputStateCommitDice
			case s.InputCrew != nil && !s.InputCrew.IsComplete():
				s.InputState = InputStateChooseObjective
			default:
				s.InputState = InputStateChooseCrew
			}

			s.ApplyModifiers()

			slices.SortFunc(s.RollingPool.Dice, func(a, b *dice.Die) int {
				if a.ID == dice.DieExtra.ID && b.ID != dice.DieExtra.ID {
					return -1
				}
				if b.ID == dice.DieExtra.ID && a.ID != dice.DieExtra.ID {
					return 1
				}
				return 0
			})
		}
	})
}

func (s *Game) ApplyModifiers() {
	for _, c := range s.GetCurrentPlayer().Crew {
		if c.Modifier != nil {
			c.Modifier(&s.SupplyPool, &s.RollingPool)
		}
	}
}

func (s *Game) ChooseCrewMember(index int) {
	s.withLock(func() {
		if len(s.CrewForHire) <= index || !s.CrewForHire[index].CanCommit(s.RollingPool, s.GetCurrentPlayer().Name) {
			return
		}
		inputCrew := s.CrewForHire[index].Copy()
		s.InputCrew = &inputCrew
	})
}

func (s *Game) ConfirmCrewMember() {
	s.withLock(func() {
		if s.InputCrew == nil {
			return
		}
		s.InputState = InputStateChooseObjective
	})
}

func (s *Game) ChooseObjective(index int) {
	s.withLock(func() {
		switch {
		case index < 0,
			index >= len(s.InputCrew.Objectives),
			s.RollingPool.ValueOf(s.InputCrew.Objectives[index].Type.ID) == 0,
			s.InputCrew.Objectives[index].IsCompleted(),
			s.InputCrew.Objectives[index].StartedBy != "" && s.InputCrew.Objectives[index].StartedBy != s.GetCurrentPlayer().Name:
			return
		}

		s.InputObjective = s.InputCrew.Objectives[index]
	})
}

func (s *Game) ConfirmObjective() {
	s.withLock(func() {
		if s.InputObjective == nil {
			return
		}
		s.InputState = InputStateCommitDice
	})
}

func (s *Game) PreviousInputStage() {
	s.withLock(func() {
		switch s.InputState {
		case InputStateChooseObjective:
			s.InputObjective = nil
			s.InputState = InputStateChooseCrew
		case InputStateCommitDice:
			s.InputState = InputStateChooseObjective
		}
	})
}

func (s *Game) HireCrewMember(index int, player *Player) error {
	return s.withErrLock(func() error {
		if index < 0 || index >= len(s.CrewForHire) {
			return errors.New("invalid_crew_index")
		}

		crew := s.CrewForHire[index]
		if crew == nil {
			return errors.New("crew_not_found")
		}

		player.AddCrew(crew, true)
		s.CrewForHire = slices.Delete(s.CrewForHire, index, index+1)
		s.dealSingleCrew()
		return nil
	})
}

func (s *Game) dealSingleCrew() {
	if len(s.Deck) == 0 {
		return
	}
	s.CrewForHire = append(s.CrewForHire, s.Deck[0])
	s.Deck = slices.Delete(s.Deck, 0, 1)
}
