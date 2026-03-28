package board

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
	tea "github.com/charmbracelet/bubbletea"
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
	if c.model.Game.GetCurrentPlayer().Name != c.model.Player.Name {
		return inputComponentStyle(false).Render(fmt.Sprintf("%s is choosing a crew card...\n", c.model.Game.GetCurrentPlayer().Name))
	}

	var output strings.Builder
	if c.model.Game.InputCrew == nil {
		output.WriteString("Choose Crew")
	} else {
		output.WriteString(c.model.style.Bold(true).Foreground(c.model.Game.InputCrew.Faction.Color).Render(c.model.Game.InputCrew.Name))
	}
	output.WriteString("\n\n")

	crewList := make([]string, 0)
	for i, crew := range c.model.Game.CrewForHire {
		if crew.CanCommit(c.model.Game.RollingPool, c.model.Game.GetCurrentPlayer().Name) {
			text := fmt.Sprintf("[%d] %s", i+1, crew.Name)
			crewList = append(crewList, c.model.style.Foreground(crew.Faction.Color).Render(text))
		}
	}
	output.WriteString(lipgloss.JoinVertical(lipgloss.Top, crewList...))

	if c.model.Game.InputCrew != nil && c.model.Game.GetCurrentPlayer().Name == c.model.Player.Name {
		fmt.Fprintf(&output, "\n\n%s to confirm", keys.GameChooseConfirm.String(c.model.style))
	}

	return inputComponentStyle(false).Render(output.String())
}

func inputStageChooseCrewHandler(s *tableScreen, msg tea.KeyMsg) (*Model, tea.Cmd) {
	game := s.model.Game

	switch {
	case keys.GameChooseCrew.TriggeredBy(msg.String()):
		i, err := strconv.Atoi(msg.String())
		if err != nil {
			return s.model, nil
		}
		game.ChooseCrewMember(i - 1)
	case keys.GameChooseConfirm.TriggeredBy(msg.String()):
		game.ConfirmCrewMember()
	}
	return s.model, nil
}
