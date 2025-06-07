package rules

import (
	"errors"

	"slices"

	"github.com/ascii-arcade/moonrollers/dice"
)

func IsUsable(crew []string, rolledDie *dice.Die, objectiveType *dice.Die) (bool, int, error) {
	if crew == nil || rolledDie == nil || objectiveType == nil {
		return false, 0, errors.New("crew, rolledDie, and objectiveType must not be nil")
	}

	if slices.Contains(crew, "dana") && rolledDie.ID == dice.DieExtra.ID && objectiveType.ID == dice.DieDamage.ID {
		return true, 2, nil
	}

	if rolledDie.ID == objectiveType.ID {
		return true, 1, nil
	}
	return false, 0, nil
}
