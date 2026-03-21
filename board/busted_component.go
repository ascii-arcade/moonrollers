package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
)

type bustedComponent struct {
	model *Model
}

func newBustedComponent(model *Model) bustedComponent {
	return bustedComponent{
		model: model,
	}
}

func (c bustedComponent) render() string {
	var output strings.Builder
	output.WriteString("Busted!\n\n")
	fmt.Fprintf(&output, "\n%s to continue", keys.GameEndTurn.String(c.model.style))
	return output.String()
}
