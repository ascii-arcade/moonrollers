package deck

import "github.com/ascii-arcade/moonrollers/factions"

type ability struct {
	Description string
}

type Crew struct {
	Ability    ability
	Faction    factions.Faction
	IsStarter  bool
	Name       string
	Objectives []objective
}
