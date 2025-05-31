package deck

import (
	"math/rand/v2"

	"github.com/ascii-arcade/moonrollers/factions"
)

func NewDeck() []*Crew {
	copiedDeck := make([]*Crew, len(deck))
	for i, c := range deck {
		copiedDeck[i] = &Crew{
			Ability:    c.Ability,
			Faction:    c.Faction,
			IsStarter:  c.IsStarter,
			Name:       c.Name,
			Objectives: c.Objectives,
		}
	}

	for i := len(copiedDeck) - 1; i > 0; i-- {
		j := rand.IntN(i + 1)
		copiedDeck[i], copiedDeck[j] = copiedDeck[j], copiedDeck[i]
	}

	return copiedDeck
}

var deck = []Crew{
	{
		Name:    "Aponi",
		Faction: factions.Blue,
		Objectives: []objective{
			{Type: DieReactor, Amount: 4, Hazard: true},
			{Type: DieThruster, Amount: 3},
			{Type: DieShield, Amount: 3, Hazard: true},
			{Type: DieWild, Amount: 2},
		},
		Ability: ability{Description: "If your rolling pool is 1-3 dice, each %reactor% or %wild% locked this roll is treated as 2 %reactors%."},
	},
	{
		Name:    "Vila",
		Faction: factions.Blue,
		Objectives: []objective{
			{Type: DieReactor, Amount: 4},
			{Type: DieReactor, Amount: 3},
			{Type: DieReactor, Amount: 2},
			{Type: DieReactor, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll no %reactors% you may re-roll 2 dice."},
	},
	{
		Name:    "Salatar",
		Faction: factions.Blue,
		Objectives: []objective{
			{Type: DieReactor, Amount: 4},
			{Type: DieThruster, Amount: 3, Hazard: true},
			{Type: DieShield, Amount: 3},
			{Type: DieDamage, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll no %reactors% you may re-roll 2 dice."},
	},
	{
		Name:    "Ada",
		Faction: factions.Blue,
		Objectives: []objective{
			{Type: DieReactor, Amount: 2},
			{Type: DieThruster, Amount: 2},
			{Type: DieShield, Amount: 1, Hazard: true},
		},
		Ability:   ability{Description: "May lock each %extra% as 2 %reactors%."},
		IsStarter: true,
	},
	{
		Name:    "Lee",
		Faction: factions.Blue,
		Objectives: []objective{
			{Type: DieReactor, Amount: 2},
			{Type: DieThruster, Amount: 2, Hazard: true},
			{Type: DieShield, Amount: 2},
			{Type: DieDamage, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll exactly 1 %reactor% you may lock it as %wild%."},
	},
	{
		Name:    "Lila",
		Faction: factions.Blue,
		Objectives: []objective{
			{Type: DieReactor, Amount: 3, Hazard: true},
			{Type: DieThruster, Amount: 3},
			{Type: DieShield, Amount: 2},
			{Type: DieDamage, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll 2+ %reactors% you may re-roll any number of dice."},
	},
	{
		Name:    "[REDACTED]",
		Faction: factions.Green,
		Objectives: []objective{
			{Type: DieShield, Amount: 4, Hazard: true},
			{Type: DieDamage, Amount: 3},
			{Type: DieThruster, Amount: 3},
			{Type: DieWild, Amount: 2},
		},
		Ability: ability{Description: "If your rolling pool is 1-3 dice, each %shield% or %wild% locked this roll is treated as 2 %shields%."},
	},
	{
		Name:    "Imdar",
		Faction: factions.Green,
		Objectives: []objective{
			{Type: DieShield, Amount: 4},
			{Type: DieShield, Amount: 3},
			{Type: DieShield, Amount: 2},
			{Type: DieShield, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll no %shields% you may draw 1 %hazard%."},
	},
	{
		Name:    "Namari",
		Faction: factions.Green,
		Objectives: []objective{
			{Type: DieShield, Amount: 4},
			{Type: DieDamage, Amount: 3, Hazard: true},
			{Type: DieThruster, Amount: 3},
			{Type: DieReactor, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "Gain 2 Prestige after busting."},
	},
	{
		Name:    "Ryle",
		Faction: factions.Green,
		Objectives: []objective{
			{Type: DieShield, Amount: 2},
			{Type: DieDamage, Amount: 2},
			{Type: DieThruster, Amount: 1},
		},
		Ability:   ability{Description: "May lock each %extra% as 2 %shields%."},
		IsStarter: true,
	},
	{
		Name:    "Bill",
		Faction: factions.Green,
		Objectives: []objective{
			{Type: DieShield, Amount: 2, Hazard: true},
			{Type: DieDamage, Amount: 2},
			{Type: DieThruster, Amount: 2},
			{Type: DieReactor, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll exactly 1 %shield% you may lock it as %wild%."},
	},
	{
		Name:    "AT-OK",
		Faction: factions.Green,
		Objectives: []objective{
			{Type: DieShield, Amount: 3, Hazard: true},
			{Type: DieDamage, Amount: 3},
			{Type: DieThruster, Amount: 2},
			{Type: DieReactor, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll 2+ %shields% you cannot bust on your next roll."},
	},
	{
		Name:    "Dr.Umbrage",
		Faction: factions.Orange,
		Objectives: []objective{
			{Type: DieDamage, Amount: 4, Hazard: true},
			{Type: DieShield, Amount: 3},
			{Type: DieReactor, Amount: 3},
			{Type: DieWild, Amount: 2},
		},
		Ability: ability{Description: "If your rolling pool is 1-3 dice, each %damage% or %wild% locked this roll is treated as 2 %damage%."},
	},
	{
		Name:    "Saghari",
		Faction: factions.Orange,
		Objectives: []objective{
			{Type: DieDamage, Amount: 4},
			{Type: DieDamage, Amount: 3},
			{Type: DieDamage, Amount: 2},
			{Type: DieDamage, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll no %damage%, roll 1 supply die and keep if %wild% or %extra%."},
	},
	{
		Name:    "Kary",
		Faction: factions.Orange,
		Objectives: []objective{
			{Type: DieDamage, Amount: 4},
			{Type: DieShield, Amount: 3, Hazard: true},
			{Type: DieReactor, Amount: 3},
			{Type: DieThruster, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "Any %damage% from your first roll may be treated as %extra%."},
	},
	{
		Name:    "Dana",
		Faction: factions.Orange,
		Objectives: []objective{
			{Type: DieDamage, Amount: 3},
			{Type: DieShield, Amount: 3},
			{Type: DieReactor, Amount: 1, Hazard: true},
		},
		Ability:   ability{Description: "May lock each %extra% as 2 %damage%."},
		IsStarter: true,
	},
	{
		Name:    "Tantin",
		Faction: factions.Orange,
		Objectives: []objective{
			{Type: DieDamage, Amount: 2},
			{Type: DieShield, Amount: 2, Hazard: true},
			{Type: DieReactor, Amount: 2},
			{Type: DieThruster, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll exactly 1 %damage% you may lock it as %wild%."},
	},
	{
		Name:    "Ryan",
		Faction: factions.Orange,
		Objectives: []objective{
			{Type: DieDamage, Amount: 3, Hazard: true},
			{Type: DieShield, Amount: 3},
			{Type: DieReactor, Amount: 2},
			{Type: DieThruster, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll 2+ %damage%, roll 2 supply dice and keep any that are %wild%."},
	},
	{
		Name:    "Moro",
		Faction: factions.Purple,
		Objectives: []objective{
			{Type: DieReactor, Amount: 4},
			{Type: DieDamage, Amount: 3, Hazard: true},
			{Type: DieShield, Amount: 3},
			{Type: DieThruster, Amount: 2, Hazard: true},
		},
		Ability: ability{Description: "If your pool is 1-3 dice, you cannot bust if you roll at least one %extra%."},
	},
	{
		Name:    "Vanta",
		Faction: factions.Purple,
		Objectives: []objective{
			{Type: DieWild, Amount: 3},
			{Type: DieWild, Amount: 2},
			{Type: DieWild, Amount: 1},
			{Type: DieDamage, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll no %extra% you may lock any 1 die as %wild%."},
	},
	{
		Name:    "Meg",
		Faction: factions.Purple,
		Objectives: []objective{
			{Type: DieThruster, Amount: 4, Hazard: true},
			{Type: DieDamage, Amount: 3},
			{Type: DieShield, Amount: 3},
			{Type: DieReactor, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "1 %wild% from your first roll may be saved for your next roll."},
	},
	{
		Name:    "Sella",
		Faction: factions.Purple,
		Objectives: []objective{
			{Type: DieThruster, Amount: 2},
			{Type: DieReactor, Amount: 2},
			{Type: DieShield, Amount: 1},
		},
		Ability:   ability{Description: "If you roll exactly 1 %extra% you may lock it as %wild%."},
		IsStarter: true,
	},
	{
		Name:    "FT-1000",
		Faction: factions.Purple,
		Objectives: []objective{
			{Type: DieShield, Amount: 3},
			{Type: DieThruster, Amount: 2, Hazard: true},
			{Type: DieDamage, Amount: 2, Hazard: true},
			{Type: DieReactor, Amount: 2},
		},
		Ability: ability{Description: "If you roll exactly 1 %wild%, you may treat it as %extra%."},
	},
	{
		Name:    "Avari",
		Faction: factions.Purple,
		Objectives: []objective{
			{Type: DieDamage, Amount: 3, Hazard: true},
			{Type: DieReactor, Amount: 3},
			{Type: DieShield, Amount: 2},
			{Type: DieThruster, Amount: 2, Hazard: true},
		},
		Ability: ability{Description: "If you roll all %wilds%, gain 3 dice for your next roll."},
	},
	{
		Name:    "Sol",
		Faction: factions.Yellow,
		Objectives: []objective{
			{Type: DieThruster, Amount: 4, Hazard: true},
			{Type: DieReactor, Amount: 3},
			{Type: DieDamage, Amount: 3, Hazard: true},
			{Type: DieWild, Amount: 2},
		},
		Ability: ability{Description: "If your rolling pool is 1-3 dice, each %thruster% or %wild% locked this roll is treated as 2 %thrusters%."},
	},
	{
		Name:    "B3-AR",
		Faction: factions.Yellow,
		Objectives: []objective{
			{Type: DieThruster, Amount: 4},
			{Type: DieThruster, Amount: 3},
			{Type: DieThruster, Amount: 2},
			{Type: DieThruster, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll 3+ %thrusters% finish the current requirement."},
	},
	{
		Name:    "Kal",
		Faction: factions.Yellow,
		Objectives: []objective{
			{Type: DieThruster, Amount: 4},
			{Type: DieReactor, Amount: 3, Hazard: true},
			{Type: DieDamage, Amount: 3},
			{Type: DieShield, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "Your rolling pool starts with 6 dice."},
	},
	{
		Name:    "Nella",
		Faction: factions.Yellow,
		Objectives: []objective{
			{Type: DieThruster, Amount: 2},
			{Type: DieReactor, Amount: 2},
			{Type: DieDamage, Amount: 1, Hazard: true},
		},
		Ability:   ability{Description: "May lock each %extra% as 2 %thrusters%."},
		IsStarter: true,
	},
	{
		Name:    "Zek",
		Faction: factions.Yellow,
		Objectives: []objective{
			{Type: DieThruster, Amount: 2},
			{Type: DieReactor, Amount: 2, Hazard: true},
			{Type: DieDamage, Amount: 2},
			{Type: DieShield, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll exactly 1 %thruster% you may lock it as %wild%."},
	},
	{
		Name:    "Myla",
		Faction: factions.Yellow,
		Objectives: []objective{
			{Type: DieThruster, Amount: 3, Hazard: true},
			{Type: DieReactor, Amount: 3},
			{Type: DieDamage, Amount: 2},
			{Type: DieShield, Amount: 1, Hazard: true},
		},
		Ability: ability{Description: "If you roll 2+ %thrusters% you may treat 1 of your dice as %extra%."},
	},
}
