package sonarsweep

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSonarSweep(t *testing.T) {
	tests := []struct {
		name     string
		input    []int64
		window   int
		expected int
	}{
		{
			name:     "part one",
			input:    []int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			window:   1,
			expected: 7,
		}, {
			name:     "part two",
			input:    []int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			window:   3,
			expected: 5,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := SonarSweep(tt.input, tt.window)
			assert.Equal(t, tt.expected, res, "returned value should match expected	")
		})
	}
}
