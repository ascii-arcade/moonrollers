package menu

import (
	"time"

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
	Width  int
	Height int
	screen screen.Screen
	style  lipgloss.Style

	error         string
	gameCodeInput textinput.Model
}

func NewModel(width, height int, style lipgloss.Style) Model {
	ti := textinput.New()
	ti.Width = 9
	ti.CharLimit = 7

	m := Model{
		Width:  width,
		Height: height,
		style:  style,

		gameCodeInput: ti,
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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Height, m.Width = msg.Height, msg.Width
		return m, nil

	case messages.SwitchScreenMsg:
		m.screen = msg.Screen.WithModel(&m)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	screenModel, cmd := m.screen.Update(msg)
	return screenModel.(*Model), cmd
}

func (m Model) View() string {
	return m.screen.View()
}

func (m *Model) setError(err string) {
	m.error = err
}

func (m *Model) clearError() {
	m.error = ""
}
