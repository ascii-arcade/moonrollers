package board

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
	tea "github.com/charmbracelet/bubbletea"
)

type inputStageOptionalHazardComponent struct {
	model *Model
}

func newInputStageOptionalHazardComponent(model *Model) inputStageOptionalHazardComponent {
	return inputStageOptionalHazardComponent{
		model: model,
	}
}

func (c inputStageOptionalHazardComponent) render() string {
	var output strings.Builder

	if c.model.Game.GetCurrentPlayer().Name != c.model.Player.Name {
		return inputComponentStyle(false).Render(fmt.Sprintf("%s is deciding to take a hazard token...\n", c.model.Game.GetCurrentPlayer().Name))
	}

	output.WriteString("Would you like to take a hazard token?")

	return inputComponentStyle(false).Render(output.String())
}

func inputStageOptionalHazardHandler(s *tableScreen, msg tea.KeyMsg) (*Model, tea.Cmd) {
	game := s.model.Game

	switch {
	case keys.GameChooseConfirm.TriggeredBy(msg.String()):
		game.GetCurrentPlayer().addHazard(hazards[rand.IntN(2)].Points)
		game.InputState = game.nextStage()
	case keys.GamePreviousInputStage.TriggeredBy(msg.String()):
		game.InputState = game.nextStage()
	}
	return s.model, nil
}
