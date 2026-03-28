package board

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
	tea "github.com/charmbracelet/bubbletea"
)

type inputStageChooseObjectiveComponent struct {
	model *Model
}

func newInputStageChooseObjectiveComponent(model *Model) inputStageChooseObjectiveComponent {
	return inputStageChooseObjectiveComponent{
		model: model,
	}
}

func (c inputStageChooseObjectiveComponent) render() string {
	if c.model.Game.GetCurrentPlayer().Name != c.model.Player.Name {
		return inputComponentStyle(false).Render(fmt.Sprintf("%s is choosing an objective...\n", c.model.Game.GetCurrentPlayer().Name))
	}

	var output strings.Builder
	output.WriteString(c.model.style.Bold(true).Foreground(c.model.Game.InputCrew.Faction.Color).Render(c.model.Game.InputCrew.Name))
	output.WriteString("\n")
	if c.model.Game.InputObjective == nil {
		output.WriteString("Choose Objective")
	}
	output.WriteString("\n\n")

	for index, objective := range c.model.Game.InputCrew.Objectives {
		switch {
		case objective.IsCompleted(),
			objective.StartedBy != "" && objective.StartedBy != c.model.Game.GetCurrentPlayer().Name,
			c.model.Game.RollingPool.ValueOf(objective.Type.ID) == 0:
			continue
		}
		fmt.Fprintf(&output, "[%d] %s\n", index+1, objective.Render(c.model.style, c.model.Game.InputObjective != nil && c.model.Game.InputObjective == objective))
	}

	if c.model.Game.InputObjective != nil {
		fmt.Fprintf(&output, "\n%s to confirm", keys.GameChooseConfirm.String(c.model.style))
	}

	fmt.Fprintf(&output, "\n%s to go back", keys.GamePreviousInputStage.String(c.model.style))

	return inputComponentStyle(false).Render(output.String())
}

func inputStageChooseObjectiveHandler(s *tableScreen, msg tea.KeyMsg) (*Model, tea.Cmd) {
	game := s.model.Game

	switch {
	case keys.GameChooseObjective.TriggeredBy(msg.String()):
		i, _ := strconv.Atoi(msg.String())
		game.ChooseObjective(i - 1)
	case keys.GameChooseConfirm.TriggeredBy(msg.String()):
		game.ConfirmObjective()
	case keys.GamePreviousInputStage.TriggeredBy(msg.String()):
		game.PreviousInputStage()
	}
	return s.model, nil
}
