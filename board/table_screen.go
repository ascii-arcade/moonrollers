package board

import (
	"fmt"
	"strconv"
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
			switch {
			case keys.GameRollDice.TriggeredBy(msg.String()):
				if game.InputObjective != nil && game.InputObjective.IsCompleted() {
					game.InputObjective = nil
				}
				if !s.isRolling {
					s.rollTickCount = 0
					s.isRolling = true
					return s.model, tea.Tick(rollInterval, func(time.Time) tea.Msg {
						return rollMsg{}
					})
				}
			case keys.GameEndTurn.TriggeredBy(msg.String()):
				game.NextTurn(false)
				return s.model, nil
			}
		case InputStateChooseCrew:
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

		case InputStateChooseObjective:
			switch {
			case keys.GameChooseObjective.TriggeredBy(msg.String()):
				i, _ := strconv.Atoi(msg.String())
				game.ChooseObjective(i - 1)
			case keys.GameChooseConfirm.TriggeredBy(msg.String()):
				game.ConfirmObjective()
			case keys.GamePreviousInputStage.TriggeredBy(msg.String()):
				game.PreviousInputStage()
			}

		case InputStateCommitDice:
			switch {
			case keys.GameChooseLeft.TriggeredBy(msg.String()):
				game.Index--
				if game.Index < 0 {
					game.Index = len(game.RollingPool.GetValidDice(game.InputObjective.Type.ID)) - 1
				}
			case keys.GameChooseRight.TriggeredBy(msg.String()):
				game.Index++
				if game.Index >= len(game.RollingPool.GetValidDice(game.InputObjective.Type.ID)) {
					game.Index = 0
				}
			case keys.GameToggle.TriggeredBy(msg.String()):
				if game.RollingPool.NumberOfSelected() >= game.InputObjective.Amount && !game.RollingPool.GetValidDice(game.InputObjective.Type.ID)[game.Index].Selected {
					return s.model, nil
				}
				game.RollingPool.GetValidDice(game.InputObjective.Type.ID)[game.Index].Selected = !game.RollingPool.GetValidDice(game.InputObjective.Type.ID)[game.Index].Selected
				if game.RollingPool.GetValidDice(game.InputObjective.Type.ID)[game.Index].Selected {
					game.InputObjective.CommittingAmount += game.RollingPool.GetValidDice(game.InputObjective.Type.ID)[game.Index].Value
				} else {
					game.InputObjective.CommittingAmount -= game.RollingPool.GetValidDice(game.InputObjective.Type.ID)[game.Index].Value
				}
			case keys.GameChooseConfirm.TriggeredBy(msg.String()):
				if game.NumberOfCommittedDice() == 0 {
					return s.model, nil
				}
				game.CommitDice()
				game.InputState = InputStateRoll
				switch {
				case game.InputCrew.IsComplete():
					game.CompleteCard(game.InputCrew, game.GetCurrentPlayer())
					game.NextTurn(false)
				case game.RollingPool.HasExtra() && len(game.SupplyPool.Dice) > 0:
					game.InputState = InputStateChooseExtraDice
				case game.InputObjective.Hazard && game.InputObjective.IsCompleted():
					game.PullHazards()
					game.InputState = InputStateChooseHazard
				}
				game.Index = 0
			case keys.GameEndTurn.TriggeredBy(msg.String()):
				game.NextTurn(false)
			case keys.GamePreviousInputStage.TriggeredBy(msg.String()):
				game.PreviousInputStage()
			}

		case InputStateChooseExtraDice:
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

		case InputStateChooseHazard:
			switch {
			case keys.GameChooseHazard.TriggeredBy(msg.String()):
				i, _ := strconv.Atoi(msg.String())
				if i < 1 || i > len(game.InputHazards) {
					return s.model, nil
				}
				chosenHazard := game.InputHazards[i-1]
				game.GetCurrentPlayer().Hazards = append(game.GetCurrentPlayer().Hazards, chosenHazard)
				game.InputState = InputStateRoll
			}

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
		if s.model.Game.GetCurrentPlayer() == s.model.Player && !s.isRolling {
			inputStageComponent = newInputStageRollComponent(s.model)
		}
	case InputStateChooseCrew:
		inputStageComponent = newInputStageChooseCrewComponent(s.model)
	case InputStateChooseObjective:
		inputStageComponent = newInputStageChooseObjectiveComponent(s.model)
	case InputStateCommitDice:
		inputStageComponent = newInputStageCommitDiceComponent(s.model)
	case InputStateChooseExtraDice:
		inputStageComponent = newInputExtraDiceComponent(s.model)
	case InputStateChooseHazard:
		if s.model.Game.GetCurrentPlayer() == s.model.Player {
			inputStageComponent = newInputStageChooseHazardComponent(s.model)
		}
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
