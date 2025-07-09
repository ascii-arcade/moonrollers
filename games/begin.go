package games

import (
	"errors"
	"slices"

	"github.com/ascii-arcade/moonrollers/deck"
	"github.com/ascii-arcade/moonrollers/factions"
	"github.com/ascii-arcade/moonrollers/messages"
	"github.com/ascii-arcade/moonrollers/rules"
)

const (
	minimumPlayers = 2
	maximumPlayers = 5
)

func (s *Game) Begin() error {
	return s.withErrLock(func() error {
		if err := s.IsPlayerCountOk(); err != nil {
			return err
		}
		s.Deck = deck.NewDeck()
		s.dealStarterCards()
		s.dealCrewForHire()
		s.CurrentTurnIndex = 0
		s.inProgress = true

		startTurn := rules.NewStartTurn(s.players[s.CurrentTurnIndex].crewIDs())
		s.initRollingPools(startTurn.RollingPoolSize)

		for _, p := range s.players {
			p.update(messages.TableScreen)
		}
		return nil
	})
}

func (s *Game) dealStarterCards() {
	if !s.Settings.UseStarterCards {
		return
	}

	for _, player := range s.players {
		for _, crew := range s.Deck {
			if crew.IsStarter && crew.Faction == *player.Faction {
				index := slices.Index(s.Deck, crew)
				s.Deck = slices.Delete(s.Deck, index, index+1)
				player.AddCrew(crew, true)
				break
			}
		}
	}
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

func (s *Game) IsPlayerCountOk() error {
	if len(s.players) > maximumPlayers {
		return errors.New("too_many_players")
	}
	if len(s.players) < minimumPlayers {
		return errors.New("not_enough_players")
	}
	return nil
}
