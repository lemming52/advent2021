package diracdie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayDice(t *testing.T) {
	tests := []struct {
		name    string
		players []string
		score   int
		quantum int
	}{
		{
			name: "backtrack",
			players: []string{
				"Player 1 starting position: 4",
				"Player 2 starting position: 8",
			},
			score:   739785,
			quantum: 444356092776315,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			score, quantum, err := PlayDice(tt.players)
			assert.Nil(t, err)
			assert.Equal(t, tt.score, score, "returned value should match expected")
			assert.Equal(t, tt.quantum, quantum, "returned value should match expected")
		})
	}
}
