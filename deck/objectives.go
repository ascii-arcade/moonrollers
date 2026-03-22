package deck

import (
	"strconv"
	"strings"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/ascii-arcade/moonrollers/dice"
	"github.com/charmbracelet/lipgloss"
)

const (
	emptyPip = "◇"
	fullPip  = "◆"
	Hazard   = "!"
)

type Objective struct {
	Type             dice.Die
	Amount           int
	Hazard           bool
	CommittingAmount int
	CompletedAmount  int
	StartedBy        string
	StartedByColor   lipgloss.Color
}

func (o *Objective) Points() int {
	if o.Type == dice.DieWild {
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
		line.WriteString(style.Foreground(o.StartedByColor).Render(fullPip))
	}
	for range o.Amount - o.CompletedAmount {
		line.WriteString(style.Foreground(o.StartedByColor).Render(emptyPip))
	}
	for range 5 - o.Amount {
		line.WriteString(" ")
	}
	line.WriteString(strconv.Itoa(o.Points()))
	return line.String()
}

func (o *Objective) RenderCommitting(style lipgloss.Style) string {
	var line strings.Builder
	line.WriteString(style.Foreground(o.Type.Color).Render(o.Type.Symbol))
	line.WriteString(" ")
	line.WriteString(o.getHazard(style))
	for range o.CommittingAmount {
		line.WriteString(style.Foreground(o.StartedByColor).Render(fullPip))
	}
	for i := range o.Amount - o.CommittingAmount {
		if o.CompletedAmount > i {
			line.WriteString(style.Foreground(o.StartedByColor).Render(fullPip))
			continue
		}
		line.WriteString(style.Foreground(o.StartedByColor).Render(emptyPip))
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

func (o *Objective) IsType(die dice.Die) bool {
	return o.Type == die || die == dice.DieWild
}

func (o *Objective) getHazard(style lipgloss.Style) string {
	if o.Hazard {
		return style.Foreground(colors.Hazard).Render(Hazard)
	}
	return " "
}
