package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
	tea "github.com/charmbracelet/bubbletea"
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
		return inputComponentStyle(false).Render(fmt.Sprintf("%s is committing ..\n", game.GetCurrentPlayer().Name))
	}

	var output strings.Builder
	output.WriteString(style.Bold(true).Foreground(game.InputCrew.Faction.Color).Render(game.InputCrew.Name))
	output.WriteString("\n")
	output.WriteString(game.InputObjective.RenderCommitting(game.RollingPool, style))
	output.WriteString("\n")

	containerStyle := style.
		Width(28).
		Align(lipgloss.Center)

	validDice := game.RollingPool.GetValidDice(game.InputObjective.Type.ID)

	topDice := make([]string, 0)
	bottomDice := make([]string, 0)
	for i, die := range validDice {
		if i < 4 {
			topDice = append(topDice, die.Render(style, !die.Selected))
			continue
		}

		bottomDice = append(bottomDice, die.Render(style, !die.Selected))
	}

	var topSelected strings.Builder
	for i := range min(len(validDice), 4) {
		topSelected.WriteString("  ")
		if game.Index == i {
			topSelected.WriteString("⬇")
		} else {
			topSelected.WriteString(" ")
		}
		topSelected.WriteString("  ")
	}

	var bottomSelected strings.Builder
	for i := range max(0, len(validDice)-4) {
		bottomSelected.WriteString("  ")
		if game.Index == i+4 {
			bottomSelected.WriteString("⬆")
		} else {
			bottomSelected.WriteString(" ")
		}
		bottomSelected.WriteString("  ")
	}

	output.WriteString(containerStyle.Render(
		style.AlignHorizontal(lipgloss.Center).Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				topSelected.String(),
				lipgloss.JoinHorizontal(lipgloss.Top, topDice...),
				lipgloss.JoinHorizontal(lipgloss.Top, bottomDice...),
				bottomSelected.String(),
			),
		),
	))

	fmt.Fprintf(&output, "\n%s/%s to select a die", keys.GameChooseLeft.String(style), keys.GameChooseRight.String(style))
	fmt.Fprintf(&output, "\n%s to add/remove a die", keys.GameToggle.String(style))
	fmt.Fprintf(&output, "\n%s to confirm", keys.GameChooseConfirm.String(style))
	fmt.Fprintf(&output, "\n%s to end turn", keys.GameEndTurn.String(style))
	fmt.Fprintf(&output, "\n%s to go back", keys.GamePreviousInputStage.String(style))

	return inputComponentStyle(false).Render(output.String())
}

func inputStageCommitDiceHandler(s *tableScreen, msg tea.KeyMsg) (*Model, tea.Cmd) {
	game := s.model.Game

	switch {
	case keys.GameChooseLeft.TriggeredBy(msg.String()):
		game.Index--
		if game.Index < 0 {
			game.Index = len(game.RollingPool.GetValidDice(game.InputObjective.Type.ID)) - 1
		}
	case keys.GameChooseRight.TriggeredBy(msg.String()):
		game.Index++
		if game.Index >= len(game.RollingPool.GetValidDice(game.InputObjective.Type.ID)) {
			game.Index = 0
		}
	case keys.GameToggle.TriggeredBy(msg.String()):
		if game.RollingPool.NumberOfSelected() >= game.InputObjective.Amount && !game.RollingPool.GetValidDice(game.InputObjective.Type.ID)[game.Index].Selected {
			return s.model, nil
		}
		game.RollingPool.GetValidDice(game.InputObjective.Type.ID)[game.Index].Selected = !game.RollingPool.GetValidDice(game.InputObjective.Type.ID)[game.Index].Selected
		if game.RollingPool.GetValidDice(game.InputObjective.Type.ID)[game.Index].Selected {
			game.InputObjective.CommittingAmount += game.RollingPool.GetValidDice(game.InputObjective.Type.ID)[game.Index].Value
		} else {
			game.InputObjective.CommittingAmount -= game.RollingPool.GetValidDice(game.InputObjective.Type.ID)[game.Index].Value
		}
	case keys.GameChooseConfirm.TriggeredBy(msg.String()):
		if game.NumberOfCommittedDice() == 0 {
			return s.model, nil
		}
		game.CommitDice()
		game.InputState = game.nextStage()
		game.Index = 0
	case keys.GameEndTurn.TriggeredBy(msg.String()):
		game.NextTurn(false)
	case keys.GamePreviousInputStage.TriggeredBy(msg.String()):
		game.PreviousInputStage()
	}
	return s.model, nil
}
