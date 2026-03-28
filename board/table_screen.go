package board

import (
	"fmt"
	"strings"
	"time"

	"github.com/ascii-arcade/moonrollers/config"
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
	game := s.model.Game
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil

	case rollMsg:
		if s.rollTickCount < rollFrames {
			s.rollTickCount++
			game.Roll(s.isRolling)
			return s.model, tea.Tick(rollInterval, func(time.Time) tea.Msg {
				return rollMsg{}
			})
		}
		s.isRolling = false
		game.Roll(s.isRolling)

	case tea.KeyMsg:
		if game.GetCurrentPlayer() != s.model.Player {
			return s.model, nil
		}

		switch game.InputState {
		case InputStateRoll:
			return inputStageRollHandler(s, msg)
		case InputStateChooseCrew:
			return inputStageChooseCrewHandler(s, msg)
		case InputStateChooseObjective:
			return inputStageChooseObjectiveHandler(s, msg)
		case InputStateCommitDice:
			return inputStageCommitDiceHandler(s, msg)
		case InputStateChooseExtraDice:
			return inputStageExtraDiceHandler(s, msg)
		case InputStateChooseHazard:
			return inputStageChooseHazardHandler(s, msg)
		case InputStateOptionalHazard:
			return inputStageOptionalHazardHandler(s, msg)
		case InputStateReroll:
			return inputStageRerollHandler(s, msg)
		case Busted:
			if keys.GameEndTurn.TriggeredBy(msg.String()) {
				game.NextTurn(true)
			}
		}

		if config.Debug {
			switch {
			case keys.GameEndTurn.TriggeredBy(msg.String()):
				game.NextTurn(false)
			case msg.String() == "!":
				_ = game.HireCrewMember(0, s.model.Player)
			case msg.String() == "@":
				_ = game.HireCrewMember(1, s.model.Player)
			case msg.String() == "#":
				_ = game.HireCrewMember(2, s.model.Player)
			case msg.String() == "$":
				_ = game.HireCrewMember(3, s.model.Player)
			case msg.String() == "%":
				_ = game.HireCrewMember(4, s.model.Player)
			case msg.String() == "^":
				_ = game.HireCrewMember(5, s.model.Player)
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

	switch s.model.Game.InputState {
	case InputStateRoll:
		inputStageComponent = newInputStageRollComponent(s.model)
	case InputStateChooseCrew:
		inputStageComponent = newInputStageChooseCrewComponent(s.model)
	case InputStateChooseObjective:
		inputStageComponent = newInputStageChooseObjectiveComponent(s.model)
	case InputStateCommitDice:
		inputStageComponent = newInputStageCommitDiceComponent(s.model)
	case InputStateChooseExtraDice:
		inputStageComponent = newInputExtraDiceComponent(s.model)
	case InputStateChooseHazard:
		inputStageComponent = newInputStageChooseHazardComponent(s.model)
	case InputStateOptionalHazard:
		inputStageComponent = newInputStageOptionalHazardComponent(s.model)
	case InputStateReroll:
		inputStageComponent = newInputStageRerollComponent(s.model)
	case Busted:
		inputStageComponent = newBustedComponent(s.model)
	}

	rightSplit := lipgloss.JoinVertical(
		lipgloss.Left,
		supplyPoolComponent.render(),
		rollingPoolComponent.render(),
		inputStageComponent.render(),
	)

	var footer strings.Builder
	fmt.Fprintf(&footer, "%s", s.model.Player.Name)
	if s.model.Player.Points > 0 {
		fmt.Fprintf(&footer, " | %d points", s.model.Player.Points)
	}

	playerName := s.style.PaddingLeft(1).Render(s.model.Game.GetCurrentPlayer().Name + "'s turn")
	if s.model.Game.GetCurrentPlayer() == s.model.Player {
		playerName = s.style.PaddingLeft(1).Render("Your turn")
	}

	footer.WriteString(" - ")
	footer.WriteString(s.style.Render(playerName))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			scoreboardComponent.render(),
			forHireComponent.render(),
			rightSplit,
		),
		playerHandComponent.render(),
		s.style.Padding(0, 1).Render(footer.String()),
	)
}
