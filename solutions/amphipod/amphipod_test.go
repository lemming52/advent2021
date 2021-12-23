package amphipod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinimum(t *testing.T) {
	tests := []struct {
		name          string
		rooms         []string
		lower         string
		expected      int
		expectedExtra int
	}{
		{
			name:     "base",
			rooms:    []string{"###B#C#B#D###", "  #A#D#C#A#  "},
			expected: 12521,
		}, {
			name:     "complex",
			rooms:    []string{"###B#C#B#D###", "  #D#C#B#A#  ", "  #D#B#A#C#  ", "  #A#D#C#A#  "},
			expected: 44169,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			expected := FindMinimum(tt.rooms)
			assert.Equal(t, tt.expected, expected, "returned value should match expected")
		})
	}
}
