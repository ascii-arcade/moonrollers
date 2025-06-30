package board

import (
	"github.com/ascii-arcade/moonrollers/keys"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type settingsScreen struct {
	model *Model
	style lipgloss.Style
	form  *huh.Form

	useStarterCards bool
}

func (m *Model) newSettingsScreen() *settingsScreen {
	s := &settingsScreen{
		model:           m,
		style:           m.style,
		useStarterCards: m.Game.Settings.UseStarterCards,
	}
	s.form = huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Use starter cards?").
				Value(&s.useStarterCards),
		),
	)
	return s
}

func (s *settingsScreen) Init() tea.Cmd {
	return s.form.Init()
}

func (s *settingsScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case keys.PreviousScreen.TriggeredBy(msg.String()):
			s.model.screen = s.model.newLobbyScreen()
			return s.model, nil
		}
	}

	form, cmd := s.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		s.form = f
	}

	if s.form.State == huh.StateCompleted {
		s.model.Game.Settings.UseStarterCards = s.useStarterCards

		s.model.screen = s.model.newLobbyScreen()
	}

	return s.model, cmd
}

func (s *settingsScreen) View() string {
	return "Settings:\n\n" + s.form.View()
}
