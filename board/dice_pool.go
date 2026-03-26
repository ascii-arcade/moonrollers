package board

import (
	"math/rand"
	"slices"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/charmbracelet/lipgloss"
)

type DicePool struct {
	Dice []*Die
}

func NewDicePool(size int) DicePool {
	dice := make([]*Die, 0)
	for range size {
		dice = append(dice, new(DieUnrolled))
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
		topDice = append(topDice, dp.Dice[i].Render(style, false))
	}

	bottomDice := make([]string, 0)
	for i := range bottomCount {
		bottomDice = append(bottomDice, dp.Dice[i+topCount].Render(style, false))
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
	all := AllDie()
	for i := range dp.Dice {
		dp.Dice[i] = new(all[rand.Intn(len(all))])
	}
}

func (dp DicePool) Length() int {
	return len(dp.Dice)
}

func (dp *DicePool) Add(die *Die) {
	dp.Dice = append(dp.Dice, die)
}

func (dp *DicePool) AddUnrolled() {
	dp.Add(new(DieUnrolled))
}

func (dp *DicePool) RemoveUnrolled() {
	removedOne := false
	dp.Dice = slices.DeleteFunc(dp.Dice, func(d *Die) bool {
		if d.ID == DieUnrolled.ID && !removedOne {
			removedOne = true
			return true
		}
		return false
	})
}

func (dp *DicePool) RemoveCommitted() {
	dp.Dice = slices.DeleteFunc(dp.Dice, func(d *Die) bool {
		return d.Selected
	})
}

func (dp *DicePool) NumberOfSelected() int {
	count := 0
	for _, d := range dp.Dice {
		if d.Selected {
			count += d.Value
		}
	}
	return count
}

func (dp DicePool) ValueOf(dieType string) int {
	count := 0
	for _, d := range dp.Dice {
		if d.ID == dieType || d.ID == DieWild.ID || d.MimicsType(dieType) {
			count += d.Value
		}
	}
	return count
}

func (dp DicePool) NumberOf(dieType string) int {
	count := 0
	for _, d := range dp.Dice {
		if d.ID == dieType {
			count++
		}
	}
	return count
}

func (dp DicePool) NumberOfUnrolled() int {
	count := 0
	for _, d := range dp.Dice {
		if d.ID == DieUnrolled.ID {
			count++
		}
	}
	return count
}

func (dp DicePool) NumberOfExtra() int {
	count := 0
	for _, d := range dp.Dice {
		if d.ID == DieExtra.ID {
			count++
		}
	}
	return count
}

func (dp DicePool) HasType(dieType string) bool {
	for _, d := range dp.Dice {
		if (d.ID == dieType || d.MimicsType(dieType)) && d.ID != DieUnrolled.ID {
			return true
		}
	}
	return false
}

func (dp DicePool) HasExtra() bool {
	return dp.HasType(DieExtra.ID)
}

func (dp *DicePool) Clear() {
	dp.Dice = []*Die{}
}

func (dp *DicePool) Contains(die *Die) bool {
	return slices.Contains(dp.Dice, die)
}

func (dp *DicePool) GetValidDice(dieType string) []*Die {
	valid := make([]*Die, 0)
	for _, die := range dp.Dice {
		if die.ID == dieType || die.MimicsType(dieType) || die.ID == DieWild.ID {
			valid = append(valid, die)
		}
	}
	return valid
}
