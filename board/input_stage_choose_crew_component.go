package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
	"github.com/ascii-arcade/moonrollers/rules"
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
	var output strings.Builder
	if c.model.Game.InputCrew == nil {
		output.WriteString("Choose Crew")
	} else {
		output.WriteString(c.model.style.Bold(true).Foreground(c.model.Game.InputCrew.Faction.Color).Render(c.model.Game.InputCrew.Name))
	}
	output.WriteString("\n\n")

	commitableToCrew := rules.CommitableToCrew(
		c.model.Player.CrewIDs(),
		c.model.Game.CrewForHire,
		c.model.Game.RollingPool,
	)
	crewList := make([]string, 0)
	for index, crew := range commitableToCrew {
		text := fmt.Sprintf("[%d] %s", index+1, crew.Name)
		crewList = append(crewList, c.model.style.Foreground(crew.Faction.Color).Render(text))
	}
	output.WriteString(lipgloss.JoinVertical(lipgloss.Top, crewList...))

	if c.model.Game.InputCrew != nil {
		output.WriteString(fmt.Sprintf("\n\n%s to confirm", keys.GameChooseConfirm.String(c.model.style)))
	}

	return inputComponentStyle(false).Render(output.String())
}
