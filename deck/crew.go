package deck

import (
	"github.com/ascii-arcade/moonrollers/dice"
	"github.com/ascii-arcade/moonrollers/factions"
)

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

var allCrew = []Crew{
	{
		Name:    "Aponi",
		Faction: factions.Blue,
		Ability: ability{Description: "If your rolling pool is 1-3 dice, each %reactor% or %wild% locked this roll is treated as 2 %reactors%."},
		Objectives: []objective{
			{Type: dice.DieReactor, Amount: 4, Hazard: true},
			{Type: dice.DieThruster, Amount: 3},
			{Type: dice.DieShield, Amount: 3, Hazard: true},
			{Type: dice.DieWild, Amount: 2},
		},
	},
	{
		Name:    "Vila",
		Faction: factions.Blue,
		Ability: ability{Description: "If you roll no %reactors% you may re-roll 2 dice."},
		Objectives: []objective{
			{Type: dice.DieReactor, Amount: 4},
			{Type: dice.DieReactor, Amount: 3},
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieReactor, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Salatar",
		Faction: factions.Blue,
		Ability: ability{Description: "If you roll no %reactors% you may re-roll 2 dice."},
		Objectives: []objective{
			{Type: dice.DieReactor, Amount: 4},
			{Type: dice.DieThruster, Amount: 3, Hazard: true},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieDamage, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Ada",
		Faction: factions.Blue,
		Ability: ability{Description: "May lock each %extra% as 2 %reactors%."},
		Objectives: []objective{
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieShield, Amount: 1, Hazard: true},
		},
		IsStarter: true,
	},
	{
		Name:    "Lee",
		Faction: factions.Blue,
		Ability: ability{Description: "If you roll exactly 1 %reactor% you may lock it as %wild%."},
		Objectives: []objective{
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieThruster, Amount: 2, Hazard: true},
			{Type: dice.DieShield, Amount: 2},
			{Type: dice.DieDamage, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Lila",
		Faction: factions.Blue,
		Ability: ability{Description: "If you roll 2+ %reactors% you may re-roll any number of dice."},
		Objectives: []objective{
			{Type: dice.DieReactor, Amount: 3, Hazard: true},
			{Type: dice.DieThruster, Amount: 3},
			{Type: dice.DieShield, Amount: 2},
			{Type: dice.DieDamage, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "[REDACTED]",
		Faction: factions.Green,
		Ability: ability{Description: "If your rolling pool is 1-3 dice, each %shield% or %wild% locked this roll is treated as 2 %shields%."},
		Objectives: []objective{
			{Type: dice.DieShield, Amount: 4, Hazard: true},
			{Type: dice.DieDamage, Amount: 3},
			{Type: dice.DieThruster, Amount: 3},
			{Type: dice.DieWild, Amount: 2},
		},
	},
	{
		Name:    "Imdar",
		Faction: factions.Green,
		Ability: ability{Description: "If you roll no %shields% you may draw 1 %hazard%."},
		Objectives: []objective{
			{Type: dice.DieShield, Amount: 4},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieShield, Amount: 2},
			{Type: dice.DieShield, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Namari",
		Faction: factions.Green,
		Ability: ability{Description: "Gain 2 Prestige after busting."},
		Objectives: []objective{
			{Type: dice.DieShield, Amount: 4},
			{Type: dice.DieDamage, Amount: 3, Hazard: true},
			{Type: dice.DieThruster, Amount: 3},
			{Type: dice.DieReactor, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Ryle",
		Faction: factions.Green,
		Ability: ability{Description: "May lock each %extra% as 2 %shields%."},
		Objectives: []objective{
			{Type: dice.DieShield, Amount: 2},
			{Type: dice.DieDamage, Amount: 2},
			{Type: dice.DieThruster, Amount: 1},
		},
		IsStarter: true,
	},
	{
		Name:    "Bill",
		Faction: factions.Green,
		Ability: ability{Description: "If you roll exactly 1 %shield% you may lock it as %wild%."},
		Objectives: []objective{
			{Type: dice.DieShield, Amount: 2, Hazard: true},
			{Type: dice.DieDamage, Amount: 2},
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieReactor, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "AT-OK",
		Faction: factions.Green,
		Ability: ability{Description: "If you roll 2+ %shields% you cannot bust on your next roll."},
		Objectives: []objective{
			{Type: dice.DieShield, Amount: 3, Hazard: true},
			{Type: dice.DieDamage, Amount: 3},
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieReactor, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Dr.Umbrage",
		Faction: factions.Orange,
		Ability: ability{Description: "If your rolling pool is 1-3 dice, each %damage% or %wild% locked this roll is treated as 2 %damage%."},
		Objectives: []objective{
			{Type: dice.DieDamage, Amount: 4, Hazard: true},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieReactor, Amount: 3},
			{Type: dice.DieWild, Amount: 2},
		},
	},
	{
		Name:    "Saghari",
		Faction: factions.Orange,
		Ability: ability{Description: "If you roll no %damage%, roll 1 supply die and keep if %wild% or %extra%."},
		Objectives: []objective{
			{Type: dice.DieDamage, Amount: 4},
			{Type: dice.DieDamage, Amount: 3},
			{Type: dice.DieDamage, Amount: 2},
			{Type: dice.DieDamage, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Kary",
		Faction: factions.Orange,
		Ability: ability{Description: "Any %damage% from your first roll may be treated as %extra%."},
		Objectives: []objective{
			{Type: dice.DieDamage, Amount: 4},
			{Type: dice.DieShield, Amount: 3, Hazard: true},
			{Type: dice.DieReactor, Amount: 3},
			{Type: dice.DieThruster, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Dana",
		Faction: factions.Orange,
		Ability: ability{Description: "May lock each %extra% as 2 %damage%."},
		Objectives: []objective{
			{Type: dice.DieDamage, Amount: 3},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieReactor, Amount: 1, Hazard: true},
		},
		IsStarter: true,
	},
	{
		Name:    "Tantin",
		Faction: factions.Orange,
		Ability: ability{Description: "If you roll exactly 1 %damage% you may lock it as %wild%."},
		Objectives: []objective{
			{Type: dice.DieDamage, Amount: 2},
			{Type: dice.DieShield, Amount: 2, Hazard: true},
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieThruster, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Ryan",
		Faction: factions.Orange,
		Ability: ability{Description: "If you roll 2+ %damage%, roll 2 supply dice and keep any that are %wild%."},
		Objectives: []objective{
			{Type: dice.DieDamage, Amount: 3, Hazard: true},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieThruster, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Moro",
		Faction: factions.Purple,
		Ability: ability{Description: "If your pool is 1-3 dice, you cannot bust if you roll at least one %extra%."},
		Objectives: []objective{
			{Type: dice.DieReactor, Amount: 4},
			{Type: dice.DieDamage, Amount: 3, Hazard: true},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieThruster, Amount: 2, Hazard: true},
		},
	},
	{
		Name:    "Vanta",
		Faction: factions.Purple,
		Ability: ability{Description: "If you roll no %extra% you may lock any 1 die as %wild%."},
		Objectives: []objective{
			{Type: dice.DieWild, Amount: 3},
			{Type: dice.DieWild, Amount: 2},
			{Type: dice.DieWild, Amount: 1},
			{Type: dice.DieDamage, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Meg",
		Faction: factions.Purple,
		Ability: ability{Description: "1 %wild% from your first roll may be saved for your next roll."},
		Objectives: []objective{
			{Type: dice.DieThruster, Amount: 4, Hazard: true},
			{Type: dice.DieDamage, Amount: 3},
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieReactor, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Sella",
		Faction: factions.Purple,
		Ability: ability{Description: "If you roll exactly 1 %extra% you may lock it as %wild%."},
		Objectives: []objective{
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieShield, Amount: 1},
		},
		IsStarter: true,
	},
	{
		Name:    "FT-1000",
		Faction: factions.Purple,
		Ability: ability{Description: "If you roll exactly 1 %wild%, you may treat it as %extra%."},
		Objectives: []objective{
			{Type: dice.DieShield, Amount: 3},
			{Type: dice.DieThruster, Amount: 2, Hazard: true},
			{Type: dice.DieDamage, Amount: 2, Hazard: true},
			{Type: dice.DieReactor, Amount: 2},
		},
	},
	{
		Name:    "Avari",
		Faction: factions.Purple,
		Ability: ability{Description: "If you roll all %wilds%, gain 3 dice for your next roll."},
		Objectives: []objective{
			{Type: dice.DieDamage, Amount: 3, Hazard: true},
			{Type: dice.DieReactor, Amount: 3},
			{Type: dice.DieShield, Amount: 2},
			{Type: dice.DieThruster, Amount: 2, Hazard: true},
		},
	},
	{
		Name:    "Sol",
		Faction: factions.Yellow,
		Ability: ability{Description: "If your rolling pool is 1-3 dice, each %thruster% or %wild% locked this roll is treated as 2 %thrusters%."},
		Objectives: []objective{
			{Type: dice.DieThruster, Amount: 4, Hazard: true},
			{Type: dice.DieReactor, Amount: 3},
			{Type: dice.DieDamage, Amount: 3, Hazard: true},
			{Type: dice.DieWild, Amount: 2},
		},
	},
	{
		Name:    "B3-AR",
		Faction: factions.Yellow,
		Ability: ability{Description: "If you roll 3+ %thrusters% finish the current requirement."},
		Objectives: []objective{
			{Type: dice.DieThruster, Amount: 4},
			{Type: dice.DieThruster, Amount: 3},
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieThruster, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Kal",
		Faction: factions.Yellow,
		Ability: ability{Description: "Your rolling pool starts with 6 dice."},
		Objectives: []objective{
			{Type: dice.DieThruster, Amount: 4},
			{Type: dice.DieReactor, Amount: 3, Hazard: true},
			{Type: dice.DieDamage, Amount: 3},
			{Type: dice.DieShield, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Nella",
		Faction: factions.Yellow,
		Ability: ability{Description: "May lock each %extra% as 2 %thrusters%."},
		Objectives: []objective{
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieReactor, Amount: 2},
			{Type: dice.DieDamage, Amount: 1, Hazard: true},
		},
		IsStarter: true,
	},
	{
		Name:    "Zek",
		Faction: factions.Yellow,
		Ability: ability{Description: "If you roll exactly 1 %thruster% you may lock it as %wild%."},
		Objectives: []objective{
			{Type: dice.DieThruster, Amount: 2},
			{Type: dice.DieReactor, Amount: 2, Hazard: true},
			{Type: dice.DieDamage, Amount: 2},
			{Type: dice.DieShield, Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Myla",
		Faction: factions.Yellow,
		Ability: ability{Description: "If you roll 2+ %thrusters% you may treat 1 of your dice as %extra%."},
		Objectives: []objective{
			{Type: dice.DieThruster, Amount: 3, Hazard: true},
			{Type: dice.DieReactor, Amount: 3},
			{Type: dice.DieDamage, Amount: 2},
			{Type: dice.DieShield, Amount: 1, Hazard: true},
		},
	},
}
