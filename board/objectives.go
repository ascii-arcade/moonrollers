package board

import (
	"strconv"
	"strings"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/charmbracelet/lipgloss"
)

const (
	symbolFullPip  = "◆"
	symbolEmptyPip = "◇"
	symbolHazard   = "!"
)

type Objective struct {
	Type             *Die
	Amount           int
	Hazard           bool
	CommittingAmount int
	CompletedAmount  int
	StartedBy        string
	StartedByColor   lipgloss.Color
}

func (o *Objective) Points() int {
	if o.Type.ID == DieWild.ID {
		return 2 * o.Amount
	}
	return o.Amount
}

func (o *Objective) Render(style lipgloss.Style, selected bool) string {
	var line strings.Builder
	line.WriteString(style.Foreground(o.Type.Color).Render(o.Type.Symbol))
	arrow := " "
	if selected {
		arrow = ">"
	}
	line.WriteString(arrow)
	line.WriteString(o.getHazard(style))
	for range o.CompletedAmount {
		line.WriteString(style.Foreground(o.StartedByColor).Render(symbolFullPip))
	}
	for range o.Amount - o.CompletedAmount {
		line.WriteString(style.Foreground(o.StartedByColor).Render(symbolEmptyPip))
	}
	for range 5 - o.Amount {
		line.WriteString(" ")
	}
	line.WriteString(strconv.Itoa(o.Points()))
	return line.String()
}

func (o *Objective) RenderCommitting(dice DicePool, style lipgloss.Style) string {
	var line strings.Builder
	line.WriteString(style.Foreground(o.Type.Color).Render(o.Type.Symbol))
	line.WriteString(" ")
	line.WriteString(o.getHazard(style))
	for i := range o.CommittingAmount {
		if i < o.Amount {
			line.WriteString(style.Foreground(o.StartedByColor).Render(symbolFullPip))
		}
	}
	for range o.Amount - o.CommittingAmount {
		// if commiting > i {
		// 	line.WriteString(style.Foreground(o.StartedByColor).Render(symbolFullPip))
		// 	continue
		// }
		line.WriteString(style.Foreground(o.StartedByColor).Render(symbolEmptyPip))
	}
	for range 5 - o.Amount {
		line.WriteString(" ")
	}
	line.WriteString(strconv.Itoa(o.Points()))
	return line.String()
}

func (o *Objective) CanCommitBy(playerName string) bool {
	return o.StartedBy == "" || o.StartedBy == playerName
}

func (o *Objective) IsCompleted() bool {
	return o.CompletedAmount == o.Amount
}

func (o *Objective) IsType(dieType string) bool {
	return o.Type.ID == dieType
}

func (o *Objective) getHazard(style lipgloss.Style) string {
	if o.Hazard {
		return style.Foreground(colors.Hazard).Render(symbolHazard)
	}
	return " "
}

func (o *Objective) ValidDie(die Die) bool {
	return o.IsType(die.ID) || o.IsType(DieWild.ID) || o.IsType(die.Mimics)
}
