package board

import (
	"strconv"
	"strings"

	"github.com/ascii-arcade/moonrollers/factions"
	"github.com/ascii-arcade/moonrollers/screen"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type lobbyScreen struct {
	model *Model
	style lipgloss.Style
}

func (m *Model) newLobbyScreen() *lobbyScreen {
	return &lobbyScreen{
		model: m,
		style: m.style,
	}
}

func (s *lobbyScreen) WithModel(model any) screen.Screen {
	s.model = model.(*Model)
	return s
}

func (s *lobbyScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1", "2", "3", "4", "5":
			i, err := strconv.Atoi(msg.String())
			if err != nil {
				return s.model, nil
			} else {
				faction := factions.All()[i-1]
				if !s.model.Game.IsFactionUsed(faction) {
					s.model.Game.SetFaction(s.model.Player, &faction)
				}
			}
		case "s":
			allHaveColor := true
			for _, p := range s.model.Game.OrderedPlayers() {
				if !p.HasFaction() {
					allHaveColor = false
					break
				}
			}
			if s.model.Player.IsHost() && allHaveColor {
				s.model.Game.Begin()
			}
		}
	}

	return s.model, nil
}

func (s *lobbyScreen) View() string {
	style := s.style.Width(s.model.Width / 3)

	header := s.model.Game.Code
	playerList := s.style.Render(s.playerList())
	footer := s.style.Render(s.footer())

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		style.Align(lipgloss.Center).MarginBottom(2).Render(header),
		style.Render(playerList),
		style.Render(footer),
	)

	return s.style.Width(s.model.Width).Height(s.model.Height).Render(
		lipgloss.Place(
			s.model.Width,
			s.model.Height,
			lipgloss.Center,
			lipgloss.Center,
			s.style.
				Padding(2, 2).
				BorderStyle(lipgloss.NormalBorder()).
				Render(content),
		),
	)
}

func (s *lobbyScreen) playerList() string {
	var playerList strings.Builder
	style := s.style

	for _, p := range s.model.Game.OrderedPlayers() {
		var listItem strings.Builder

		listItem.WriteString("* " + p.Name)
		if p.Name == s.model.Player.Name {
			listItem.WriteString(" (you)")
		}
		if p.IsHost() {
			listItem.WriteString(" (host)")
		}
		if !p.HasFaction() {
			listItem.WriteString(" (no faction)")
		}

		if p.HasFaction() {
			playerList.WriteString(style.Foreground(p.Faction.Color).Render(listItem.String()) + "\n")
		} else {
			playerList.WriteString(listItem.String() + "\n")
		}
	}

	return playerList.String()
}

func (s *lobbyScreen) footer() string {
	var sb strings.Builder
	colorList := make([]string, 0)

	for i, faction := range factions.All() {
		style := s.style
		word := style.Foreground(faction.Color).Render(faction.Name)

		if s.model.Game.IsFactionUsed(faction) {
			style = style.Italic(true)
			word = faction.Name + " (used)"
		}

		item := style.Render("Press '" + strconv.Itoa(i+1) + "' to choose " + word)
		colorList = append(colorList, item)
	}

	sb.WriteString(lipgloss.JoinVertical(lipgloss.Left, colorList...))

	sb.WriteString("\n")
	if s.model.Player.IsHost() {
		sb.WriteString("Press 's' to start the game.")
	} else {
		sb.WriteString("Waiting for host to start the game...")
	}
	sb.WriteString("\nPress 'ctrl+c' to quit.")

	return sb.String()
}
