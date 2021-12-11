package dumbooctopus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModelOctopus(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		flashes int
		synced  int
	}{
		{
			name: "base",
			input: []string{
				"5483143223",
				"2745854711",
				"5264556173",
				"6141336146",
				"6357385478",
				"4167524645",
				"2176841721",
				"6882881134",
				"4846848554",
				"5283751526",
			},
			flashes: 1656,
			synced:  195,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			flashes, synced := ModelOctopus(tt.input, 100)
			assert.Equal(t, tt.flashes, flashes, "returned value should match expected")
			assert.Equal(t, tt.synced, synced, "returned value should match expected")
		})
	}
}
