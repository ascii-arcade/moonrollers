package board

type Crew struct {
	Faction    Faction
	ID         string
	IsStarter  bool
	Name       string
	Objectives []*Objective
	Modifier   func(game *Game)
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

func (c *Crew) CanCommit(pool DicePool, playerName string) bool {
	for _, objective := range c.AvailableObjectives() {
		for _, die := range pool.Dice {
			if (objective.IsType(die.ID) || die.ID == DieWild.ID || objective.IsType(die.Mimics)) && objective.CanCommitBy(playerName) && !objective.IsCompleted() {
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
		Name:    "Ada",
		ID:      "ada",
		Faction: Blue,
		Objectives: []*Objective{
			{Type: new(DieReactor), Amount: 2},
			{Type: new(DieThruster), Amount: 2},
			{Type: new(DieShield), Amount: 1, Hazard: true},
		},
		IsStarter: true,
		Modifier: func(game *Game) {
			for _, die := range game.RollingPool.Dice {
				if die.ID == DieExtra.ID {
					die.Mimics = DieReactor.ID
					die.Value = 2
				}
			}
		},
	},
	{
		Name:    "Aponi",
		ID:      "aponi",
		Faction: Blue,
		Objectives: []*Objective{
			{Type: new(DieReactor), Amount: 4, Hazard: true},
			{Type: new(DieThruster), Amount: 3},
			{Type: new(DieShield), Amount: 3, Hazard: true},
			{Type: new(DieWild), Amount: 2},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.Length() <= 3 {
				for _, die := range game.RollingPool.Dice {
					if die.ID == DieReactor.ID || die.ID == DieWild.ID {
						die.Mimics = DieReactor.ID
						die.Value = 2
					}
				}
			}
		},
	},
	{
		Name:    "AT-OK",
		ID:      "at-ok",
		Faction: Green,
		Objectives: []*Objective{
			{Type: new(DieShield), Amount: 3, Hazard: true},
			{Type: new(DieDamage), Amount: 3},
			{Type: new(DieThruster), Amount: 2},
			{Type: new(DieReactor), Amount: 1, Hazard: true},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.NumberOf(DieShield.ID) >= 2 {
				game.PreventBust = true
			}
		},
	},
	{
		Name:    "Avari",
		ID:      "avari",
		Faction: Purple,
		Objectives: []*Objective{
			{Type: new(DieDamage), Amount: 3, Hazard: true},
			{Type: new(DieReactor), Amount: 3},
			{Type: new(DieShield), Amount: 2},
			{Type: new(DieThruster), Amount: 2, Hazard: true},
		},
		Modifier: func(game *Game) {
			for _, die := range game.SupplyPool.Dice {
				if die.ID != DieWild.ID {
					return
				}
			}

			for range 3 {
				if game.SupplyPool.Length() == 0 {
					return
				}
				game.RollingPool.AddUnrolled()
				game.SupplyPool.RemoveUnrolled()
			}
		},
	},
	{
		Name:    "B3-AR",
		ID:      "b3ar",
		Faction: Yellow,
		Objectives: []*Objective{
			{Type: new(DieThruster), Amount: 4},
			{Type: new(DieThruster), Amount: 3},
			{Type: new(DieThruster), Amount: 2},
			{Type: new(DieThruster), Amount: 1, Hazard: true},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.NumberOf(DieThruster.ID) >= 3 {
				game.InputObjective.CompletedAmount = game.InputObjective.Amount
				game.InputState = InputStateRoll
			}
		},
	},
	{
		Name:    "Bill",
		ID:      "bill",
		Faction: Green,
		Objectives: []*Objective{
			{Type: new(DieShield), Amount: 2, Hazard: true},
			{Type: new(DieDamage), Amount: 2},
			{Type: new(DieThruster), Amount: 2},
			{Type: new(DieReactor), Amount: 1, Hazard: true},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.NumberOf(DieShield.ID) == 1 {
				for _, die := range game.RollingPool.Dice {
					if die.ID == DieShield.ID {
						die.Mimics = DieWild.ID
					}
				}
			}
		},
	},
	{
		Name:    "Dana",
		ID:      "dana",
		Faction: Orange,
		Objectives: []*Objective{
			{Type: new(DieDamage), Amount: 3},
			{Type: new(DieShield), Amount: 3},
			{Type: new(DieReactor), Amount: 1, Hazard: true},
		},
		IsStarter: true,
		Modifier: func(game *Game) {
			for _, die := range game.RollingPool.Dice {
				if die.ID == DieExtra.ID {
					die.Mimics = DieDamage.ID
					die.Value = 2
				}
			}
		},
	},
	{
		Name:    "Dr.Umbrage",
		ID:      "drumbrage",
		Faction: Orange,
		Objectives: []*Objective{
			{Type: new(DieDamage), Amount: 4, Hazard: true},
			{Type: new(DieShield), Amount: 3},
			{Type: new(DieReactor), Amount: 3},
			{Type: new(DieWild), Amount: 2},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.Length() <= 3 {
				for _, die := range game.RollingPool.Dice {
					if die.ID == DieDamage.ID || die.ID == DieWild.ID {
						die.Mimics = DieDamage.ID
						die.Value = 2
					}
				}
			}
		},
	},
	{
		Name:    "FT-1000",
		ID:      "ft1000",
		Faction: Purple,
		Objectives: []*Objective{
			{Type: new(DieShield), Amount: 3},
			{Type: new(DieThruster), Amount: 2, Hazard: true},
			{Type: new(DieDamage), Amount: 2, Hazard: true},
			{Type: new(DieReactor), Amount: 2},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.NumberOf(DieWild.ID) == 1 {
				for _, die := range game.RollingPool.Dice {
					if die.ID == DieWild.ID {
						die.Mimics = DieExtra.ID
					}
				}
			}
		},
	},
	{
		Name:    "Imdar",
		ID:      "imdar",
		Faction: Green,
		Objectives: []*Objective{
			{Type: new(DieShield), Amount: 4},
			{Type: new(DieShield), Amount: 3},
			{Type: new(DieShield), Amount: 2},
			{Type: new(DieShield), Amount: 1, Hazard: true},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.NumberOf(DieShield.ID) == 0 {
				game.InputState = InputStateOptionalHazard
			}
		},
	},
	{
		Name:    "Kal",
		ID:      "kal",
		Faction: Yellow,
		Objectives: []*Objective{
			{Type: new(DieThruster), Amount: 4},
			{Type: new(DieReactor), Amount: 3, Hazard: true},
			{Type: new(DieDamage), Amount: 3},
			{Type: new(DieShield), Amount: 1, Hazard: true},
		},
		Modifier: func(game *Game) {
			if game.RollCount == 0 {
				game.RollingPool.AddUnrolled()
				game.SupplyPool.RemoveUnrolled()
			}
		},
	},
	{
		Name:    "Kary",
		ID:      "kary",
		Faction: Orange,
		Objectives: []*Objective{
			{Type: new(DieDamage), Amount: 4},
			{Type: new(DieShield), Amount: 3, Hazard: true},
			{Type: new(DieReactor), Amount: 3},
			{Type: new(DieThruster), Amount: 1, Hazard: true},
		},
		Modifier: func(game *Game) {
			if game.RollCount == 1 {
				for _, die := range game.RollingPool.Dice {
					if die.ID == DieDamage.ID {
						die.Mimics = DieExtra.ID
					}
				}
			}
		},
	},
	{
		Name:    "Lee",
		ID:      "lee",
		Faction: Blue,
		Objectives: []*Objective{
			{Type: new(DieReactor), Amount: 2},
			{Type: new(DieThruster), Amount: 2, Hazard: true},
			{Type: new(DieShield), Amount: 2},
			{Type: new(DieDamage), Amount: 1, Hazard: true},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.NumberOf(DieReactor.ID) == 1 {
				for _, die := range game.RollingPool.Dice {
					if die.ID == DieReactor.ID {
						die.Mimics = DieWild.ID
					}
				}
			}
		},
	},
	{
		Name:    "Lila",
		ID:      "lila",
		Faction: Blue,
		Objectives: []*Objective{
			{Type: new(DieReactor), Amount: 3, Hazard: true},
			{Type: new(DieThruster), Amount: 3},
			{Type: new(DieShield), Amount: 2},
			{Type: new(DieDamage), Amount: 1, Hazard: true},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.NumberOf(DieReactor.ID) >= 2 {
				game.InputState = InputStateReroll
			}
		},
	},
	{
		Name:    "Meg",
		ID:      "meg",
		Faction: Purple,
		Objectives: []*Objective{
			{Type: new(DieThruster), Amount: 4, Hazard: true},
			{Type: new(DieDamage), Amount: 3},
			{Type: new(DieShield), Amount: 3},
			{Type: new(DieReactor), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Moro",
		ID:      "moro",
		Faction: Purple,
		Objectives: []*Objective{
			{Type: new(DieReactor), Amount: 4},
			{Type: new(DieDamage), Amount: 3, Hazard: true},
			{Type: new(DieShield), Amount: 3},
			{Type: new(DieThruster), Amount: 2, Hazard: true},
		},
	},
	{
		Name:    "Myla",
		ID:      "myla",
		Faction: Yellow,
		Objectives: []*Objective{
			{Type: new(DieThruster), Amount: 3, Hazard: true},
			{Type: new(DieReactor), Amount: 3},
			{Type: new(DieDamage), Amount: 2},
			{Type: new(DieShield), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Namari",
		ID:      "namari",
		Faction: Green,
		Objectives: []*Objective{
			{Type: new(DieShield), Amount: 4},
			{Type: new(DieDamage), Amount: 3, Hazard: true},
			{Type: new(DieThruster), Amount: 3},
			{Type: new(DieReactor), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Nella",
		ID:      "nella",
		Faction: Yellow,
		Objectives: []*Objective{
			{Type: new(DieThruster), Amount: 2},
			{Type: new(DieReactor), Amount: 2},
			{Type: new(DieDamage), Amount: 1, Hazard: true},
		},
		IsStarter: true,
		Modifier: func(game *Game) {
			for _, die := range game.RollingPool.Dice {
				if die.ID == DieExtra.ID {
					die.Mimics = DieThruster.ID
					die.Value = 2
				}
			}
		},
	},
	{
		Name:    "[REDACTED]",
		ID:      "redacted",
		Faction: Green,
		Objectives: []*Objective{
			{Type: new(DieShield), Amount: 4, Hazard: true},
			{Type: new(DieDamage), Amount: 3},
			{Type: new(DieThruster), Amount: 3},
			{Type: new(DieWild), Amount: 2},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.Length() <= 3 {
				for _, die := range game.RollingPool.Dice {
					if die.ID == DieShield.ID || die.ID == DieWild.ID {
						die.Mimics = DieShield.ID
						die.Value = 2
					}
				}
			}
		},
	},
	{
		Name:    "Ryan",
		ID:      "ryan",
		Faction: Orange,
		Objectives: []*Objective{
			{Type: new(DieDamage), Amount: 3, Hazard: true},
			{Type: new(DieShield), Amount: 3},
			{Type: new(DieReactor), Amount: 2},
			{Type: new(DieThruster), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Ryle",
		ID:      "ryle",
		Faction: Green,
		Objectives: []*Objective{
			{Type: new(DieShield), Amount: 2},
			{Type: new(DieDamage), Amount: 2},
			{Type: new(DieThruster), Amount: 1},
		},
		IsStarter: true,
		Modifier: func(game *Game) {
			for _, die := range game.RollingPool.Dice {
				if die.ID == DieExtra.ID {
					die.Mimics = DieShield.ID
					die.Value = 2
				}
			}
		},
	},
	{
		Name:    "Saghari",
		ID:      "saghari",
		Faction: Orange,
		Objectives: []*Objective{
			{Type: new(DieDamage), Amount: 4},
			{Type: new(DieDamage), Amount: 3},
			{Type: new(DieDamage), Amount: 2},
			{Type: new(DieDamage), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Salatar",
		ID:      "salatar",
		Faction: Blue,
		Objectives: []*Objective{
			{Type: new(DieReactor), Amount: 4},
			{Type: new(DieThruster), Amount: 3, Hazard: true},
			{Type: new(DieShield), Amount: 3},
			{Type: new(DieDamage), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Sella",
		ID:      "sella",
		Faction: Purple,
		Objectives: []*Objective{
			{Type: new(DieThruster), Amount: 2},
			{Type: new(DieReactor), Amount: 2},
			{Type: new(DieShield), Amount: 1},
		},
		IsStarter: true,
		Modifier: func(game *Game) {
			if game.RollingPool.NumberOfExtra() == 1 {
				for _, die := range game.RollingPool.Dice {
					if die.ID == DieExtra.ID {
						die.Mimics = DieWild.ID
					}
				}
			}
		},
	},
	{
		Name:    "Sol",
		ID:      "sol",
		Faction: Yellow,
		Objectives: []*Objective{
			{Type: new(DieThruster), Amount: 4, Hazard: true},
			{Type: new(DieReactor), Amount: 3},
			{Type: new(DieDamage), Amount: 3, Hazard: true},
			{Type: new(DieWild), Amount: 2},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.Length() <= 3 {
				for _, die := range game.RollingPool.Dice {
					if die.ID == DieThruster.ID || die.ID == DieWild.ID {
						die.Mimics = DieThruster.ID
						die.Value = 2
					}
				}
			}
		},
	},
	{
		Name:    "Tantin",
		ID:      "tantin",
		Faction: Orange,
		Objectives: []*Objective{
			{Type: new(DieDamage), Amount: 2},
			{Type: new(DieShield), Amount: 2, Hazard: true},
			{Type: new(DieReactor), Amount: 2},
			{Type: new(DieThruster), Amount: 1, Hazard: true},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.NumberOf(DieDamage.ID) == 1 {
				for _, die := range game.RollingPool.Dice {
					if die.ID == DieDamage.ID {
						die.Mimics = DieWild.ID
					}
				}
			}
		},
	},
	{
		Name:    "Vanta",
		ID:      "vanta",
		Faction: Purple,
		Objectives: []*Objective{
			{Type: new(DieWild), Amount: 3},
			{Type: new(DieWild), Amount: 2},
			{Type: new(DieWild), Amount: 1},
			{Type: new(DieDamage), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Vila",
		ID:      "vila",
		Faction: Blue,
		Objectives: []*Objective{
			{Type: new(DieReactor), Amount: 4},
			{Type: new(DieReactor), Amount: 3},
			{Type: new(DieReactor), Amount: 2},
			{Type: new(DieReactor), Amount: 1, Hazard: true},
		},
	},
	{
		Name:    "Zek",
		ID:      "zek",
		Faction: Yellow,
		Objectives: []*Objective{
			{Type: new(DieThruster), Amount: 2},
			{Type: new(DieReactor), Amount: 2, Hazard: true},
			{Type: new(DieDamage), Amount: 2},
			{Type: new(DieShield), Amount: 1, Hazard: true},
		},
		Modifier: func(game *Game) {
			if game.RollingPool.NumberOf(DieThruster.ID) == 1 {
				for _, die := range game.RollingPool.Dice {
					if die.ID == DieThruster.ID {
						die.Mimics = DieWild.ID
					}
				}
			}
		},
	},
}
