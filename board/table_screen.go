package board

import (
	"github.com/ascii-arcade/moonrollers/keys"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tableScreen struct {
	model *Model
	style lipgloss.Style
}

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

	case tea.KeyMsg:
		if keys.GameIncrementPoint.TriggeredBy(msg.String()) {
			_ = s.model.Game.AddPoints(s.model.Player, 1)
		}
	}

	return s.model, nil
}

func (s *tableScreen) View() string {
	forHire := newForHire(s.model)
	playerHand := newPlayerHand(s.model)
	scoreboard := newScoreboard(s.model)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			scoreboard.render(),
			forHire.render(),
		),
		playerHand.render(),
	)
}
