package games

import (
	"context"

	"github.com/ascii-arcade/moonrollers/deck"
	"github.com/ascii-arcade/moonrollers/factions"
	"github.com/ascii-arcade/moonrollers/generaterandom"
	"github.com/ascii-arcade/moonrollers/language"
	"github.com/charmbracelet/ssh"
)

var players = make(map[string]*Player)

func NewPlayer(ctx context.Context, sess ssh.Session, langPref *language.LanguagePreference) *Player {
	player, exists := players[sess.User()]

	if exists {
		player.UpdateChan = make(chan int)
		player.connected = true
		player.ctx = ctx
	} else {
		crew := make(map[string]*deck.Crew)
		crewCount := make(map[string]int)
		for _, faction := range factions.All() {
			crewCount[faction.Name] = 0
		}

		player = &Player{
			Name:               generaterandom.Name(langPref.Lang),
			Crew:               crew,
			CrewCount:          crewCount,
			Points:             0,
			UpdateChan:         make(chan int),
			LanguagePreference: langPref,
			Sess:               sess,
			connected:          true,
			ctx:                ctx,
		}
		players[sess.User()] = player
	}

	go func() {
		<-player.ctx.Done()
		player.connected = false
		for _, fn := range player.onDisconnect {
			fn()
		}
	}()

	return player
}

func RemovePlayer(player *Player) {
	if _, exists := players[player.Sess.User()]; exists {
		close(player.UpdateChan)
		delete(players, player.Sess.User())
	}
}

func GetPlayerCount() int {
	return len(players)
}

func GetConnectedPlayerCount() int {
	count := 0
	for _, player := range players {
		if player.connected {
			count++
		}
	}
	return count
}
