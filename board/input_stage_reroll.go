package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type inputStageRerollComponent struct {
	model *Model
}

func newInputStageRerollComponent(model *Model) inputStageRerollComponent {
	return inputStageRerollComponent{
		model: model,
	}
}

func (c inputStageRerollComponent) render() string {
	game := c.model.Game
	style := c.model.style

	var output strings.Builder

	if c.model.Game.GetCurrentPlayer().Name != c.model.Player.Name {
		return inputComponentStyle(false).Render(fmt.Sprintf("%s is deciding to reroll dice...\n", c.model.Game.GetCurrentPlayer().Name))
	}

	output.WriteString("Which dice would you like to reroll?")
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
	fmt.Fprintf(&output, "\n%s to reroll", keys.GameRollDice.String(style))

	return inputComponentStyle(false).Render(output.String())
}

func inputStageRerollHandler(s *tableScreen, msg tea.KeyMsg) (*Model, tea.Cmd) {
	game := s.model.Game

	switch {
	case keys.GameChooseLeft.TriggeredBy(msg.String()):
		if game.Index > 0 {
			game.Index--
		}
	case keys.GameChooseRight.TriggeredBy(msg.String()):
		if game.Index < game.RollingPool.Length()-1 {
			game.Index++
		}
	case keys.GameToggle.TriggeredBy(msg.String()):
		if game.RollingPool.Length() > 0 {
			die := game.RollingPool.Dice[game.Index]
			die.Selected = !die.Selected
		}
	case keys.GameRollDice.TriggeredBy(msg.String()):
		var tempPool DicePool
		for _, die := range game.RollingPool.Dice {
			if die.Selected {
				tempPool.Add(die)
				game.RollingPool.Remove(die)
			}
		}
		if tempPool.Length() > 0 {
			return s.model, nil
		}

		tempPool.Roll()
		for _, die := range tempPool.Dice {
			game.RollingPool.Add(die)
		}

		game.InputState = game.nextStage()
	}

	return s.model, nil
}
