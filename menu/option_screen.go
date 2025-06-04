package menu

import (
	"fmt"
	"strings"

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
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil

	case tea.KeyMsg:
		if keys.MenuEnglish.TriggeredBy(msg.String()) {
			s.model.player.LanguagePreference.SetLanguage("EN")
		}
		if keys.MenuSpanish.TriggeredBy(msg.String()) {
			s.model.player.LanguagePreference.SetLanguage("ES")
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
	var content strings.Builder
	content.WriteString(s.model.lang().Get("menu", "welcome") + "\n\n")
	content.WriteString(fmt.Sprintf(s.model.lang().Get("menu", "press_to_create"), keys.MenuStartNewGame.String(s.style)) + "\n")
	content.WriteString(fmt.Sprintf(s.model.lang().Get("menu", "press_to_join"), keys.MenuJoinGame.String(s.style)) + "\n")
	content.WriteString("\n\n")

	if s.model.lang() == language.Languages["EN"] {
		content.WriteString(fmt.Sprintf(language.Languages["ES"].Get("menu", "choose_language"), keys.MenuSpanish.String(s.style)))
	} else if s.model.lang() == language.Languages["ES"] {
		content.WriteString(fmt.Sprintf(language.Languages["EN"].Get("menu", "choose_language"), keys.MenuEnglish.String(s.style)))
	}

	return content.String()
}
