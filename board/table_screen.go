package board

import (
	"github.com/ascii-arcade/moonrollers/keys"
	"github.com/ascii-arcade/moonrollers/screen"
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

func (s *tableScreen) WithModel(model any) screen.Screen {
	s.model = model.(*Model)
	return s
}

func (s *tableScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil

	case tea.KeyMsg:
		if keys.GameIncrementPoint.TriggeredBy(msg.String()) {
			_ = s.model.Game.AddPoints(s.model.Player.Name, 1)
		}
	}

	return s.model, nil
}

func (s *tableScreen) View() string {
	scoreboard := newScoreboard(s.model)

	cards := make([]string, 0)
	for _, card := range s.model.Game.CrewForHire {
		cards = append(cards, newCard(s.model, card).renderForHire())
	}

	var cardRows []string
	if len(cards) > 0 {
		row1 := lipgloss.JoinHorizontal(lipgloss.Left, cards[:min(3, len(cards))]...)
		cardRows = append(cardRows, row1)
	}
	if len(cards) > 3 {
		row2 := lipgloss.JoinHorizontal(lipgloss.Left, cards[3:min(6, len(cards))]...)
		cardRows = append(cardRows, row2)
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		scoreboard.render(),
		lipgloss.JoinVertical(lipgloss.Left, cardRows...),
	)
}
