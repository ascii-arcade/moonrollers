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
		return containerStyle.Render("")
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

func (dp DicePool) Length() int {
	return len(dp.Dice)
}

func (dp *DicePool) Add(die Die) {
	dp.Dice = append(dp.Dice, die)
}

func (dp *DicePool) AddUnrolled(count int) {
	for range count {
		dp.Add(DieUnrolled)
	}
}

func (dp *DicePool) RemoveExtra() {
	removedOne := false
	dp.Dice = slices.DeleteFunc(dp.Dice, func(d Die) bool {
		if d.ID == DieUnrolled.ID && !removedOne {
			removedOne = true
			return true
		}
		return false
	})
}

func (dp *DicePool) Remove(die Die, count int) {
	i := 0
	dp.Dice = slices.DeleteFunc(dp.Dice, func(d Die) bool {
		if d == die || d == DieWild {
			i++
		}
		return i <= count && (d == die || d == DieWild)
	})
}

func (dp DicePool) NumberOf(die Die) int {
	count := 0
	for _, d := range dp.Dice {
		if d.Mimics != nil && d.Mimics.ID == die.ID {
			count += d.Value
			continue
		}
		if d.ID == die.ID || d.ID == DieWild.ID {
			count++
		}
	}
	return count
}

func (dp DicePool) HasType(die Die) bool {
	for _, d := range dp.Dice {
		if (d.ID == die.ID || d.Mimics != nil && d.Mimics.ID == die.ID) && d.ID != DieUnrolled.ID {
			return true
		}
	}
	return false
}

func (dp DicePool) HasExtra() bool {
	return dp.HasType(DieExtra)
}

func (dp *DicePool) Clear() {
	dp.Dice = []Die{}
}
