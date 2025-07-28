package board

import (
	"fmt"
	"time"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/ascii-arcade/moonrollers/keys"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type inputStageRollComponent struct {
	model *Model
}

func newInputStageRollComponent(model *Model) inputStageRollComponent {
	return inputStageRollComponent{
		model: model,
	}
}

func (c inputStageRollComponent) render() string {
	containerStyle := c.model.style.
		Width(30).
		Height(14).
		Padding(1).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(colors.InputStageBorder)

	return containerStyle.Render(
		fmt.Sprintf("Press %s to roll!", keys.GameRollDice.String(c.model.style)),
	)
}

func (s inputStageRollComponent) update(msg tea.Msg) (any, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if keys.GameRollDice.TriggeredBy(msg.String()) {
			if !s.model.Game.IsRolling {
				s.model.Game.RollTick = 0
				s.model.Game.IsRolling = true
				return s.model, tea.Tick(rollInterval, func(time.Time) tea.Msg {
					return rollMsg{}
				})
			}
		}
	}

	return s.model, nil
}
