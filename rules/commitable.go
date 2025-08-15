package rules

import (
	"slices"

	"github.com/ascii-arcade/moonrollers/deck"
	"github.com/ascii-arcade/moonrollers/dice"
)

type CommitableDie struct {
	dice.Die
	amount int
}

func CommitableFor(handCrewIDs []string, target *dice.Die, rolled dice.DicePool) []CommitableDie {
	commitableDice := make([]CommitableDie, 0)

	for _, die := range rolled.Dice {
		if die.ID == target.ID {
			commitableDie := CommitableDie{Die: die, amount: 1}
			commitableDice = append(commitableDice, commitableDie)
		}
	}

	return commitableDice
}

func CommitableToCrew(handCrewIDs []string, crew []*deck.Crew, rolled dice.DicePool) []*deck.Crew {
	commitableCrew := make([]*deck.Crew, 0)
	for _, c := range crew {
		if slices.Contains(commitableCrew, c) {
			continue
		}
		for _, objective := range c.AvailableObjectives() {
			if len(CommitableFor(handCrewIDs, &objective.Type, rolled)) > 0 {
				commitableCrew = append(commitableCrew, c)
				break
			}
		}
	}
	return commitableCrew
}
