package rules

import (
	"testing"

	"github.com/ascii-arcade/moonrollers/dice"
)

func TestIsUsable(t *testing.T) {
	type input struct {
		crew          []string
		rolledDie     *dice.Die
		objectiveType *dice.Die
	}
	type output struct {
		usable bool
		count  int
	}
	type scenario struct {
		name    string
		input   input
		want    output
		wantErr bool
	}
	scenarios := []scenario{
		{
			name:    "no crew",
			input:   input{crew: nil, rolledDie: &dice.DieDamage, objectiveType: &dice.DieDamage},
			want:    output{usable: false, count: 0},
			wantErr: true,
		},
		{
			name:    "no rolledDie",
			input:   input{crew: []string{"dana"}, rolledDie: nil, objectiveType: &dice.DieDamage},
			want:    output{usable: false, count: 0},
			wantErr: true,
		},
		{
			name:    "no objectiveType",
			input:   input{crew: []string{"dana"}, rolledDie: &dice.DieDamage, objectiveType: nil},
			want:    output{usable: false, count: 0},
			wantErr: true,
		},
		{
			name:    "empty crew, damage as damage",
			input:   input{crew: []string{}, rolledDie: &dice.DieDamage, objectiveType: &dice.DieDamage},
			want:    output{usable: true, count: 1},
			wantErr: false,
		},
		{
			name:    "empty crew, damage as shield",
			input:   input{crew: []string{}, rolledDie: &dice.DieDamage, objectiveType: &dice.DieShield},
			want:    output{usable: false, count: 0},
			wantErr: false,
		},
		{
			name:    "dana, extra as 2 damage",
			input:   input{crew: []string{"dana"}, rolledDie: &dice.DieExtra, objectiveType: &dice.DieDamage},
			want:    output{usable: true, count: 2},
			wantErr: false,
		},
		{
			name:    "dana, extra as extra",
			input:   input{crew: []string{"dana"}, rolledDie: &dice.DieExtra, objectiveType: &dice.DieExtra},
			want:    output{usable: true, count: 1},
			wantErr: false,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			isUsable, count, err := IsUsable(s.input.crew, s.input.rolledDie, s.input.objectiveType)

			if s.wantErr && err == nil {
				t.Errorf("expected an error, got none")
			} else if !s.wantErr && err != nil {
				t.Errorf("expected no error, got %v", err)
			} else {
				if isUsable != s.want.usable || count != s.want.count {
					t.Errorf("IsUsable() = (%v, %d), want (%v, %d)", isUsable, count, s.want.usable, s.want.count)
				}
			}
		})
	}
}
