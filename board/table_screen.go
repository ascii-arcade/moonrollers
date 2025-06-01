package board

import (
	"fmt"

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
		cards = append(cards, newCard(s.model, card).render())
	}

	return s.style.Render(fmt.Sprintf("You are %s", s.model.Player.Name)) +
		"\n\n" + scoreboard.render() +
		"\n\n" + lipgloss.JoinHorizontal(lipgloss.Top, cards...) +
		"\n\n" + s.style.Render("Press 'ctrl+c' to quit")
}
