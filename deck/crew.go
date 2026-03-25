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
	Objectives []*Objective
	Modifier   func(pool *dice.DicePool)
}

func (c *Crew) AvailableObjectives() []*Objective {
	return c.Objectives
}

func (c *Crew) Copy() Crew {
	objectives := make([]*Objective, len(c.Objectives))
	for i, objective := range c.Objectives {
		objectives[i] = new(*objective)
	}
	return Crew{
		Faction:    c.Faction,
		ID:         c.ID,
		IsStarter:  c.IsStarter,
		Name:       c.Name,
		Objectives: objectives,
		Modifier:   c.Modifier,
	}
}

func (c *Crew) CanCommit(pool dice.DicePool, playerName string) bool {
	for _, objective := range c.AvailableObjectives() {
		for _, die := range pool.Dice {
			if objective.IsType(die.ID) && objective.CanCommitBy(playerName) && !objective.IsCompleted() {
				return true
			}
		}
	}
	return false
}

func (c *Crew) IsComplete() bool {
	for _, objective := range c.Objectives {
		if !objective.IsCompleted() {
			return false
		}
	}
	return true
}

var allCrew = []Crew{
	{
		Name:    "Aponi",
		ID:      "aponi",
		Faction: factions.Blue,
		Objectives: []*Objective{
			{Type: new(dice.DieReactor), Amount: 4, Hazard: true},
			{Type: new(dice.DieThruster), Amount: 3},
			{Type: new(dice.DieShield), Amount: 3, Hazard: true},
			{Type: new(dice.DieWild), Amount: 2},
		},
	},
	{
		Name:    "Vila",
		ID:      "vila",
		Faction: factions.Blue,
		Objectives: []*Objective{
			{Type: new(dice.DieReactor), Amount: 4},
			{Type: new(dice.DieReactor), Amount: 3},
			{Type: new(dice.DieReactor), Amount: 2},
			{Type: new(dice.DieReactor), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Salatar",
		ID:      "salatar",
		Faction: factions.Blue,
		Objectives: []*Objective{
			{Type: new(dice.DieReactor), Amount: 4},
			{Type: new(dice.DieThruster), Amount: 3, Hazard: true},
			{Type: new(dice.DieShield), Amount: 3},
			{Type: new(dice.DieDamage), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Ada",
		ID:      "ada",
		Faction: factions.Blue,
		Objectives: []*Objective{
			{Type: new(dice.DieReactor), Amount: 2},
			{Type: new(dice.DieThruster), Amount: 2},
			{Type: new(dice.DieShield), Amount: 1, Hazard: true},
		},
		IsStarter: true,
		Modifier: func(pool *dice.DicePool) {
			newPool := make([]*dice.Die, 0, len(pool.Dice))
			for _, die := range pool.Dice {
				if die.ID == dice.DieExtra.ID {
					replace := &dice.Die{
						Symbol: die.Symbol,
						Color:  die.Color,
						ID:     die.ID,
						Value:  2,
						Mimics: &dice.DieReactor,
					}
					newPool = append(newPool, replace)
					continue
				}
				newPool = append(newPool, die)
			}
			pool.Dice = newPool
		},
	},
	{
		Name:    "Lee",
		ID:      "lee",
		Faction: factions.Blue,
		Objectives: []*Objective{
			{Type: new(dice.DieReactor), Amount: 2},
			{Type: new(dice.DieThruster), Amount: 2, Hazard: true},
			{Type: new(dice.DieShield), Amount: 2},
			{Type: new(dice.DieDamage), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Lila",
		ID:      "lila",
		Faction: factions.Blue,
		Objectives: []*Objective{
			{Type: new(dice.DieReactor), Amount: 3, Hazard: true},
			{Type: new(dice.DieThruster), Amount: 3},
			{Type: new(dice.DieShield), Amount: 2},
			{Type: new(dice.DieDamage), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "[REDACTED]",
		ID:      "redacted",
		Faction: factions.Green,
		Objectives: []*Objective{
			{Type: new(dice.DieShield), Amount: 4, Hazard: true},
			{Type: new(dice.DieDamage), Amount: 3},
			{Type: new(dice.DieThruster), Amount: 3},
			{Type: new(dice.DieWild), Amount: 2},
		},
	},
	{
		Name:    "Imdar",
		ID:      "imdar",
		Faction: factions.Green,
		Objectives: []*Objective{
			{Type: new(dice.DieShield), Amount: 4},
			{Type: new(dice.DieShield), Amount: 3},
			{Type: new(dice.DieShield), Amount: 2},
			{Type: new(dice.DieShield), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Namari",
		ID:      "namari",
		Faction: factions.Green,
		Objectives: []*Objective{
			{Type: new(dice.DieShield), Amount: 4},
			{Type: new(dice.DieDamage), Amount: 3, Hazard: true},
			{Type: new(dice.DieThruster), Amount: 3},
			{Type: new(dice.DieReactor), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Ryle",
		ID:      "ryle",
		Faction: factions.Green,
		Objectives: []*Objective{
			{Type: new(dice.DieShield), Amount: 2},
			{Type: new(dice.DieDamage), Amount: 2},
			{Type: new(dice.DieThruster), Amount: 1},
		},
		IsStarter: true,
		Modifier: func(pool *dice.DicePool) {
			newPool := make([]*dice.Die, 0, len(pool.Dice))
			for _, die := range pool.Dice {
				if die.ID == dice.DieExtra.ID {
					replace := &dice.Die{
						Symbol: die.Symbol,
						Color:  die.Color,
						ID:     die.ID,
						Value:  2,
						Mimics: &dice.DieShield,
					}
					newPool = append(newPool, replace)
					continue
				}
				newPool = append(newPool, die)
			}
			pool.Dice = newPool
		},
	},
	{
		Name:    "Bill",
		ID:      "bill",
		Faction: factions.Green,
		Objectives: []*Objective{
			{Type: new(dice.DieShield), Amount: 2, Hazard: true},
			{Type: new(dice.DieDamage), Amount: 2},
			{Type: new(dice.DieThruster), Amount: 2},
			{Type: new(dice.DieReactor), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "AT-OK",
		ID:      "at-ok",
		Faction: factions.Green,
		Objectives: []*Objective{
			{Type: new(dice.DieShield), Amount: 3, Hazard: true},
			{Type: new(dice.DieDamage), Amount: 3},
			{Type: new(dice.DieThruster), Amount: 2},
			{Type: new(dice.DieReactor), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Dr.Umbrage",
		ID:      "drumbrage",
		Faction: factions.Orange,
		Objectives: []*Objective{
			{Type: new(dice.DieDamage), Amount: 4, Hazard: true},
			{Type: new(dice.DieShield), Amount: 3},
			{Type: new(dice.DieReactor), Amount: 3},
			{Type: new(dice.DieWild), Amount: 2},
		},
	},
	{
		Name:    "Saghari",
		ID:      "saghari",
		Faction: factions.Orange,
		Objectives: []*Objective{
			{Type: new(dice.DieDamage), Amount: 4},
			{Type: new(dice.DieDamage), Amount: 3},
			{Type: new(dice.DieDamage), Amount: 2},
			{Type: new(dice.DieDamage), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Kary",
		ID:      "kary",
		Faction: factions.Orange,
		Objectives: []*Objective{
			{Type: new(dice.DieDamage), Amount: 4},
			{Type: new(dice.DieShield), Amount: 3, Hazard: true},
			{Type: new(dice.DieReactor), Amount: 3},
			{Type: new(dice.DieThruster), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Dana",
		ID:      "dana",
		Faction: factions.Orange,
		Objectives: []*Objective{
			{Type: new(dice.DieDamage), Amount: 3},
			{Type: new(dice.DieShield), Amount: 3},
			{Type: new(dice.DieReactor), Amount: 1, Hazard: true},
		},
		IsStarter: true,
		Modifier: func(pool *dice.DicePool) {
			newPool := make([]*dice.Die, 0, len(pool.Dice))
			for _, die := range pool.Dice {
				if die.ID == dice.DieExtra.ID {
					replace := &dice.Die{
						Symbol: die.Symbol,
						Color:  die.Color,
						ID:     die.ID,
						Value:  2,
						Mimics: &dice.DieDamage,
					}
					newPool = append(newPool, replace)
					continue
				}
				newPool = append(newPool, die)
			}
			pool.Dice = newPool
		},
	},
	{
		Name:    "Tantin",
		ID:      "tantin",
		Faction: factions.Orange,
		Objectives: []*Objective{
			{Type: new(dice.DieDamage), Amount: 2},
			{Type: new(dice.DieShield), Amount: 2, Hazard: true},
			{Type: new(dice.DieReactor), Amount: 2},
			{Type: new(dice.DieThruster), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Ryan",
		ID:      "ryan",
		Faction: factions.Orange,
		Objectives: []*Objective{
			{Type: new(dice.DieDamage), Amount: 3, Hazard: true},
			{Type: new(dice.DieShield), Amount: 3},
			{Type: new(dice.DieReactor), Amount: 2},
			{Type: new(dice.DieThruster), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Moro",
		ID:      "moro",
		Faction: factions.Purple,
		Objectives: []*Objective{
			{Type: new(dice.DieReactor), Amount: 4},
			{Type: new(dice.DieDamage), Amount: 3, Hazard: true},
			{Type: new(dice.DieShield), Amount: 3},
			{Type: new(dice.DieThruster), Amount: 2, Hazard: true},
		},
	},
	{
		Name:    "Vanta",
		ID:      "vanta",
		Faction: factions.Purple,
		Objectives: []*Objective{
			{Type: new(dice.DieWild), Amount: 3},
			{Type: new(dice.DieWild), Amount: 2},
			{Type: new(dice.DieWild), Amount: 1},
			{Type: new(dice.DieDamage), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Meg",
		ID:      "meg",
		Faction: factions.Purple,
		Objectives: []*Objective{
			{Type: new(dice.DieThruster), Amount: 4, Hazard: true},
			{Type: new(dice.DieDamage), Amount: 3},
			{Type: new(dice.DieShield), Amount: 3},
			{Type: new(dice.DieReactor), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Sella",
		ID:      "sella",
		Faction: factions.Purple,
		Objectives: []*Objective{
			{Type: new(dice.DieThruster), Amount: 2},
			{Type: new(dice.DieReactor), Amount: 2},
			{Type: new(dice.DieShield), Amount: 1},
		},
		IsStarter: true,
		Modifier: func(pool *dice.DicePool) {
			newPool := make([]*dice.Die, 0, len(pool.Dice))
			if pool.NumberOf(dice.DieExtra.ID) == 1 {
				for _, die := range pool.Dice {
					if die.ID == dice.DieExtra.ID {
						replace := &dice.Die{
							Symbol: die.Symbol,
							Color:  die.Color,
							ID:     die.ID,
							Value:  1,
							Mimics: &dice.DieWild,
						}
						newPool = append(newPool, replace)
						continue
					}
					newPool = append(newPool, die)
				}
				pool.Dice = newPool
			}
		},
	},
	{
		Name:    "FT-1000",
		ID:      "ft1000",
		Faction: factions.Purple,
		Objectives: []*Objective{
			{Type: new(dice.DieShield), Amount: 3},
			{Type: new(dice.DieThruster), Amount: 2, Hazard: true},
			{Type: new(dice.DieDamage), Amount: 2, Hazard: true},
			{Type: new(dice.DieReactor), Amount: 2},
		},
	},
	{
		Name:    "Avari",
		ID:      "avari",
		Faction: factions.Purple,
		Objectives: []*Objective{
			{Type: new(dice.DieDamage), Amount: 3, Hazard: true},
			{Type: new(dice.DieReactor), Amount: 3},
			{Type: new(dice.DieShield), Amount: 2},
			{Type: new(dice.DieThruster), Amount: 2, Hazard: true},
		},
	},
	{
		Name:    "Sol",
		ID:      "sol",
		Faction: factions.Yellow,
		Objectives: []*Objective{
			{Type: new(dice.DieThruster), Amount: 4, Hazard: true},
			{Type: new(dice.DieReactor), Amount: 3},
			{Type: new(dice.DieDamage), Amount: 3, Hazard: true},
			{Type: new(dice.DieWild), Amount: 2},
		},
	},
	{
		Name:    "B3-AR",
		ID:      "b3ar",
		Faction: factions.Yellow,
		Objectives: []*Objective{
			{Type: new(dice.DieThruster), Amount: 4},
			{Type: new(dice.DieThruster), Amount: 3},
			{Type: new(dice.DieThruster), Amount: 2},
			{Type: new(dice.DieThruster), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Kal",
		ID:      "kal",
		Faction: factions.Yellow,
		Objectives: []*Objective{
			{Type: new(dice.DieThruster), Amount: 4},
			{Type: new(dice.DieReactor), Amount: 3, Hazard: true},
			{Type: new(dice.DieDamage), Amount: 3},
			{Type: new(dice.DieShield), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Nella",
		ID:      "nella",
		Faction: factions.Yellow,
		Objectives: []*Objective{
			{Type: new(dice.DieThruster), Amount: 2},
			{Type: new(dice.DieReactor), Amount: 2},
			{Type: new(dice.DieDamage), Amount: 1, Hazard: true},
		},
		IsStarter: true,
		Modifier: func(pool *dice.DicePool) {
			newPool := make([]*dice.Die, 0, len(pool.Dice))
			for _, die := range pool.Dice {
				if die.ID == dice.DieExtra.ID {
					replace := &dice.Die{
						Symbol: die.Symbol,
						Color:  die.Color,
						ID:     die.ID,
						Value:  2,
						Mimics: &dice.DieThruster,
					}
					newPool = append(newPool, replace)
					continue
				}
				newPool = append(newPool, die)
			}
			pool.Dice = newPool
		},
	},
	{
		Name:    "Zek",
		ID:      "zek",
		Faction: factions.Yellow,
		Objectives: []*Objective{
			{Type: new(dice.DieThruster), Amount: 2},
			{Type: new(dice.DieReactor), Amount: 2, Hazard: true},
			{Type: new(dice.DieDamage), Amount: 2},
			{Type: new(dice.DieShield), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Myla",
		ID:      "myla",
		Faction: factions.Yellow,
		Objectives: []*Objective{
			{Type: new(dice.DieThruster), Amount: 3, Hazard: true},
			{Type: new(dice.DieReactor), Amount: 3},
			{Type: new(dice.DieDamage), Amount: 2},
			{Type: new(dice.DieShield), Amount: 1, Hazard: true},
		},
	},
}
