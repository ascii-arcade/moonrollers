package games

import (
	"context"

	"github.com/ascii-arcade/moonrollers/generaterandom"
	"github.com/ascii-arcade/moonrollers/language"
	"github.com/charmbracelet/ssh"
)

var players = make(map[string]*Player)

func NewPlayer(ctx context.Context, sess ssh.Session, langPref *language.LanguagePreference) *Player {
	player, exists := players[sess.User()]
	if exists {
		player.UpdateChan = make(chan struct{})
		player.connected = true
		player.ctx = ctx

		goto RETURN
	}

	player = &Player{
		Name:               generaterandom.Name(langPref.Lang),
		Points:             0,
		UpdateChan:         make(chan struct{}),
		LanguagePreference: langPref,
		Sess:               sess,
		connected:          true,
		ctx:                ctx,
	}
	players[sess.User()] = player

RETURN:
	go func() {
		<-player.ctx.Done()
		player.connected = false
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
