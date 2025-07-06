package games

type Settings struct {
	CardsOfAFactionToWin int
	UseStarterCards bool
}

func NewSettings() Settings {
	return Settings{
		CardsOfAFactionToWin: 3,
		UseStarterCards: true,
	}
}
