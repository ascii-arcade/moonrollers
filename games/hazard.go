package games

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Hazard struct {
	Points  int
	Hazards int
	Count   int
}

var hazards = []Hazard{
	{Points: 1, Hazards: 0},
	{Points: 2, Hazards: 1},
	{Points: 5, Hazards: 2},
}

func (h *Hazard) Render(style lipgloss.Style) string {

	var hazardText strings.Builder
	hazardText.WriteString(strconv.Itoa(h.Points))
	if h.Hazards > 0 {
		for range h.Hazards {
			hazardText.WriteString("⚠")
		}
	}
	if h.Count > 1 {
		hazardText.WriteString(" x" + strconv.Itoa(h.Count))
	}

	return style.
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("red")).
		PaddingLeft(1).
		PaddingRight(1).
		Width(6).Render(hazardText.String())
}
