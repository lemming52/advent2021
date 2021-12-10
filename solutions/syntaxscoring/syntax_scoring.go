package syntaxscoring

import (
	"bufio"
	"log"
	"os"
	"sort"
)

type ChunkParser struct {
	brackets         []rune
	corruptScores    map[rune]int
	corruptTotal     int
	incompleteScores map[rune]int
	incompleteTotals []int
	incompleteIndex  int
}

func newChunkParser(count int) *ChunkParser {
	return &ChunkParser{
		brackets: []rune{},
		corruptScores: map[rune]int{
			')': 3,
			']': 57,
			'}': 1197,
			'>': 25137,
		},
		incompleteScores: map[rune]int{
			')': 1,
			']': 2,
			'}': 3,
			'>': 4,
		},
		incompleteTotals: make([]int, count),
		incompleteIndex:  0,
	}
}

func (c *ChunkParser) resetParser() {
	c.brackets = []rune{}
}

func (c *ChunkParser) parseLine(l string) {
	c.brackets = make([]rune, len(l))
	index := 0
	for _, r := range l {
		switch r {
		case '{', '<', '[', '(':
			c.brackets[index] = r
			index++
		case '}', '>', ']', ')':
			if c.brackets[index-1] != complement(r) {
				// Corrupted
				c.corruptTotal += c.corruptScores[r]
				return
			}
			c.brackets[index-1] = 0
			index--
		}
	}
	c.scoreIncomplete(index)
}

func complement(r rune) rune {
	switch r {
	case '}':
		return '{'
	case '>':
		return '<'
	case ']':
		return '['
	case ')':
		return '('
	case '{':
		return '}'
	case '<':
		return '>'
	case '[':
		return ']'
	case '(':
		return ')'
	}
	return 0
}

func (c *ChunkParser) scoreIncomplete(index int) {
	score := 0
	for i := index - 1; i >= 0; i-- {
		score *= 5
		r := c.brackets[i]
		score += c.incompleteScores[complement(r)]
	}
	c.addIncompleteScore(score)
}

func (c *ChunkParser) addIncompleteScore(score int) {
	sortIndex := sort.Search(c.incompleteIndex, func(i int) bool { return c.incompleteTotals[i] >= score })
	copy(c.incompleteTotals[sortIndex+1:], c.incompleteTotals[sortIndex:])
	c.incompleteTotals[sortIndex] = score
	c.incompleteIndex++
}

func ParseLines(lines []string) (int, int) {
	parser := newChunkParser(len(lines))
	for _, l := range lines {
		parser.parseLine(l)
		parser.resetParser()
	}
	return parser.corruptTotal, parser.incompleteTotals[parser.incompleteIndex/2]
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ParseLines(lines)
}
