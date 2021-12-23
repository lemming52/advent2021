package amphipod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinimum(t *testing.T) {
	tests := []struct {
		name     string
		upper    string
		lower    string
		expected int
	}{
		{
			name:  "base",
			upper: "###B#C#B#D###",
			lower: "  #A#D#C#A#  ",

			expected: 12521,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			expected := FindMinimum(tt.upper, tt.lower)
			assert.Equal(t, tt.expected, expected, "returned value should match expected")
		})
	}
}
