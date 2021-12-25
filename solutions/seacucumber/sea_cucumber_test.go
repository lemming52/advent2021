package seacucumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	tests := []struct {
		name     string
		cucumber []string
		expected int
	}{
		{
			name: "base",
			cucumber: []string{
				"v...>>.vv>",
				".vv>>.vv..",
				">>.>v>...v",
				">>v>>.>.v.",
				"v>v.vv.v..",
				">.>>..v...",
				".vv..>.>v.",
				"v.v..>>v.v",
				"....v..v.>",
			},
			expected: 58,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			expected := Migrate(tt.cucumber)
			assert.Equal(t, tt.expected, expected, "returned value should match expected")
		})
	}
}
