package dive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDive(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
		flag     bool
	}{
		{
			name: "part one",
			input: []string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2",
			},
			expected: 150,
			flag:     false,
		}, {
			name: "part two",
			input: []string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2",
			},
			expected: 900,
			flag:     true,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := Dive(tt.input, tt.flag)
			assert.Equal(t, tt.expected, res, "returned value should match expected	")
		})
	}
}
