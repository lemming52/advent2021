package treacheryofwhales

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrabSubmarines(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  int
		expected2 int
	}{
		{
			name:      "base",
			input:     "16,1,2,0,4,2,7,1,2,14",
			expected:  37,
			expected2: 168,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res1, res2, err := CrabSubmarines(tt.input)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, res1, "returned value should match expected")
			assert.Equal(t, tt.expected2, res2, "returned value should match expected")
		})
	}
}
