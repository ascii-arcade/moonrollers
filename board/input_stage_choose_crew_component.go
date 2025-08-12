package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
)

type inputStageChooseCrewComponent struct {
	model *Model
}

func newInputStageChooseCrewComponent(model *Model) inputStageChooseCrewComponent {
	return inputStageChooseCrewComponent{
		model: model,
	}
}

func (c inputStageChooseCrewComponent) render() string {
	var output strings.Builder
	if c.model.Game.InputCrew == nil {
		output.WriteString("Choose Crew")
	} else {
		output.WriteString(c.model.style.Bold(true).Foreground(c.model.Game.InputCrew.Faction.Color).Render(c.model.Game.InputCrew.Name))
	}
	output.WriteString("\n\n")

	for index, crew := range c.model.Game.CrewForHire {
		text := fmt.Sprintf("[%d] %s\n", index+1, crew.Name)
		output.WriteString(c.model.style.Foreground(crew.Faction.Color).Render(text))
	}

	if c.model.Game.InputCrew != nil {
		output.WriteString(fmt.Sprintf("\n\n%s to confirm", keys.GameChooseConfirm.String(c.model.style)))
	}

	return inputComponentStyle(false).Render(output.String())
}
