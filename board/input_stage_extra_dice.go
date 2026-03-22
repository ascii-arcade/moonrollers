package board

import (
	"fmt"
	"strings"

	"github.com/ascii-arcade/moonrollers/keys"
)

type inputExtraDiceComponent struct {
	model *Model
}

func newInputExtraDiceComponent(model *Model) inputExtraDiceComponent {
	return inputExtraDiceComponent{
		model: model,
	}
}

func (c inputExtraDiceComponent) render() string {
	var output strings.Builder
	output.WriteString("How many extra die would you like to add?\n\n")
	fmt.Fprintf(&output, "%s to add from Supply\n", keys.GameChooseExtraDice.String(c.model.style))
	fmt.Fprintf(&output, "%s to remove extra from Pool\n", keys.GameUncommitDie.String(c.model.style))
	fmt.Fprintf(&output, "%s to continue", keys.GameChooseConfirm.String(c.model.style))
	return inputComponentStyle(true).Render(output.String())
}
