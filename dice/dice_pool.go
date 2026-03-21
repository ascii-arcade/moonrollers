package dice

import (
	"math/rand"
	"slices"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/charmbracelet/lipgloss"
)

type DicePool struct {
	Dice []Die
}

func NewDicePool(size int) DicePool {
	dice := make([]Die, 0)
	for range size {
		dice = append(dice, DieUnrolled)
	}
	return DicePool{
		Dice: dice,
	}
}

func (dp *DicePool) Render(style lipgloss.Style) string {
	containerStyle := style.
		Width(30).
		Height(6).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(colors.DieUnrolled)

	diceCount := len(dp.Dice)
	if diceCount == 0 {
		return ""
	}
	topCount := (diceCount + 1) / 2
	bottomCount := diceCount / 2

	topDice := make([]string, 0)
	for i := range topCount {
		topDice = append(topDice, dp.Dice[i].Render(style))
	}

	bottomDice := make([]string, 0)
	for i := range bottomCount {
		bottomDice = append(bottomDice, dp.Dice[i+topCount].Render(style))
	}

	return containerStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			lipgloss.JoinHorizontal(lipgloss.Top, topDice...),
			lipgloss.JoinHorizontal(lipgloss.Top, bottomDice...),
		),
	)
}

func (dp DicePool) Roll() {
	all := All()
	for i := range dp.Dice {
		dp.Dice[i] = all[rand.Intn(len(all))]
	}
}

func (dp *DicePool) Remove(die Die, count int) {
	i := 0
	dp.Dice = slices.DeleteFunc(dp.Dice, func(d Die) bool {
		if d == die {
			i++
		}
		return i <= count && d == die
	})
}

func (dp DicePool) NumberOf(die Die) int {
	count := 0
	for _, d := range dp.Dice {
		if d == die {
			count++
		}
	}
	return count
}
