package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
	"github.com/charmbracelet/lipgloss"
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
	header := "Choose Crew"
	if c.model.Game.InputCrew != nil {
		header = c.model.style.Bold(true).Foreground(c.model.Game.InputCrew.Faction.Color).Render(c.model.Game.InputCrew.Name)
	}

	var crewList []string
	for index, crew := range c.model.Game.CrewForHire {
		text := fmt.Sprintf("[%d] %s", index+1, crew.Name)
		crewList = append(crewList, c.model.style.Foreground(crew.Faction.Color).Render(text))
	}

	confirm := ""
	if c.model.Game.InputCrew != nil {
		confirm = fmt.Sprintf("%s to confirm", keys.GameChooseConfirm.String(c.model.style))
	}

	return inputComponentStyle(false).Render(lipgloss.JoinVertical(
		lipgloss.Top,
		header+"\n\n",
		strings.Join(crewList, "\n"),
		"\n\n"+confirm,
	))
}
