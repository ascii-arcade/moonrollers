package games

import (
	"slices"

	"github.com/ascii-arcade/moonrollers/deck"
	"github.com/ascii-arcade/moonrollers/factions"
)

func (s *Game) Begin() {
	s.withLock(func() {
		s.Deck = deck.NewDeck()
		s.dealCrewForHire()
		s.inProgress = true
	})
}

func (s *Game) dealCrewForHire() {
	var draw int
	switch len(s.players) {
	case 2:
		draw = 4
	case 3:
		draw = 5
	default:
		draw = 6
	}

	skippedCrew := make([]*deck.Crew, 0)
	for len(s.CrewForHire) < draw {
		card := s.Deck[0]
		if s.hasFactionForHire(card.Faction) && len(s.CrewForHire) < 5 {
			skippedCrew = append(skippedCrew, card)
			s.Deck = slices.Delete(s.Deck, 0, 1)
			continue
		}
		s.CrewForHire = append(s.CrewForHire, card)
		s.Deck = slices.Delete(s.Deck, 0, 1)
	}

	s.Deck = append(s.Deck, skippedCrew...)
	s.Deck.Shuffle()
}

func (s *Game) hasFactionForHire(faction factions.Faction) bool {
	for _, crew := range s.CrewForHire {
		if crew.Faction == faction {
			return true
		}
	}
	return false
}
