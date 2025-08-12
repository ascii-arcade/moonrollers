package board

import (
	"strconv"
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

	rollTickCount int
	isRolling     bool
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
		if s.rollTickCount < rollFrames {
			s.rollTickCount++
			s.model.Game.Roll(s.isRolling)
			return s.model, tea.Tick(rollInterval, func(time.Time) tea.Msg {
				return rollMsg{}
			})
		}
		s.isRolling = false
		s.model.Game.Roll(s.isRolling)

	case tea.KeyMsg:
		if s.model.Game.GetCurrentPlayer() != s.model.Player {
			return s.model, nil
		}

		switch s.model.Game.InputState {
		case games.InputStateRoll:
			if keys.GameRollDice.TriggeredBy(msg.String()) {
				if !s.isRolling {
					s.rollTickCount = 0
					s.isRolling = true
					return s.model, tea.Tick(rollInterval, func(time.Time) tea.Msg {
						return rollMsg{}
					})
				}
			}
		case games.InputStateChooseCrew:
			switch {
			case keys.GameChooseCrew.TriggeredBy(msg.String()):
				i, err := strconv.Atoi(msg.String())
				if err != nil {
					return s.model, nil
				}

				s.model.Game.ChooseCrewMember(i - 1)
				return s.model, nil
			case keys.GameChooseConfirm.TriggeredBy(msg.String()):
				s.model.Game.ConfirmCrewMember()
			}

		case games.InputStateChooseObjective:
			switch {
			case keys.GameChooseObjective.TriggeredBy(msg.String()):
				i, err := strconv.Atoi(msg.String())
				if err != nil {
					return s.model, nil
				}

				s.model.Game.ChooseObjective(i - 1)
				return s.model, nil
			case keys.GameChooseConfirm.TriggeredBy(msg.String()):
				s.model.Game.ConfirmObjective()
			case keys.GamePreviousInputStage.TriggeredBy(msg.String()):
				s.model.Game.PreviousInputStage()
			}
		}

		if config.Debug {
			switch {
			case keys.GameEndTurn.TriggeredBy(msg.String()):
				s.model.Game.NextTurn()
			case msg.String() == "!":
				_ = s.model.Game.HireCrewMember(0, s.model.Player)
			case msg.String() == "@":
				_ = s.model.Game.HireCrewMember(1, s.model.Player)
			case msg.String() == "#":
				_ = s.model.Game.HireCrewMember(2, s.model.Player)
			case msg.String() == "$":
				_ = s.model.Game.HireCrewMember(3, s.model.Player)
			case msg.String() == "%":
				_ = s.model.Game.HireCrewMember(4, s.model.Player)
			case msg.String() == "^":
				_ = s.model.Game.HireCrewMember(5, s.model.Player)
			}
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
	inputStageComponent = newInputStageEmptyComponent()

	if s.model.Game.GetCurrentPlayer() == s.model.Player {
		switch s.model.Game.InputState {
		case games.InputStateRoll:
			if !s.isRolling {
				inputStageComponent = newInputStageRollComponent(s.model)
			}
		case games.InputStateChooseCrew:
			inputStageComponent = newInputStageChooseCrewComponent(s.model)
		case games.InputStateChooseObjective:
			inputStageComponent = newInputStageChooseObjectiveComponent(s.model)
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
