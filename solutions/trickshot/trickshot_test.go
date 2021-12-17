package trickshot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrickshot(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		count    int
	}{
		{
			name:     "base",
			input:    "target area: x=20..30, y=-10..-5",
			expected: 45,
			count:    112,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res, count, err := Trickshot(tt.input)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res, "returned value should match expected")
			assert.Equal(t, tt.count, count, "returned value should match expected")

		})
	}
}
