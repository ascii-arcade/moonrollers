package games

import (
	"errors"
	"slices"

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
		s.RollCount++
		s.RollingPool.Roll()
		if !isRolling {
			s.InputState = InputStateChooseCrew
		}
	})
}

func (s *Game) ChooseCrewMember(index int) {
	s.withLock(func() {
		if index < 0 || index >= len(s.CrewForHire) {
			return
		}

		s.InputCrew = s.CrewForHire[index]
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
		if index < 0 || index >= len(s.InputCrew.Objectives) {
			return
		}

		s.InputObjective = &s.InputCrew.Objectives[index]
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
