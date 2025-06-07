package menu

import (
	"errors"
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
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil

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

				game, err := games.GetOpenGame(code)
				if err != nil {
					if !errors.Is(err, games.ErrGameInProgress) && !game.HasPlayer(s.model.player) {
						s.model.setError(err.Error())
						return s.model, nil
					}
				}

				if err := s.model.joinGame(code, false); err != nil {
					s.model.setError(err.Error())
					return s.model, nil
				}

				return s.model, func() tea.Msg { return messages.SwitchToBoardMsg{Game: game} }
			}
		}

		s.model.clearError()
	}

	s.model.gameCodeInput, cmd = s.model.gameCodeInput.Update(msg)
	s.model.gameCodeInput.SetValue(strings.ToUpper(s.model.gameCodeInput.Value()))

	if len(s.model.gameCodeInput.Value()) == 3 && !strings.Contains(s.model.gameCodeInput.Value(), "-") {
		s.model.gameCodeInput.SetValue(s.model.gameCodeInput.Value() + "-")
		s.model.gameCodeInput.CursorEnd()
	}
	return s.model, cmd
}

func (s *joinScreen) View() string {
	errorMessage := ""
	if s.model.errorCode != "" {
		errorMessage = s.model.lang().Get("error", s.model.errorCode)
	}

	var content strings.Builder
	content.WriteString(s.model.lang().Get("menu", "enter_code") + "\n\n")
	content.WriteString(s.model.gameCodeInput.View() + "\n\n")
	content.WriteString(s.style.Foreground(colors.Error).Render(errorMessage))

	return content.String()
}
