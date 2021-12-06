package giantsquid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewboard(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected *Board
	}{
		{
			name: "base",
			input: []string{
				"15 79 24",
				"25 41 63",
				"13 83 69",
			},
			expected: &Board{
				grid: [][]*Entry{
					[]*Entry{
						&Entry{value: 15}, &Entry{value: 79}, &Entry{value: 24},
					}, []*Entry{
						&Entry{value: 25}, &Entry{value: 41}, &Entry{value: 63},
					}, []*Entry{
						&Entry{value: 13}, &Entry{value: 83}, &Entry{value: 69},
					},
				},
				values: map[int][]int{
					15: []int{0, 0},
					79: []int{0, 1},
					24: []int{0, 2},
					25: []int{1, 0},
					41: []int{1, 1},
					63: []int{1, 2},
					13: []int{2, 0},
					83: []int{2, 1},
					69: []int{2, 2},
				},
				number: 1,
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res, err := newBoard(tt.input, 1, nil, nil)
			assert.Equal(t, tt.expected, res, "returned value should match expected")
			assert.Nil(t, err)
		})
	}
}

func TestMark(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		value    int
		expected *Board
	}{
		{
			name: "base",
			input: []string{
				"15 79 24",
				"25 41 63",
				"13 83 69",
			},
			value: 41,
			expected: &Board{
				grid: [][]*Entry{
					[]*Entry{
						&Entry{value: 15}, &Entry{value: 79}, &Entry{value: 24},
					}, []*Entry{
						&Entry{value: 25}, &Entry{value: 41, marked: true}, &Entry{value: 63},
					}, []*Entry{
						&Entry{value: 13}, &Entry{value: 83}, &Entry{value: 69},
					},
				},
				values: map[int][]int{
					15: []int{0, 0},
					79: []int{0, 1},
					24: []int{0, 2},
					25: []int{1, 0},
					41: []int{1, 1},
					63: []int{1, 2},
					13: []int{2, 0},
					83: []int{2, 1},
					69: []int{2, 2},
				},
				number: 1,
				played: 1,
			},
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res, err := newBoard(tt.input, 1, nil, nil)
			assert.Nil(t, err)
			res.mark(tt.value)
			assert.Equal(t, tt.expected, res, "returned value should match expected")
		})
	}
}

