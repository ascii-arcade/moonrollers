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
		switch {
		case keys.LobbyStartGame.TriggeredBy(msg.String()):
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
		case keys.LobbyJoinFaction.TriggeredBy(msg.String()):
			i, _ := strconv.Atoi(msg.String())

			if i == 0 {
				_ = s.model.Game.SetFaction(s.model.Player, nil)
				return s.model, nil
			}

			faction := factions.All()[i-1]
			if !s.model.Game.IsFactionUsed(faction) {
				_ = s.model.Game.SetFaction(s.model.Player, &faction)
			}
		case keys.LobbySettings.TriggeredBy(msg.String()):
			if s.model.Player.IsHost() {
				settingsScreen := s.model.newSettingsScreen()
				settingsScreen.Init()
				s.model.screen = settingsScreen
			}
		}
	}

	return s.model, nil
}

func (s *lobbyScreen) View() string {
	style := s.style.Width(s.model.width / 2)

	header := s.model.Game.Code
	playerList := s.style.Render(s.playerList())
	spectators := s.style.Render(s.spectators())
	footer := s.style.Render(s.footer())

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		style.Align(lipgloss.Center).MarginBottom(2).Render(header),
		style.Render(playerList),
		spectators,
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

func (s *lobbyScreen) spectators() string {
	style := s.style
	numberOfSpectators := 0
	for _, p := range s.model.Game.OrderedPlayers() {
		if !p.HasFaction() {
			numberOfSpectators++
		}
	}
	return style.Render(fmt.Sprintf("%d spectator%s\n", numberOfSpectators, func() string {
		if numberOfSpectators != 1 {
			return "s"
		}
		return ""
	}()))
}

func (s *lobbyScreen) playerList() string {
	var playerList strings.Builder
	style := s.style

	players := s.model.Game.OrderedPlayers()
	for i := range 5 {
		var listItem strings.Builder
		if len(players) >= i+1 {
			p := players[i]

			if !p.HasFaction() {
				goto SPECTATING
			}

			listItem.WriteString("* " + p.Name)
			if p.Name == s.model.Player.Name {
				fmt.Fprintf(&listItem, " (%s)", s.model.lang().Get("board", "player_list_you"))
			}
			if p.IsHost() {
				fmt.Fprintf(&listItem, " (%s)", s.model.lang().Get("board", "player_list_host"))
			}
			playerList.WriteString(style.Foreground(p.Faction.Color).Render(listItem.String()) + "\n")
			continue
		}
	SPECTATING:
		playerList.WriteString(style.Foreground(lipgloss.Color("#ffffff")).Render("* "+s.model.lang().Get("board", "no_player")) + "\n")
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

	sb.WriteString("\n\n")
	sb.WriteString("Press '0' to spectate")
	sb.WriteString("\n\n")

	if s.model.Player.IsHost() {
		err := s.model.Game.IsPlayerCountOk()
		if err == nil {
			fmt.Fprintf(&sb, s.model.lang().Get("board", "press_to_start"), keys.LobbyStartGame.String(s.style))
		} else {
			errorMessage := s.model.lang().Get("error", err.Error())
			sb.WriteString(s.style.Foreground(colors.Error).Render(errorMessage))
		}
		sb.WriteString("\n")
	} else {
		sb.WriteString(s.model.lang().Get("board", "waiting_for_start"))
	}
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf(s.model.lang().Get("settings", "press_to_open"), keys.LobbySettings.String(s.style)))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf(s.model.lang().Get("global", "quit"), keys.ExitApplication.String(s.style)))

	return sb.String()
}
