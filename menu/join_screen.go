package menu

import (
	"strings"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/ascii-arcade/moonrollers/games"
	"github.com/ascii-arcade/moonrollers/keys"
	"github.com/ascii-arcade/moonrollers/messages"
	"github.com/ascii-arcade/moonrollers/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type joinScreen struct {
	model *Model
	style lipgloss.Style
}

func (m *Model) newJoinScreen() *joinScreen {
	m.gameCodeInput.Focus()
	m.gameCodeInput.SetValue("")
	return &joinScreen{
		model: m,
		style: m.style,
	}
}

func (s *joinScreen) WithModel(model any) screen.Screen {
	s.model = model.(*Model)
	return s
}

func (s *joinScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if keys.PreviousScreen.TriggeredBy(msg.String()) {
			return s.model, func() tea.Msg {
				return messages.SwitchScreenMsg{
					Screen: s.model.newOptionScreen(),
				}
			}
		}
		if keys.Submit.TriggeredBy(msg.String()) {
			if len(s.model.gameCodeInput.Value()) == 7 {
				code := strings.ToUpper(s.model.gameCodeInput.Value())
				_, err := games.GetOpenGame(code)
				if err != nil {
					s.model.setError(err.Error())
					return s.model, nil
				}
				return s.model, func() tea.Msg { return messages.JoinGame{GameCode: code} }
			}
		}

		s.model.clearError()
		val := s.model.gameCodeInput.Value()
		if len(val) == 3 && msg.Type == tea.KeyRunes && msg.Runes[0] != '-' {
			val = val + "-"
			s.model.gameCodeInput.SetValue(val)
			s.model.gameCodeInput.CursorEnd()
		}
	}

	s.model.gameCodeInput, cmd = s.model.gameCodeInput.Update(msg)
	return s.model, cmd
}

func (s *joinScreen) View() string {
	style := s.style.Width(s.model.Width).Height(s.model.Height)
	paneStyle := s.style.Width(s.model.Width).Height(s.model.Height / 2)

	content := s.model.lang().Get("menu.enter_code") + "\n\n" + s.model.gameCodeInput.View()

	panes := lipgloss.JoinVertical(
		lipgloss.Center,
		paneStyle.MarginBottom(2).Align(lipgloss.Center, lipgloss.Bottom).Foreground(colors.Logo).Render(logo),
		paneStyle.Align(lipgloss.Center, lipgloss.Top).Render(content+"\n\n"+s.style.Foreground(colors.Error).Render(s.model.error)),
	)

	return style.Render(panes)
}
