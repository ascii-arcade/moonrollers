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
	game := c.model.Game
	style := c.model.style

	if game.GetCurrentPlayer().Name != c.model.Player.Name {
		return inputComponentStyle(false).Render(fmt.Sprintf("%s is committing dice...\n", game.GetCurrentPlayer().Name))
	}

	var output strings.Builder
	output.WriteString(style.Bold(true).Foreground(game.InputCrew.Faction.Color).Render(game.InputCrew.Name))
	output.WriteString("\n")
	output.WriteString(game.InputObjective.RenderCommitting(style))
	output.WriteString("\n")
	output.WriteString("How many dice do you want to commit?\n")

	containerStyle := style.
		Align(lipgloss.Center)

	diceCount := game.InputObjective.Committing.Length()

	topDice := make([]string, 0)
	bottomDice := make([]string, 0)
	for i := range diceCount {
		if i <= 3 {
			topDice = append(topDice, game.InputObjective.Type.Render(style))
			continue
		}

		bottomDice = append(bottomDice, game.InputObjective.Type.Render(style))
	}

	output.WriteString(containerStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			lipgloss.JoinHorizontal(lipgloss.Top, topDice...),
			lipgloss.JoinHorizontal(lipgloss.Top, bottomDice...),
		),
	))

	fmt.Fprintf(&output, "\n%s to commit a die", keys.GameCommitDie.String(style))
	fmt.Fprintf(&output, "\n%s to remove a die", keys.GameUncommitDie.String(style))
	fmt.Fprintf(&output, "\n%s to commit a special die", keys.GameCommitSpecialDie.String(style))
	fmt.Fprintf(&output, "\n%s to remove a special die", keys.GameUncommitSpecialDie.String(style))
	fmt.Fprintf(&output, "\n%s to confirm", keys.GameChooseConfirm.String(style))
	fmt.Fprintf(&output, "\n%s to end turn", keys.GameEndTurn.String(style))
	fmt.Fprintf(&output, "\n%s to go back", keys.GamePreviousInputStage.String(style))

	return inputComponentStyle(false).Render(output.String())
}
