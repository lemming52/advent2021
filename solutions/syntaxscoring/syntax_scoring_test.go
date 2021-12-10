package syntaxscoring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpecifyParseLines(t *testing.T) {
	tests := []struct {
		name               string
		input              []string
		expectedCorrupt    int
		expectedIncomplete int
	}{
		{
			name: "vertical",
			input: []string{
				"[({(<(())[]>[[{[]{<()<>>",
				"[(()[<>])]({[<{<<[]>>(",
				"{([(<{}[<>[]}>{[]{[(<()>",
				"(((({<>}<{<{<>}{[]{[]{}",
				"[[<[([]))<([[{}[[()]]]",
				"[{[{({}]{}}([{[{{{}}([]",
				"{<[[]]>}<{[{[{[]{()[[[]",
				"[<(<(<(<{}))><([]([]()",
				"<{([([[(<>()){}]>(<<{{",
				"<{([{{}}[<[[[<>{}]]]>[]]",
			},
			expectedCorrupt:    26397,
			expectedIncomplete: 288957,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			corrupt, incomplete := ParseLines(tt.input)
			assert.Equal(t, tt.expectedCorrupt, corrupt, "returned value should match expected	")
			assert.Equal(t, tt.expectedIncomplete, incomplete, "returned value should match expected	")

		})
	}
}
