package games

import (
	"slices"

	"github.com/ascii-arcade/moonrollers/deck"
	"github.com/ascii-arcade/moonrollers/factions"
)

const (
	minimumPlayers = 2
	maximumPlayers = 5
)

func (s *Game) Begin() {
	s.withLock(func() {
		if _, ok := s.IsPlayerCountOk(); !ok {
			return
		}
		s.Deck = deck.NewDeck()
		s.dealCrewForHire()
		s.inProgress = true
	})
}

func (s *Game) dealCrewForHire() {
	draw := min(len(s.players)+2, 6)

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

func (s *Game) IsPlayerCountOk() (string, bool) {
	if len(s.players) > maximumPlayers {
		return "Too many players", false
	}
	if len(s.players) < minimumPlayers {
		return "Not enough players", false
	}
	return "", true
}