func TestHasWon(t *testing.T) {
	tests := []struct {
		name     string
		input    *Board
		expected bool
	}{
		{
			name: "played limit",
			input: &Board{
				grid: [][]*Entry{
					[]*Entry{
						&Entry{value: 15}, &Entry{value: 79, marked: true}, &Entry{value: 24},
					}, []*Entry{
						&Entry{value: 25}, &Entry{value: 41, marked: true}, &Entry{value: 63},
					}, []*Entry{
						&Entry{value: 13}, &Entry{value: 83, marked: true}, &Entry{value: 69},
					},
				},
				values: map[int][]int{
					15: []int{0, 0},
					79: []int{0, 1},
					24: []int{0, 2},
					25: []int{1, 0},
					41: []int{1, 1},
					63: []int{1, 2},
					13: []int{2, 0},
					83: []int{2, 1},
					69: []int{2, 2},
				},
				played: 1,
			},
			expected: false,
		}, {
			name: "column",
			input: &Board{
				grid: [][]*Entry{
					[]*Entry{
						&Entry{value: 15}, &Entry{value: 79, marked: true}, &Entry{value: 24},
					}, []*Entry{
						&Entry{value: 25}, &Entry{value: 41, marked: true}, &Entry{value: 63},
					}, []*Entry{
						&Entry{value: 13}, &Entry{value: 83, marked: true}, &Entry{value: 69},
					},
				},
				values: map[int][]int{
					15: []int{0, 0},
					79: []int{0, 1},
					24: []int{0, 2},
					25: []int{1, 0},
					41: []int{1, 1},
					63: []int{1, 2},
					13: []int{2, 0},
					83: []int{2, 1},
					69: []int{2, 2},
				},
				played: 3,
			},
			expected: true,
		}, {
			name: "row",
			input: &Board{
				grid: [][]*Entry{
					[]*Entry{
						&Entry{value: 15, marked: true}, &Entry{value: 79, marked: true}, &Entry{value: 24, marked: true},
					}, []*Entry{
						&Entry{value: 25}, &Entry{value: 41, marked: true}, &Entry{value: 63},
					}, []*Entry{
						&Entry{value: 13}, &Entry{value: 83, marked: false}, &Entry{value: 69},
					},
				},
				values: map[int][]int{
					15: []int{0, 0},
					79: []int{0, 1},
					24: []int{0, 2},
					25: []int{1, 0},
					41: []int{1, 1},
					63: []int{1, 2},
					13: []int{2, 0},
					83: []int{2, 1},
					69: []int{2, 2},
				},
				played: 3,
			},
			expected: true,
		}, {
			name: "not won",
			input: &Board{
				grid: [][]*Entry{
					[]*Entry{
						&Entry{value: 15, marked: false}, &Entry{value: 79, marked: true}, &Entry{value: 24, marked: true},
					}, []*Entry{
						&Entry{value: 25}, &Entry{value: 41, marked: false}, &Entry{value: 63},
					}, []*Entry{
						&Entry{value: 13}, &Entry{value: 83, marked: true}, &Entry{value: 69},
					},
				},
				values: map[int][]int{
					15: []int{0, 0},
					79: []int{0, 1},
					24: []int{0, 2},
					25: []int{1, 0},
					41: []int{1, 1},
					63: []int{1, 2},
					13: []int{2, 0},
					83: []int{2, 1},
					69: []int{2, 2},
				},
				played: 3,
			},
			expected: false,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			res := tt.input.hasWon(0, 1)
			assert.Equal(t, tt.expected, res, "returned value should match expected")
		})
	}
}

func TestPlay(t *testing.T) {
	tests := []struct {
		name     string
		numbers  string
		boards   [][]string
		expected int
	}{
		{
			name:    "base",
			numbers: "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
			boards: [][]string{
				[]string{
					"22 13 17 11  0",
					"8  2 23  4 24",
					"21  9 14 16  7",
					"6 10  3 18  5",
					"1 12 20 15 19",
				}, []string{
					"3 15  0  2 22",
					"9 18 13 17  5",
					"19  8  7 25 23",
					"20 11 10 24  4",
					"14 21 16 12  6",
				}, []string{
					"14 21 17 24  4",
					"10 16 15  9 19",
					"18  8 23 26 20",
					"22 11 13  6  5",
					" 2  0 12  3  7",
				},
			},
			expected: 4512,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			bingo, err := newBingoGame(tt.numbers, tt.boards)
			assert.Nil(t, err)
			res := bingo.Play()
			assert.Equal(t, tt.expected, res, "returned value should match expected")
		})
	}
}

func TestPlayToLose(t *testing.T) {
	tests := []struct {
		name     string
		numbers  string
		boards   [][]string
		expected int
	}{
		{
			name:    "base",
			numbers: "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
			boards: [][]string{
				[]string{
					"22 13 17 11  0",
					"8  2 23  4 24",
					"21  9 14 16  7",
					"6 10  3 18  5",
					"1 12 20 15 19",
				}, []string{
					"3 15  0  2 22",
					"9 18 13 17  5",
					"19  8  7 25 23",
					"20 11 10 24  4",
					"14 21 16 12  6",
				}, []string{
					"14 21 17 24  4",
					"10 16 15  9 19",
					"18  8 23 26 20",
					"22 11 13  6  5",
					" 2  0 12  3  7",
				},
			},
			expected: 1924,
		},
	}
	for _, test := range tests {
		tt := test
		t.Run(tt.name, func(t *testing.T) {
			bingo, err := newBingoGame(tt.numbers, tt.boards)
			assert.Nil(t, err)
			res := bingo.PlayToLose()
			assert.Equal(t, tt.expected, res, "returned value should match expected")
		})
	}
}
