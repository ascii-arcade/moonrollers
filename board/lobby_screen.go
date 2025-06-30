package board

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/ascii-arcade/moonrollers/factions"
	"github.com/ascii-arcade/moonrollers/keys"
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

func (s *lobbyScreen) Update(msg tea.Msg) (any, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.model.height, s.model.width = msg.Height, msg.Width
		return s.model, nil

	case tea.KeyMsg:
		if keys.LobbyStartGame.TriggeredBy(msg.String()) {
			allHaveColor := true
			for _, p := range s.model.Game.OrderedPlayers() {
				if !p.HasFaction() {
					allHaveColor = false
					break
				}
			}
			if s.model.Player.IsHost() && allHaveColor {
				_ = s.model.Game.Begin()
			}
		}
		if keys.LobbyJoinFaction.TriggeredBy(msg.String()) {
			i, err := strconv.Atoi(msg.String())
			if err != nil {
				return s.model, nil
			} else {
				faction := factions.All()[i-1]
				if !s.model.Game.IsFactionUsed(faction) {
					_ = s.model.Game.SetFaction(s.model.Player, &faction)
				}
			}
		}
	}

	return s.model, nil
}

func (s *lobbyScreen) View() string {
	style := s.style.Width(s.model.width / 2)

	header := s.model.Game.Code
	playerList := s.style.Render(s.playerList())
	footer := s.style.Render(s.footer())

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		style.Align(lipgloss.Center).MarginBottom(2).Render(header),
		style.Render(playerList),
		style.Render(footer),
	)

	return s.style.Width(s.model.width).Height(s.model.height).Render(
		lipgloss.Place(
			s.model.width,
			s.model.height,
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
			listItem.WriteString(fmt.Sprintf(" (%s)", s.model.lang().Get("board", "player_list_you")))
		}
		if p.IsHost() {
			listItem.WriteString(fmt.Sprintf(" (%s)", s.model.lang().Get("board", "player_list_host")))
		}
		if !p.HasFaction() {
			listItem.WriteString(fmt.Sprintf(" (%s)", s.model.lang().Get("board", "no_faction")))
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
			word = fmt.Sprintf("%s (%s)", faction.Name, s.model.lang().Get("board", "used"))
		}

		item := style.Render(fmt.Sprintf(s.model.lang().Get("board", "choose_faction"), keys.LobbyJoinFaction.IndexedString(i, s.style), word))
		colorList = append(colorList, item)
	}

	sb.WriteString(lipgloss.JoinVertical(lipgloss.Left, colorList...))

	sb.WriteString("\n")
	if s.model.Player.IsHost() {
		err := s.model.Game.IsPlayerCountOk()
		if err == nil {
			sb.WriteString(fmt.Sprintf(s.model.lang().Get("board", "press_to_start"), keys.LobbyStartGame.String(s.style)))
		} else {
			errorMessage := s.model.lang().Get("error", err.Error())
			sb.WriteString(s.style.Foreground(colors.Error).Render(errorMessage))
		}
	} else {
		sb.WriteString(s.model.lang().Get("board", "waiting_for_start"))
	}
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf(s.model.lang().Get("global", "quit"), keys.ExitApplication.String(s.style)))

	return sb.String()
}
