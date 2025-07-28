package board

import (
	"time"

	"github.com/ascii-arcade/moonrollers/config"
	"github.com/ascii-arcade/moonrollers/games"
	"github.com/ascii-arcade/moonrollers/keys"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tableScreen struct {
	model *Model
	style lipgloss.Style
}

type rollMsg struct{}

func (m *Model) newTableScreen() *tableScreen {
	return &tableScreen{
		model: m,
		style: m.style,
	}
}

func (s *tableScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil

	case rollMsg:
		if s.model.Game.RollTick < rollFrames {
			s.model.Game.RollTick++
			s.model.Game.Roll(s.model.Game.IsRolling)
			return s.model, tea.Tick(rollInterval, func(time.Time) tea.Msg {
				return rollMsg{}
			})
		}
		s.model.Game.IsRolling = false
		s.model.Game.Roll(s.model.Game.IsRolling)

	case tea.KeyMsg:
		if s.model.Game.GetCurrentPlayer() != s.model.Player {
			return s.model, nil
		}

		switch {
		case keys.GameEndTurn.TriggeredBy(msg.String()):
			s.model.Game.NextTurn()
			return s.model, nil
		}

		if config.Debug {
			switch {
			case msg.String() == "a":
				_ = s.model.Game.HireCrewMember(0, s.model.Player)
			case msg.String() == "r":
				_ = s.model.Game.HireCrewMember(1, s.model.Player)
			case msg.String() == "s":
				_ = s.model.Game.HireCrewMember(2, s.model.Player)
			case msg.String() == "t":
				_ = s.model.Game.HireCrewMember(3, s.model.Player)
			case msg.String() == "d":
				_ = s.model.Game.HireCrewMember(4, s.model.Player)
			case msg.String() == "h":
				_ = s.model.Game.HireCrewMember(5, s.model.Player)
			}
			return s.model, nil
		}
	}

	return s.model, nil
}

func (s *tableScreen) View() string {
	rollingPoolComponent := newDiceComponent(s.model, s.model.Game.RollingPool)
	supplyPoolComponent := newDiceComponent(s.model, s.model.Game.SupplyPool)
	forHireComponent := newForHireComponent(s.model)
	playerHandComponent := newPlayerHandComponent(s.model)
	scoreboardComponent := newScoreboardComponent(s.model)

	var inputStageComponent inputStageComponent
	inputStageComponent = newInputStageEmptyComponent(s.model)

	if s.model.Game.GetCurrentPlayer() == s.model.Player {
		if s.model.Game.InputState == games.InputStateRoll && !s.model.Game.IsRolling {
			inputStageComponent = newInputStageRollComponent(s.model)
		}
	}

	rightSplit := lipgloss.JoinVertical(
		lipgloss.Left,
		supplyPoolComponent.render(),
		rollingPoolComponent.render(),
		inputStageComponent.render(),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			scoreboardComponent.render(),
			forHireComponent.render(),
			rightSplit,
		),
		playerHandComponent.render(),
	)
}
