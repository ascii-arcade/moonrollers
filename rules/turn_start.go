package rules

import "slices"

type startTurn struct {
	RollingPoolSize int
}

func NewStartTurn(crewIDs []string) startTurn {
	rollingPoolSize := 5

	if slices.Contains(crewIDs, "kal") {
		rollingPoolSize = 6
	}

	return startTurn{
		RollingPoolSize: rollingPoolSize,
	}
}
