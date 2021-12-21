package diracdie

import (
	"advent/solutions/utils"
	"log"
	"regexp"
	"strconv"
)

const startPattern = `Player \d starting position: (\d+)`

type Board struct {
	size      int
	positions []int
}

func newBoard(size int, positions []int) (*Board, []*Player) {
	b := &Board{
		size:      size,
		positions: make([]int, len(positions)),
	}
	players := make([]*Player, len(positions))
	for i, p := range positions {
		b.positions[i] = p
		players[i] = newPlayer(i)
	}
	return b, players
}

func (b *Board) move(player, move int) int {
	position := (b.positions[player] + move) % b.size
	b.positions[player] = position
	if position == 0 {
		return 10
	}
	return position
}

type Player struct {
	name  int
	score int
}

func newPlayer(name int) *Player {
	return &Player{
		name:  name,
		score: 0,
	}
}

type Die struct {
	rolls int
	count int
}

func (d *Die) roll() int {
	d.rolls += 3
	if d.count <= 97 {
		roll := (d.count + 2) * 3
		d.count += 3
		return roll
	}
	switch d.count {
	case 98:
		d.count = 1
		return 200
	case 99:
		d.count = 2
		return 103
	default:
		d.count = 3
		return 6
	}
}

func takeTurn(p *Player, d *Die, b *Board) bool {
	roll := d.roll()
	score := b.move(p.name, roll)
	p.score += score
	return p.score >= 1000
}

func PlayDice(starts []string) (int, int, error) {
	pattern, err := regexp.Compile(startPattern)
	if err != nil {
		return 0, 0, err
	}
	positions := make([]int, len(starts))
	for i, s := range starts {
		startPosition := pattern.FindStringSubmatch(s)
		val, err := strconv.Atoi(startPosition[1])
		if err != nil {
			return 0, 0, err
		}
		positions[i] = val
	}
	board, players := newBoard(10, positions)
	die := &Die{rolls: 0, count: 0}
	currentPlayer, playerCount := 0, len(players)
	victor := false
	for !victor {
		p := players[currentPlayer]
		victor = takeTurn(p, die, board)
		currentPlayer = (currentPlayer + 1) % playerCount
	}
	return players[currentPlayer].score * die.rolls, PlayQuantum(positions[0], positions[1]), nil
}

func Challenge(path string) (int, int) {
	lines, err := utils.ReadStrings(path)
	if err != nil {
		log.Fatal(err)
	}
	a, b, err := PlayDice(lines)
	if err != nil {
		log.Fatal(err)
	}
	return a, b
}
