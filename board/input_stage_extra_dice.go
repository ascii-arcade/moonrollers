package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
	tea "github.com/charmbracelet/bubbletea"
)

type inputExtraDiceComponent struct {
	model *Model
}

func newInputExtraDiceComponent(model *Model) inputExtraDiceComponent {
	return inputExtraDiceComponent{
		model: model,
	}
}

func (c inputExtraDiceComponent) render() string {
	if c.model.Game.GetCurrentPlayer().Name != c.model.Player.Name {
		return inputComponentStyle(false).Render(fmt.Sprintf("%s is choosing extra ..\n", c.model.Game.GetCurrentPlayer().Name))
	}
	var output strings.Builder
	output.WriteString("How many extra die would you like to add?\n\n")
	fmt.Fprintf(&output, "%s to add from Supply\n", keys.GameChooseExtraDice.String(c.model.style))
	fmt.Fprintf(&output, "%s to remove extra from Pool\n", keys.GameRemove.String(c.model.style))
	fmt.Fprintf(&output, "%s to continue", keys.GameChooseConfirm.String(c.model.style))
	return inputComponentStyle(true).Render(output.String())
}

func inputStageExtraDiceHandler(s *tableScreen, msg tea.KeyMsg) (*Model, tea.Cmd) {
	game := s.model.Game

	switch {
	case keys.GameChooseExtraDice.TriggeredBy(msg.String()):
		if game.RollingPool.NumberOfUnrolled() < game.RollingPool.NumberOfExtra() && game.SupplyPool.Length() > 0 {
			game.RollingPool.AddUnrolled()
			game.SupplyPool.RemoveUnrolled()
		}
	case keys.GameRemove.TriggeredBy(msg.String()):
		if game.RollingPool.NumberOfUnrolled() > 0 {
			game.RollingPool.RemoveUnrolled()
			game.SupplyPool.AddUnrolled()
		}
	case keys.GameChooseConfirm.TriggeredBy(msg.String()):
		game.InputState = InputStateRoll
		if game.InputObjective.Hazard && game.InputObjective.IsCompleted() {
			game.PullHazards()
			game.InputState = InputStateChooseHazard
		}
	}
	return s.model, nil
}
