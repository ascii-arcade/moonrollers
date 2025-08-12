package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
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
	var output strings.Builder
	output.WriteString(c.model.style.Bold(true).Foreground(c.model.Game.InputCrew.Faction.Color).Render(c.model.Game.InputCrew.Name))
	output.WriteString("\n")
	if c.model.Game.InputObjective == nil {
		output.WriteString("Choose Objective")
	} else {
		output.WriteString(c.model.Game.InputObjective.Render(c.model.style))
	}
	output.WriteString("\n\n")

	for index, objective := range c.model.Game.InputCrew.Objectives {
		output.WriteString(fmt.Sprintf("[%d] %s\n", index+1, objective.Render(c.model.style)))
	}

	if c.model.Game.InputObjective != nil {
		output.WriteString(fmt.Sprintf("\n\n%s to confirm", keys.GameChooseConfirm.String(c.model.style)))
	}

	output.WriteString(fmt.Sprintf("\n%s to go back", keys.GamePreviousInputStage.String(c.model.style)))

	return inputComponentStyle(false).Render(output.String())
}
