package games

type Settings struct {
	UseStarterCards bool
}

func NewSettings() Settings {
	return Settings{
		UseStarterCards: true,
	}
}
