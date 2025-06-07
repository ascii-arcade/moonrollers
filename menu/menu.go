package menu

import (
	"errors"
	"math/rand/v2"
	"time"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/ascii-arcade/moonrollers/config"
	"github.com/ascii-arcade/moonrollers/dice"
	"github.com/ascii-arcade/moonrollers/games"
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
	width       int
	height      int
	screen      screen.Screen
	style       lipgloss.Style
	displayDice []string

	errorCode     string
	gameCodeInput textinput.Model

	player *games.Player
}

func NewModel(width, height int, style lipgloss.Style, player *games.Player) Model {
	ti := textinput.New()
	ti.Width = 9
	ti.CharLimit = 7

	m := Model{
		width:         width,
		height:        height,
		style:         style,
		displayDice:   make([]string, 0),
		gameCodeInput: ti,
		player:        player,
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
	return m.player.LanguagePreference.Lang
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
	if m.width < config.MinimumWidth {
		return m.lang().Get("error", "window_too_narrow")
	}
	if m.height < config.MinimumHeight {
		return m.lang().Get("error", "window_too_short")
	}

	style := m.style.Width(m.width).Height(m.height)
	paneStyle := m.style.Width(m.width).PaddingTop(1)

	panes := lipgloss.JoinVertical(
		lipgloss.Center,
		paneStyle.Align(lipgloss.Center, lipgloss.Bottom).Foreground(colors.Logo).Height(m.height/2).Render(m.style.Align(lipgloss.Left).Render(logo)),
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

func (m *Model) joinGame(code string, isNew bool) error {
	game, err := games.GetOpenGame(code)
	if err != nil && !(errors.Is(err, games.ErrGameInProgress) && game.HasPlayer(m.player)) {
		return err
	}
	if err := game.AddPlayer(m.player, isNew); err != nil {
		return err
	}
	return nil
}
