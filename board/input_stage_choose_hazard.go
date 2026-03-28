package board

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
	tea "github.com/charmbracelet/bubbletea"
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

	if c.model.Game.GetCurrentPlayer().Name != c.model.Player.Name {
		return inputComponentStyle(false).Render(fmt.Sprintf("%s is choosing a hazard...\n", c.model.Game.GetCurrentPlayer().Name))
	}

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

func inputStageChooseHazardHandler(s *tableScreen, msg tea.KeyMsg) (*Model, tea.Cmd) {
	game := s.model.Game

	switch {
	case keys.GameChooseHazard.TriggeredBy(msg.String()):
		i, _ := strconv.Atoi(msg.String())
		if i < 1 || i > len(game.InputHazards) {
			return s.model, nil
		}
		chosenHazard := game.InputHazards[i-1]
		game.GetCurrentPlayer().addHazard(chosenHazard.Points)
		game.InputState = InputStateRoll
	}
	return s.model, nil
}
