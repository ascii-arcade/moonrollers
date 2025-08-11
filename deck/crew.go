package deck

import (
	"github.com/ascii-arcade/moonrollers/dice"
	"github.com/ascii-arcade/moonrollers/factions"
)

type Crew struct {
	Faction    factions.Faction
	ID         string
	IsStarter  bool
	Name       string
	Objectives []Objective
}

var allCrew = []Crew{
	{
		Name:    "Aponi",
		ID:      "aponi",
		Faction: factions.Blue,
		Objectives: []Objective{
			{Type: dice.DieReactor, Amount: 4, Hazard: true},
			{Type: dice.DieThruster, Amount: 3},
			{Type: dice.DieShield, Amount: 3, Hazard: true},
			{Type: dice.DieWild, Amount: 2},
		},
	},
	{
		Name:    "Vila",
		ID:      "vila",
		Faction: factions.Blue,
		Objectives: []Objective{
			{Type: dice.DieReactor, Amount: 4},
			{Type: dice.DieReactor, Amount: 3},
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieReactor, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Salatar",
		ID:      "salatar",
		Faction: factions.Blue,
		Objectives: []Objective{
			{Type: dice.DieReactor, Amount: 4},
			{Type: dice.DieThruster, Amount: 3, Hazard: true},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieDamage, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Ada",
		ID:      "ada",
		Faction: factions.Blue,
		Objectives: []Objective{
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieShield, Amount: 1, Hazard: true},
		},
		IsStarter: true,
	},
	{
		Name:    "Lee",
		ID:      "lee",
		Faction: factions.Blue,
		Objectives: []Objective{
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieThruster, Amount: 2, Hazard: true},
			{Type: dice.DieShield, Amount: 2},
			{Type: dice.DieDamage, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Lila",
		ID:      "lila",
		Faction: factions.Blue,
		Objectives: []Objective{
			{Type: dice.DieReactor, Amount: 3, Hazard: true},
			{Type: dice.DieThruster, Amount: 3},
			{Type: dice.DieShield, Amount: 2},
			{Type: dice.DieDamage, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "[REDACTED]",
		ID:      "redacted",
		Faction: factions.Green,
		Objectives: []Objective{
			{Type: dice.DieShield, Amount: 4, Hazard: true},
			{Type: dice.DieDamage, Amount: 3},
			{Type: dice.DieThruster, Amount: 3},
			{Type: dice.DieWild, Amount: 2},
		},
	},
	{
		Name:    "Imdar",
		ID:      "imdar",
		Faction: factions.Green,
		Objectives: []Objective{
			{Type: dice.DieShield, Amount: 4},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieShield, Amount: 2},
			{Type: dice.DieShield, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Namari",
		ID:      "namari",
		Faction: factions.Green,
		Objectives: []Objective{
			{Type: dice.DieShield, Amount: 4},
			{Type: dice.DieDamage, Amount: 3, Hazard: true},
			{Type: dice.DieThruster, Amount: 3},
			{Type: dice.DieReactor, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Ryle",
		ID:      "ryle",
		Faction: factions.Green,
		Objectives: []Objective{
			{Type: dice.DieShield, Amount: 2},
			{Type: dice.DieDamage, Amount: 2},
			{Type: dice.DieThruster, Amount: 1},
		},
		IsStarter: true,
	},
	{
		Name:    "Bill",
		ID:      "bill",
		Faction: factions.Green,
		Objectives: []Objective{
			{Type: dice.DieShield, Amount: 2, Hazard: true},
			{Type: dice.DieDamage, Amount: 2},
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieReactor, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "AT-OK",
		ID:      "at-ok",
		Faction: factions.Green,
		Objectives: []Objective{
			{Type: dice.DieShield, Amount: 3, Hazard: true},
			{Type: dice.DieDamage, Amount: 3},
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieReactor, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Dr.Umbrage",
		ID:      "drumbrage",
		Faction: factions.Orange,
		Objectives: []Objective{
			{Type: dice.DieDamage, Amount: 4, Hazard: true},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieReactor, Amount: 3},
			{Type: dice.DieWild, Amount: 2},
		},
	},
	{
		Name:    "Saghari",
		ID:      "saghari",
		Faction: factions.Orange,
		Objectives: []Objective{
			{Type: dice.DieDamage, Amount: 4},
			{Type: dice.DieDamage, Amount: 3},
			{Type: dice.DieDamage, Amount: 2},
			{Type: dice.DieDamage, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Kary",
		ID:      "kary",
		Faction: factions.Orange,
		Objectives: []Objective{
			{Type: dice.DieDamage, Amount: 4},
			{Type: dice.DieShield, Amount: 3, Hazard: true},
			{Type: dice.DieReactor, Amount: 3},
			{Type: dice.DieThruster, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Dana",
		ID:      "dana",
		Faction: factions.Orange,
		Objectives: []Objective{
			{Type: dice.DieDamage, Amount: 3},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieReactor, Amount: 1, Hazard: true},
		},
		IsStarter: true,
	},
	{
		Name:    "Tantin",
		ID:      "tantin",
		Faction: factions.Orange,
		Objectives: []Objective{
			{Type: dice.DieDamage, Amount: 2},
			{Type: dice.DieShield, Amount: 2, Hazard: true},
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieThruster, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Ryan",
		ID:      "ryan",
		Faction: factions.Orange,
		Objectives: []Objective{
			{Type: dice.DieDamage, Amount: 3, Hazard: true},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieThruster, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Moro",
		ID:      "moro",
		Faction: factions.Purple,
		Objectives: []Objective{
			{Type: dice.DieReactor, Amount: 4},
			{Type: dice.DieDamage, Amount: 3, Hazard: true},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieThruster, Amount: 2, Hazard: true},
		},
	},
	{
		Name:    "Vanta",
		ID:      "vanta",
		Faction: factions.Purple,
		Objectives: []Objective{
			{Type: dice.DieWild, Amount: 3},
			{Type: dice.DieWild, Amount: 2},
			{Type: dice.DieWild, Amount: 1},
			{Type: dice.DieDamage, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Meg",
		ID:      "meg",
		Faction: factions.Purple,
		Objectives: []Objective{
			{Type: dice.DieThruster, Amount: 4, Hazard: true},
			{Type: dice.DieDamage, Amount: 3},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieReactor, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Sella",
		ID:      "sella",
		Faction: factions.Purple,
		Objectives: []Objective{
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieShield, Amount: 1},
		},
		IsStarter: true,
	},
	{
		Name:    "FT-1000",
		ID:      "ft1000",
		Faction: factions.Purple,
		Objectives: []Objective{
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieThruster, Amount: 2, Hazard: true},
			{Type: dice.DieDamage, Amount: 2, Hazard: true},
			{Type: dice.DieReactor, Amount: 2},
		},
	},
	{
		Name:    "Avari",
		ID:      "avari",
		Faction: factions.Purple,
		Objectives: []Objective{
			{Type: dice.DieDamage, Amount: 3, Hazard: true},
			{Type: dice.DieReactor, Amount: 3},
			{Type: dice.DieShield, Amount: 2},
			{Type: dice.DieThruster, Amount: 2, Hazard: true},
		},
	},
	{
		Name:    "Sol",
		ID:      "sol",
		Faction: factions.Yellow,
		Objectives: []Objective{
			{Type: dice.DieThruster, Amount: 4, Hazard: true},
			{Type: dice.DieReactor, Amount: 3},
			{Type: dice.DieDamage, Amount: 3, Hazard: true},
			{Type: dice.DieWild, Amount: 2},
		},
	},
	{
		Name:    "B3-AR",
		ID:      "b3ar",
		Faction: factions.Yellow,
		Objectives: []Objective{
			{Type: dice.DieThruster, Amount: 4},
			{Type: dice.DieThruster, Amount: 3},
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieThruster, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Kal",
		ID:      "kal",
		Faction: factions.Yellow,
		Objectives: []Objective{
			{Type: dice.DieThruster, Amount: 4},
			{Type: dice.DieReactor, Amount: 3, Hazard: true},
			{Type: dice.DieDamage, Amount: 3},
			{Type: dice.DieShield, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Nella",
		ID:      "nella",
		Faction: factions.Yellow,
		Objectives: []Objective{
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieDamage, Amount: 1, Hazard: true},
		},
		IsStarter: true,
	},
	{
		Name:    "Zek",
		ID:      "zek",
		Faction: factions.Yellow,
		Objectives: []Objective{
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieReactor, Amount: 2, Hazard: true},
			{Type: dice.DieDamage, Amount: 2},
			{Type: dice.DieShield, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Myla",
		ID:      "myla",
		Faction: factions.Yellow,
		Objectives: []Objective{
			{Type: dice.DieThruster, Amount: 3, Hazard: true},
			{Type: dice.DieReactor, Amount: 3},
			{Type: dice.DieDamage, Amount: 2},
			{Type: dice.DieShield, Amount: 1, Hazard: true},
		},
	},
}
