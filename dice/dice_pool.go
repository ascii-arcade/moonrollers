package dice

import (
	"math/rand"

	"github.com/charmbracelet/lipgloss"
)

type DicePool []Die

func NewDicePool(size int) DicePool {
	dicePool := make(DicePool, size)
	for range size {
		dicePool = append(dicePool, DieUnrolled)
	}
	return dicePool
}

func (dp DicePool) Render(style lipgloss.Style) []string {
	rendered := make([]string, len(dp))
	for i, die := range dp {
		rendered[i] = die.Render(style)
	}
	return rendered
}

func (dp DicePool) Roll() {
	all := All()
	for i := range dp {
		dp[i] = all[rand.Intn(len(all))]
	}
}
