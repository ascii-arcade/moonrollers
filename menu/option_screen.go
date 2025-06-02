package menu

import (
	"fmt"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/ascii-arcade/moonrollers/keys"
	"github.com/ascii-arcade/moonrollers/language"
	"github.com/ascii-arcade/moonrollers/messages"
	"github.com/ascii-arcade/moonrollers/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type optionScreen struct {
	model *Model
	style lipgloss.Style
}

func (m *Model) newOptionScreen() *optionScreen {
	return &optionScreen{
		model: m,
		style: m.style,
	}
}

func (s *optionScreen) WithModel(model any) screen.Screen {
	s.model = model.(*Model)
	return s
}

func (s *optionScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if keys.MenuEnglish.TriggeredBy(msg.String()) {
			s.model.languagePreference.SetLanguage("EN")
		}
		if keys.MenuSpanish.TriggeredBy(msg.String()) {
			s.model.languagePreference.SetLanguage("ES")
		}
		if keys.MenuStartNewGame.TriggeredBy(msg.String()) {
			return s.model, func() tea.Msg { return messages.NewGame{} }
		}
		if keys.MenuJoinGame.TriggeredBy(msg.String()) {
			return s.model, func() tea.Msg {
				return messages.SwitchScreenMsg{
					Screen: s.model.newJoinScreen(),
				}
			}
		}
	}

	return s.model, nil
}

func (s *optionScreen) View() string {
	style := s.style.Width(s.model.Width).Height(s.model.Height)
	paneStyle := s.style.Width(s.model.Width).Height(s.model.Height / 2)

	content := "Welcome to the Game!\n\n"
	content += "Press " + keys.MenuStartNewGame.String(s.style) + " to create a new game.\n"
	content += "Press " + keys.MenuJoinGame.String(s.style) + " to join an existing game.\n"

	if s.model.lang() == language.Languages["EN"] {
		content += fmt.Sprintf(language.Languages["ES"].Get("menu.choose_language"), keys.MenuSpanish.String(s.style))
	} else if s.model.lang() == language.Languages["ES"] {
		content += fmt.Sprintf(language.Languages["EN"].Get("menu.choose_language"), keys.MenuEnglish.String(s.style))
	}

	panes := lipgloss.JoinVertical(
		lipgloss.Center,
		paneStyle.MarginBottom(2).Align(lipgloss.Center, lipgloss.Bottom).Foreground(colors.Logo).Render(logo),
		paneStyle.Align(lipgloss.Center, lipgloss.Top).Render(content),
	)

	return style.Render(panes)
}
