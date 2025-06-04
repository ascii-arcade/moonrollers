package board

import (
	"strconv"
	"strings"

	"github.com/ascii-arcade/moonrollers/colors"
	"github.com/ascii-arcade/moonrollers/games"
	"github.com/charmbracelet/lipgloss"
)

const (
	scoreboardPip = "■"
)

type scoreboard struct {
	model   *Model
	players []*games.Player
	short   bool
	style   lipgloss.Style
}

func newScoreboard(model *Model) scoreboard {
	return scoreboard{
		model:   model,
		players: model.Game.OrderedPlayers(),
		short:   false,
		style:   model.style,
	}
}

func (s *scoreboard) render() string {
	if s.short {
		out := make([]string, 0)
		for _, p := range s.players {
			points := s.style.
				Foreground(lipgloss.Color(p.Faction.Color)).
				Render(p.Name + ": " + strconv.Itoa(p.Points))
			out = append(out, points)
		}

		return lipgloss.JoinVertical(lipgloss.Left, s.style.Padding(1).Render(strings.Join(out, "\n")))
	}

	rows := make([]string, 0)
	for row := range 10 {
		var rowCells []string
		for col := range 5 {
			rowCells = append(rowCells, s.pointCell(row, col))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, rowCells...))
	}

	morePointsRow := make([]string, 0)
	for playerIndex := range 5 {
		morePointsRow = append(morePointsRow, s.additionalPointsCell(playerIndex))
	}
	rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, morePointsRow...))

	return s.style.Render(strings.Join(rows, "\n"))
}

func (s *scoreboard) additionalPointsCell(playerIndex int) string {
	if playerIndex >= len(s.players) {
		return s.emptyCellStyle().Render("")
	}

	points := s.players[playerIndex].Points

	if points < 50 {
		return s.emptyCellStyle().Render("")
	}

	output := make([]string, 0)

	for range points / 50 {
		output = append(output, s.style.Foreground(s.players[playerIndex].Faction.Color).Render(scoreboardPip))
	}

	return s.populatedCellStyle().Render(strings.Join(output, " "))
}

func (s *scoreboard) pointCell(row int, col int) string {
	points := (row)*5 + col + 1
	var pips []string
	for _, player := range s.players {
		if player.Points%50 == points {
			pip := s.style.Foreground(player.Faction.Color).Render(scoreboardPip)
			pips = append(pips, pip)
		}
	}
	if len(pips) == 0 {
		return s.emptyCellStyle().Render(strconv.Itoa(points))
	}
	return s.populatedCellStyle().Render(strings.Join(pips, ""))
}

func (s *scoreboard) populatedCellStyle() lipgloss.Style {
	return s.style.
		Border(lipgloss.NormalBorder(), true).
		BorderForeground(colors.Border).
		Width(5).
		Height(1).
		Align(lipgloss.Center)
}

func (s *scoreboard) emptyCellStyle() lipgloss.Style {
	return s.populatedCellStyle().Foreground(colors.Border)
}
