package snailmaths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnailSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
		mag      int
	}{
		{
			name: "base",
			input: []string{
				"[1,2]",
				"[[3,4],5]",
			},
			expected: 143,
			mag:      305,
		}, {
			name: "reduction",
			input: []string{
				"[[[[4,3],4],4],[7,[[8,4],9]]]",
				"[1,1]",
			},
			expected: 1384,
			mag:      2778,
		}, {
			name: "alpha",
			input: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
			},
			expected: 445,
			mag:      100,
		}, {
			name: "beta",
			input: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
			},
			expected: 791,
			mag:      125,
		}, {
			name: "gamma",
			input: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
				"[6,6]",
			},
			expected: 1137,
			mag:      150,
		}, {
			name: "bigger",
			input: []string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
				"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
				"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
				"[7,[5,[[3,8],[1,4]]]]",
				"[[2,[2,2]],[8,[8,1]]]",
				"[2,9]",
				"[1,[[[9,3],9],[[9,0],[0,7]]]]",
				"[[[5,[7,4]],7],1]",
				"[[[[4,2],2],6],[8,7]]",
			},
			expected: 3488,
			mag:      3946,
		}, {
			name: "final",
			input: []string{
				"[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
				"[[[5,[2,8]],4],[5,[[9,9],0]]]",
				"[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
				"[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]",
				"[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]",
				"[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]",
				"[[[[5,4],[7,7]],8],[[8,3],8]]",
				"[[9,3],[[9,9],[6,[4,9]]]]",
				"[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]",
				"[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]",
			},
			expected: 4140,
			mag:      3993,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res, mag := SnailSum(tt.input)
			assert.Equal(t, tt.expected, res, "returned value should match expected")
			assert.Equal(t, tt.mag, mag, "returned value should match expected")

		})
	}
}

/*


 */
