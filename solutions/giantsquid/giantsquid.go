package giantsquid

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	noWin   = -1
	endGame = -1
)

type Entry struct {
	value  int
	marked bool
}

type Board struct {
	grid     [][]*Entry
	values   map[int][]int
	played   int
	number   int
	incoming chan int
	results  chan int
}

func newBoard(values []string, boardNumber int, numbersChan, resultsChan chan int) (*Board, error) {
	grid := make([][]*Entry, len(values))
	for i, v := range values {
		grid[i] = make([]*Entry, len(strings.Fields(v)))
	}
	b := &Board{grid: grid, values: map[int][]int{}, number: boardNumber, incoming: numbersChan, results: resultsChan}
	for y, row := range values {
		for x, val := range strings.Fields(row) {
			v, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			b.grid[y][x] = &Entry{value: v}
			b.values[v] = []int{y, x}
		}
	}
	return b, nil
}

func (b *Board) mark(value int) (int, int, bool) {
	coords, ok := b.values[value]
	if !ok {
		return 0, 0, false
	}
	b.grid[coords[0]][coords[1]].marked = true
	b.played++
	return coords[0], coords[1], true
}

func (b *Board) hasWon(y, x int) bool {
	if b.played < len(b.grid) {
		return false
	}
	if b.hasWonColumns(x) {
		return true
	}
	return b.hasWonRows(y)
}

func (b *Board) hasWonColumns(x int) bool {
	for i := range b.grid {
		if !b.grid[i][x].marked {
			return false
		}
	}
	return true
}

func (b *Board) hasWonRows(y int) bool {
	for i := range b.grid[0] {
		if !b.grid[y][i].marked {
			return false
		}
	}
	return true
}

func (b *Board) score(value int) int {
	total := 0
	for _, row := range b.grid {
		for _, entry := range row {
			if !entry.marked {
				total += entry.value
			}
		}
	}
	return total * value
}

func (b *Board) Play() {
	for {
		value := <-b.incoming
		if value == endGame {
			return
		}
		y, x, played := b.mark(value)
		if played && b.hasWon(y, x) {
			b.results <- b.number
			return
		} else {
			b.results <- noWin
		}
	}
}

type BingoGame struct {
	numbers     []int
	boards      []*Board
	boardCount  int
	numbersChan chan int
	resultsChan chan int
}

func newBingoGame(numbers string, boards [][]string) (*BingoGame, error) {
	values := strings.Split(numbers, ",")
	boardCount := len(boards)
	b := &BingoGame{
		numbers:     make([]int, len(values)),
		boards:      make([]*Board, boardCount),
		boardCount:  boardCount,
		numbersChan: make(chan int),
		resultsChan: make(chan int),
	}
	for i, v := range values {
		val, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		b.numbers[i] = val
	}
	for i, v := range boards {
		var err error
		b.boards[i], err = newBoard(v, i, b.numbersChan, b.resultsChan)
		if err != nil {
			return nil, err
		}
	}
	return b, nil
}

func (b *BingoGame) Play() int {
	for _, board := range b.boards {
		go board.Play()
	}
	winner, value := b.callNumbers()
	b.finishGame()
	return b.boards[winner].score(value)
}

func (b *BingoGame) callNumbers() (int, int) {
	for _, val := range b.numbers {
		for i := 0; i < b.boardCount; i++ {
			b.numbersChan <- val
		}
		for i := 0; i < b.boardCount; i++ {
			res := <-b.resultsChan
			if res != noWin {
				b.boardCount--
				return res, val
			}
		}
	}
	return noWin, 0
}

func (b *BingoGame) finishGame() {
	close(b.numbersChan)
}

func (b *BingoGame) PlayToLose() int {
	for _, board := range b.boards {
		go board.Play()
	}
	winner, value := b.callNumbersLosing()
	b.finishGame()
	return b.boards[winner].score(value)
}

func (b *BingoGame) callNumbersLosing() (int, int) {
	for _, val := range b.numbers {
		fmt.Println("board count", b.boardCount)
		for i := 0; i < b.boardCount; i++ {
			fmt.Print(".")
			b.numbersChan <- val
		}
		active := b.boardCount
		for i := 0; i < active; i++ {
			fmt.Print("-")
			res := <-b.resultsChan
			if res != noWin {
				fmt.Println("result", res, b.boardCount)
				if b.boardCount == 1 {
					return res, val
				}
				b.boardCount--
			}
		}
	}
	return noWin, 0
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // first line is numbers
	numbers := scanner.Text()
	scanner.Scan() // skip next empty line
	boards := [][]string{}
	board := []string{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			boards = append(boards, board)
			board = []string{}
			continue
		}
		board = append(board, text)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	bingo, err := newBingoGame(numbers, boards)
	if err != nil {
		log.Fatal(err)
	}
	losing, err := newBingoGame(numbers, boards)
	if err != nil {
		log.Fatal(err)
	}
	return bingo.Play(), losing.PlayToLose()
}
