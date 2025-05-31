package deck

import (
	"math/rand/v2"
)

type Deck []*Crew

func NewDeck() Deck {
	copiedDeck := make(Deck, len(allCrew))
	for i, c := range allCrew {
		copiedDeck[i] = &Crew{
			Ability:    c.Ability,
			Faction:    c.Faction,
			IsStarter:  c.IsStarter,
			Name:       c.Name,
			Objectives: c.Objectives,
		}
	}
	copiedDeck.Shuffle()

	return copiedDeck
}

func (d *Deck) Shuffle() {
	for i := len(*d) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	}
}
