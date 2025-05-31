package board

import (
	"fmt"

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

func (s *tableScreen) setModel(model *Model) {
	s.model = model
}

func (s *tableScreen) update(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "a":
		s.model.Game.AddPoints(s.model.Player.Name, 1)
	}

	return s.model, nil
}

func (s *tableScreen) view() string {
	scoreboard := newScoreboard(s.model)

	card := newCard(s.model, s.model.Game.Deck[0])

	return s.style.Render(fmt.Sprintf("You are %s", s.model.Player.Name)) +
		"\n\n" + scoreboard.render() +
		"\n\n" + card.render() +
		"\n\n" + s.style.Render("Press 'ctrl+c' to quit")
}
