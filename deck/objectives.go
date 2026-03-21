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
}

func (o *Objective) Points() int {
	if o.Type == dice.DieWild {
		return 2 * o.Amount
	}

	return o.Amount
}

func (o *Objective) Render(style lipgloss.Style) string {
	var line strings.Builder
	line.WriteString(style.Foreground(o.Type.Color).Render(o.Type.Symbol))
	line.WriteString(" ")
	if o.Hazard {
		line.WriteString(style.Foreground(colors.Hazard).Render(Hazard))
	} else {
		line.WriteString(" ")
	}
	for range o.CompletedAmount {
		line.WriteString(fullPip)
	}
	for range o.Amount - o.CompletedAmount {
		line.WriteString(emptyPip)
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
	if o.Hazard {
		line.WriteString(style.Foreground(colors.Hazard).Render(Hazard))
	} else {
		line.WriteString(" ")
	}
	for range o.CommittingAmount {
		line.WriteString(fullPip)
	}
	for i := range o.Amount - o.CommittingAmount {
		if o.CompletedAmount > i {
			line.WriteString(fullPip)
			continue
		}
		line.WriteString(emptyPip)
	}
	for range 5 - o.Amount {
		line.WriteString(" ")
	}
	line.WriteString(strconv.Itoa(o.Points()))
	return line.String()
}
