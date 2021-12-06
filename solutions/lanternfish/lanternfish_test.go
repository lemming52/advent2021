package lanternfish

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleOfLife(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		cycleLength int
		expected    int
	}{
		{
			name:        "short",
			input:       "3,4,3,1,2",
			cycleLength: 18,
			expected:    26,
		}, {
			name:        "long",
			input:       "3,4,3,1,2",
			cycleLength: 80,
			expected:    5934,
		}, {
			name:        "so long and thanks for all the fish",
			input:       "3,4,3,1,2",
			cycleLength: 256,
			expected:    5934,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res, err := CircleOfLife(tt.input, tt.cycleLength)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res, "returned value should match expected	")
		})
	}
}
