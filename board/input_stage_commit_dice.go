package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
	"github.com/charmbracelet/lipgloss"
)

type inputStageCommitDiceComponent struct {
	model *Model
}

func newInputStageCommitDiceComponent(model *Model) inputStageCommitDiceComponent {
	return inputStageCommitDiceComponent{
		model: model,
	}
}

func (c inputStageCommitDiceComponent) render() string {
	var output strings.Builder
	output.WriteString(c.model.style.Bold(true).Foreground(c.model.Game.InputCrew.Faction.Color).Render(c.model.Game.InputCrew.Name))
	output.WriteString("\n")
	output.WriteString(c.model.Game.InputObjective.RenderCommitting(c.model.style))
	output.WriteString("\n")
	output.WriteString("How many dice do you want to commit?\n")

	containerStyle := c.model.style.
		Align(lipgloss.Center)

	diceCount := c.model.Game.InputObjective.CommittingAmount

	topDice := make([]string, 0)
	bottomDice := make([]string, 0)
	for i := range diceCount {
		if i <= 3 {
			topDice = append(topDice, c.model.Game.InputObjective.Type.Render(c.model.style))
			continue
		}

		bottomDice = append(bottomDice, c.model.Game.InputObjective.Type.Render(c.model.style))
	}

	output.WriteString(containerStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			lipgloss.JoinHorizontal(lipgloss.Top, topDice...),
			lipgloss.JoinHorizontal(lipgloss.Top, bottomDice...),
		),
	))

	fmt.Fprintf(&output, "\n'1-%d' to commit a die", min(c.model.Game.RollingPool.NumberOf(c.model.Game.InputObjective.Type), c.model.Game.InputObjective.Amount-c.model.Game.InputObjective.CompletedAmount))
	fmt.Fprintf(&output, "\n%s to remove a die", keys.GameUncommitDie.String(c.model.style))
	fmt.Fprintf(&output, "\n%s to confirm", keys.GameChooseConfirm.String(c.model.style))
	fmt.Fprintf(&output, "\n%s to end turn", keys.GameEndTurn.String(c.model.style))
	fmt.Fprintf(&output, "\n%s to go back", keys.GamePreviousInputStage.String(c.model.style))

	return inputComponentStyle(false).Render(output.String())
}
