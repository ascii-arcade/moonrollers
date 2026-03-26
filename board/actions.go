package board

import (
	"errors"
	"slices"
)

func (g *Game) SetFaction(player *Player, faction *Faction) error {
	return g.withErrLock(func() error {
		player.Faction = faction
		return nil
	})
}

func (g *Game) Roll(isRolling bool) {
	g.withLock(func() {
		g.RollCount++
		g.RollingPool.Roll()
		if !isRolling {
			switch {
			case g.InputCrew != nil && !g.InputCrew.CanCommit(g.RollingPool, g.GetCurrentPlayer().Sess.User()),
				g.InputObjective != nil && g.RollingPool.ValueOf(g.InputObjective.Type.ID) == 0:
				if !g.PreventBust {
					g.InputState = Busted
					break
				}
				g.PreventBust = false
			case g.InputObjective != nil && g.InputObjective.CompletedAmount < g.InputObjective.Amount:
				g.InputState = InputStateCommitDice
			case g.InputCrew != nil && !g.InputCrew.IsComplete():
				g.InputState = InputStateChooseObjective
			default:
				g.InputState = InputStateChooseCrew
			}

			g.ApplyModifiers()

			slices.SortFunc(g.RollingPool.Dice, func(a, b *Die) int {
				if a.ID == DieExtra.ID && b.ID != DieExtra.ID {
					return -1
				}
				if b.ID == DieExtra.ID && a.ID != DieExtra.ID {
					return 1
				}
				return 0
			})
		}
	})
}

func (g *Game) ApplyModifiers() {
	for _, c := range g.GetCurrentPlayer().Crew {
		if c.Modifier != nil {
			c.Modifier(g)
		}
	}
}

func (g *Game) ChooseCrewMember(index int) {
	g.withLock(func() {
		if len(g.CrewForHire) <= index || !g.CrewForHire[index].CanCommit(g.RollingPool, g.GetCurrentPlayer().Name) {
			return
		}
		inputCrew := g.CrewForHire[index].Copy()
		g.InputCrew = &inputCrew
	})
}

func (g *Game) ConfirmCrewMember() {
	g.withLock(func() {
		if g.InputCrew == nil {
			return
		}
		g.InputState = InputStateChooseObjective
	})
}

func (g *Game) ChooseObjective(index int) {
	g.withLock(func() {
		switch {
		case index < 0,
			index >= len(g.InputCrew.Objectives),
			g.RollingPool.ValueOf(g.InputCrew.Objectives[index].Type.ID) == 0,
			g.InputCrew.Objectives[index].IsCompleted(),
			g.InputCrew.Objectives[index].StartedBy != "" && g.InputCrew.Objectives[index].StartedBy != g.GetCurrentPlayer().Name:
			return
		}

		g.InputObjective = g.InputCrew.Objectives[index]
	})
}

func (g *Game) ConfirmObjective() {
	g.withLock(func() {
		if g.InputObjective == nil {
			return
		}
		g.InputState = InputStateCommitDice
	})
}

func (g *Game) PreviousInputStage() {
	g.withLock(func() {
		switch g.InputState {
		case InputStateChooseObjective:
			g.InputObjective = nil
			g.InputState = InputStateChooseCrew
		case InputStateCommitDice:
			g.InputState = InputStateChooseObjective
		}
	})
}

func (g *Game) HireCrewMember(index int, player *Player) error {
	return g.withErrLock(func() error {
		if index < 0 || index >= len(g.CrewForHire) {
			return errors.New("invalid_crew_index")
		}

		crew := g.CrewForHire[index]
		if crew == nil {
			return errors.New("crew_not_found")
		}

		player.AddCrew(crew, true)
		g.CrewForHire = slices.Delete(g.CrewForHire, index, index+1)
		g.dealSingleCrew()
		return nil
	})
}

func (g *Game) dealSingleCrew() {
	if len(g.Deck) == 0 {
		return
	}
	g.CrewForHire = append(g.CrewForHire, g.Deck[0])
	g.Deck = slices.Delete(g.Deck, 0, 1)
}
