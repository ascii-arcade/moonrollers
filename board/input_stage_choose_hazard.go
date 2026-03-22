package board

import (
	"fmt"
	"strings"
)

type inputStageChooseHazardComponent struct {
	model *Model
}

func newInputStageChooseHazardComponent(model *Model) inputStageChooseHazardComponent {
	return inputStageChooseHazardComponent{
		model: model,
	}
}

func (c inputStageChooseHazardComponent) render() string {
	var output strings.Builder
	output.WriteString("Choose a hazard to keep:\n")

	for i, hazard := range c.model.Game.InputHazards {
		fmt.Fprintf(&output, "[%d] %d", i+1, hazard.Points)
		for range hazard.Hazards {
			output.WriteString("⚠")
		}
		output.WriteString("\n")
	}

	return inputComponentStyle(false).Render(output.String())
}
