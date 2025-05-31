package board

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type tableScreen struct {
	model *Model
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
	scoreboard := scoreboard{
		model:   s.model,
		players: s.model.Game.OrderedPlayers(),
		short:   false,
	}

	return s.model.renderer.NewStyle().Render(fmt.Sprintf("You are %s", s.model.Player.Name)) +
		"\n\n" + scoreboard.render() +
		"\n\n" + s.model.renderer.NewStyle().Render("Press 'ctrl+c' to quit")
}
