package menu

import (
	"math/rand/v2"
	"time"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/ascii-arcade/moonrollers/config"
	"github.com/ascii-arcade/moonrollers/dice"
	"github.com/ascii-arcade/moonrollers/keys"
	"github.com/ascii-arcade/moonrollers/language"
	"github.com/ascii-arcade/moonrollers/messages"
	"github.com/ascii-arcade/moonrollers/screen"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const logo = `            ++++*+
		 +++  *+  +*+
	 *++    *++++++   +++
  *+    ++++*+++++++++    ++
+    +++++++++++ ++++++++    +
+  +++++++ ++++++++ ++++++++ +
+       +++++ ++++++++++++++ +
+       ++++++++++++++++++++ +
+          +++++++ +++++++++ +
+             ++++*++++*++++ +
+             ++++*+++++++++ +
+             ++++++++++++++ +
+             +++++++++*++++ +
+             +++++++++      +
  ++          ++++++++    ++
     +*+      ++++    *+*
         +*+   +* +++*
            +++*++            `

type doneMsg struct{}

type Model struct {
	Width              int
	Height             int
	screen             screen.Screen
	style              lipgloss.Style
	languagePreference *language.LanguagePreference
	displayDice        []string

	errorCode     string
	gameCodeInput textinput.Model
}

func NewModel(width, height int, style lipgloss.Style, languagePreference *language.LanguagePreference) Model {
	ti := textinput.New()
	ti.Width = 9
	ti.CharLimit = 7

	m := Model{
		Width:              width,
		Height:             height,
		style:              style,
		languagePreference: languagePreference,
		displayDice:        make([]string, 0),

		gameCodeInput: ti,
	}
	for range 12 {
		i := rand.IntN(len(dice.All()))
		m.displayDice = append(m.displayDice, dice.All()[i].Render(style))
	}

	m.screen = m.newSplashScreen()
	return m
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.Tick(time.Second, func(t time.Time) tea.Msg {
			return doneMsg{}
		}),
		tea.WindowSize(),
		textinput.Blink,
	)
}

func (m *Model) lang() *language.Language {
	return m.languagePreference.Lang
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Height, m.Width = msg.Height, msg.Width
		return m, nil

	case messages.SwitchScreenMsg:
		m.screen = msg.Screen.WithModel(&m)
		return m, nil

	case tea.KeyMsg:
		if keys.ExitApplication.TriggeredBy(msg.String()) {
			return m, tea.Quit
		}
	}

	screenModel, cmd := m.screen.Update(msg)
	return screenModel.(*Model), cmd
}

func (m Model) View() string {
	if m.Width < config.MinimumWidth {
		return m.lang().Get("error.window_too_narrow")
	}
	if m.Height < config.MinimumHeight {
		return m.lang().Get("error.window_too_short")
	}

	style := m.style.Width(m.Width).Height(m.Height)
	paneStyle := m.style.Width(m.Width).PaddingTop(1)

	panes := lipgloss.JoinVertical(
		lipgloss.Center,
		paneStyle.Align(lipgloss.Center, lipgloss.Bottom).Foreground(colors.Logo).Height(m.Height/2).Render(m.style.Align(lipgloss.Left).Render(logo)),
		paneStyle.Align(lipgloss.Center, lipgloss.Top).Render(lipgloss.JoinHorizontal(lipgloss.Top, m.displayDice...)),
		paneStyle.Align(lipgloss.Center, lipgloss.Top).Render(m.screen.View()),
	)

	return style.Render(panes)
}

func (m *Model) setError(err string) {
	m.errorCode = err
}

func (m *Model) clearError() {
	m.errorCode = ""
}
