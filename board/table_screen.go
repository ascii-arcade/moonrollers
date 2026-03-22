package board

import (
	"strconv"
	"time"

	"github.com/ascii-arcade/moonrollers/config"
	"github.com/ascii-arcade/moonrollers/dice"
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
			switch {
			case keys.GameRollDice.TriggeredBy(msg.String()):
				if s.model.Game.InputObjective != nil && s.model.Game.InputObjective.IsCompleted() {
					s.model.Game.InputObjective = nil
				}
				if !s.isRolling {
					s.rollTickCount = 0
					s.isRolling = true
					return s.model, tea.Tick(rollInterval, func(time.Time) tea.Msg {
						return rollMsg{}
					})
				}
			case keys.GameEndTurn.TriggeredBy(msg.String()):
				s.model.Game.NextTurn(false)
				return s.model, nil
			}
		case games.InputStateChooseCrew:
			switch {
			case keys.GameChooseCrew.TriggeredBy(msg.String()):
				i, err := strconv.Atoi(msg.String())
				if err != nil {
					return s.model, nil
				}

				s.model.Game.ChooseCrewMember(i - 1)
			case keys.GameChooseConfirm.TriggeredBy(msg.String()):
				s.model.Game.ConfirmCrewMember()
			}

		case games.InputStateChooseObjective:
			switch {
			case keys.GameChooseObjective.TriggeredBy(msg.String()):
				i, _ := strconv.Atoi(msg.String())
				s.model.Game.ChooseObjective(i - 1)
			case keys.GameChooseConfirm.TriggeredBy(msg.String()):
				s.model.Game.ConfirmObjective()
			case keys.GamePreviousInputStage.TriggeredBy(msg.String()):
				s.model.Game.PreviousInputStage()
			}

		case games.InputStateCommitDice:
			switch {
			case keys.GameCommitDie.TriggeredBy(msg.String()):
				i, _ := strconv.Atoi(msg.String())
				switch {
				case s.model.Game.InputObjective.Amount-s.model.Game.InputObjective.CompletedAmount < i:
					i = min(
						s.model.Game.RollingPool.NumberOf(s.model.Game.InputObjective.Type),
						s.model.Game.InputObjective.Amount-s.model.Game.InputObjective.CompletedAmount,
					)
				case s.model.Game.RollingPool.NumberOf(s.model.Game.InputObjective.Type) < i:
					i = s.model.Game.RollingPool.NumberOf(s.model.Game.InputObjective.Type)
				case s.model.Game.InputObjective.CompletedAmount+i > s.model.Game.InputObjective.Amount:
					i = s.model.Game.InputObjective.Amount - s.model.Game.InputObjective.CompletedAmount
				}
				s.model.Game.InputObjective.CommittingAmount = i
			case keys.GameUncommitDie.TriggeredBy(msg.String()):
				s.model.Game.InputObjective.CommittingAmount -= 1
				if s.model.Game.InputObjective.CommittingAmount < 0 {
					s.model.Game.InputObjective.CommittingAmount = 0
				}
			case keys.GameChooseConfirm.TriggeredBy(msg.String()):
				s.model.Game.CommitDice()
				s.model.Game.InputState = games.InputStateRoll
				switch {
				case s.model.Game.RollingPool.HasExtra() && len(s.model.Game.SupplyPool.Dice) > 0:
					s.model.Game.InputState = games.InputStateChooseExtraDice
				case s.model.Game.InputObjective.Hazard && s.model.Game.InputObjective.IsCompleted():
					s.model.Game.PullHazards()
					s.model.Game.InputState = games.InputStateChooseHazard
				}
			case keys.GameEndTurn.TriggeredBy(msg.String()):
				s.model.Game.NextTurn(false)
			case keys.GamePreviousInputStage.TriggeredBy(msg.String()):
				s.model.Game.PreviousInputStage()
			}

		case games.InputStateChooseExtraDice:
			switch {
			case keys.GameChooseExtraDice.TriggeredBy(msg.String()):
				if s.model.Game.RollingPool.NumberOf(dice.DieUnrolled) < s.model.Game.RollingPool.NumberOf(dice.DieExtra) && s.model.Game.SupplyPool.Has(dice.DieUnrolled) {
					s.model.Game.RollingPool.Add(1)
					s.model.Game.SupplyPool.RemoveExtra()
				}
			case keys.GameUncommitDie.TriggeredBy(msg.String()):
				if s.model.Game.RollingPool.Has(dice.DieUnrolled) {
					s.model.Game.RollingPool.RemoveExtra()
					s.model.Game.SupplyPool.Add(1)
				}
			case keys.GameChooseConfirm.TriggeredBy(msg.String()):
				s.model.Game.InputState = games.InputStateRoll
				if s.model.Game.InputObjective.Hazard && s.model.Game.InputObjective.IsCompleted() {
					s.model.Game.PullHazards()
					s.model.Game.InputState = games.InputStateChooseHazard
				}
			}

		case games.InputStateChooseHazard:
			switch {
			case keys.GameChooseHazard.TriggeredBy(msg.String()):
				i, _ := strconv.Atoi(msg.String())
				if i < 1 || i > len(s.model.Game.InputHazards) {
					return s.model, nil
				}
				chosenHazard := s.model.Game.InputHazards[i-1]
				s.model.Game.GetCurrentPlayer().Hazards = append(s.model.Game.GetCurrentPlayer().Hazards, chosenHazard)
				s.model.Game.InputState = games.InputStateRoll
			}

		case games.Busted:
			if keys.GameEndTurn.TriggeredBy(msg.String()) {
				s.model.Game.NextTurn(true)
			}
		}

		if config.Debug {
			switch {
			case keys.GameEndTurn.TriggeredBy(msg.String()):
				s.model.Game.NextTurn(false)
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

	switch s.model.Game.InputState {
	case games.InputStateRoll:
		if s.model.Game.GetCurrentPlayer() == s.model.Player && !s.isRolling {
			inputStageComponent = newInputStageRollComponent(s.model)
		}
	case games.InputStateChooseCrew:
		inputStageComponent = newInputStageChooseCrewComponent(s.model)
	case games.InputStateChooseObjective:
		inputStageComponent = newInputStageChooseObjectiveComponent(s.model)
	case games.InputStateCommitDice:
		inputStageComponent = newInputStageCommitDiceComponent(s.model)
	case games.InputStateChooseExtraDice:
		inputStageComponent = newInputExtraDiceComponent(s.model)
	case games.InputStateChooseHazard:
		if s.model.Game.GetCurrentPlayer() == s.model.Player {
			inputStageComponent = newInputStageChooseHazardComponent(s.model)
		}
	case games.Busted:
		inputStageComponent = newBustedComponent(s.model)
	}

	playerName := s.style.PaddingLeft(1).Render(s.model.Game.GetCurrentPlayer().Name + "'s turn")
	if s.model.Game.GetCurrentPlayer() == s.model.Player {
		playerName = s.style.PaddingLeft(1).Render("Your turn")
	}

	rightSplit := lipgloss.JoinVertical(
		lipgloss.Left,
		supplyPoolComponent.render(),
		rollingPoolComponent.render(),
		inputStageComponent.render(),
		playerName,
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
