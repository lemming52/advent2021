package amphipod

import (
	"advent/solutions/utils"
	"container/heap"
	"fmt"
	"log"
)

const runeOffset = 65

type Board struct {
	rooms    [][]rune
	hallway  []rune
	name     string
	cost     int
	index    int
	visited  bool
	previous []*Board
}

func newEmptyBoard() *Board {
	b := &Board{
		rooms:   make([][]rune, 4),
		hallway: make([]rune, 11),
	}
	for i := 0; i < 4; i++ {
		b.rooms[i] = make([]rune, 2)
	}
	return b
}

func newBoard(upper, lower string) *Board {
	b := &Board{
		hallway:  make([]rune, 11),
		rooms:    make([][]rune, 4),
		previous: []*Board{},
	}
	for i := range b.hallway {
		b.hallway[i] = 0
	}
	for i, r := range upper[3:] {
		if r != '#' {
			b.rooms[i/2] = []rune{r, 0}
		}
	}
	for i, r := range lower[3:10] {
		if r != '#' {
			b.rooms[i/2][1] = r
		}
	}
	b.generateName()
	return b
}

func (b *Board) IsOrganised() bool {
	for i, r := range b.rooms {
		if r[0] != rune(i)+runeOffset {
			return false
		}
		if r[1] != rune(i)+runeOffset {
			return false
		}
	}
	return true
}

func (b *Board) generateName() {
	b.name = string(b.hallway)
	for _, r := range b.rooms {
		b.name += string(r)
	}
}

func (b *Board) getCopy() *Board {
	nb := newEmptyBoard()
	copy(nb.hallway, b.hallway)
	for i, r := range b.rooms {
		copy(nb.rooms[i], r)
	}
	nb.cost = b.cost
	nb.previous = append(b.previous, b)
	return nb
}

func (b *Board) print() {
	fmt.Println("Cost: ", b.cost, len(b.rooms), len(b.hallway))
	fmt.Println("#############")
	s := "#"
	for _, r := range b.hallway {
		if r >= runeOffset && r <= runeOffset+4 {
			s += string(r)
		} else {
			s += " "
		}
	}
	s += "#"
	s1, s2 := "###", "  #"
	for _, r := range b.rooms {
		if r[0] >= runeOffset && r[0] < runeOffset+4 {
			s1 += fmt.Sprintf("%s#", string(r[0]))
		} else {
			s1 += " #"
		}
		if r[1] >= runeOffset && r[1] < runeOffset+4 {
			s2 += fmt.Sprintf("%s#", string(r[1]))
		} else {
			s2 += " #"
		}
	}
	s1 += "##"
	s2 += "  "
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println("  #########  ")
}

func (b *Board) allowedMoves() []*Board {
	moves := []*Board{}
	for i, r := range b.hallway {
		if r != 0 {
			m := b.findHallwayMove(i, r)
			if m != nil {
				moves = append(moves, m)

			}
		}
	}
	for i, r := range b.rooms {
		if r[0] != 0 {
			newMoves := b.findRoomMoves(i, 0, r[0])
			if newMoves != nil {
				moves = append(moves, newMoves...)
			}
			moves = append(moves)
		} else if r[1] != 0 {
			newMoves := b.findRoomMoves(i, 1, r[1])
			if newMoves != nil {
				moves = append(moves, newMoves...)
			}
		}
	}
	return moves
}

func (b *Board) findHallwayMove(i int, r rune) *Board {
	destination := int(r - runeOffset)
	distance := i - (destination+1)*2
	if distance < 0 {
		distance *= -1
	}
	roomIndex := 0
	if b.rooms[destination][1] == 0 {
		roomIndex = 1
		distance += 2
	} else if b.rooms[destination][0] == 0 {
		distance += 1
	} else {
		return nil
	}
	increment := 1
	if i > (destination+1)*2 {
		increment = -1
	}
	for j := i + increment; j != (destination+1)*2; j += increment {
		if b.hallway[j] != 0 {
			return nil
		}
	}
	nb := b.getCopy()
	nb.hallway[i] = 0
	nb.rooms[destination][roomIndex] = r
	for j := 0; j < destination; j++ {
		distance *= 10
	}
	nb.cost += distance
	nb.generateName()
	return nb
}

func (b *Board) findRoomMoves(i, roomIndex int, r rune) []*Board {
	destination := int(r - runeOffset)
	if i == destination {
		if roomIndex == 1 {
			return nil
		}
		if b.rooms[i][1] == r {
			return nil
		}
	}
	position := (i + 1) * 2
	moves := []*Board{}
	// Look left
	for j := position - 1; j >= 0; j-- {
		if b.hallway[j] != 0 {
			break
		}
		if j%2 == 0 && j != 0 {
			continue
		}
		moves = append(moves, b.newRoomMove(i, position, j, roomIndex, r))
	}
	for j := position + 1; j < 11; j++ {
		if b.hallway[j] != 0 {
			break
		}
		if j%2 == 0 && (j != 0 && j != 10) {
			continue
		}
		moves = append(moves, b.newRoomMove(i, position, j, roomIndex, r))
	}
	return moves
}

func (b *Board) newRoomMove(i, position, j, roomIndex int, r rune) *Board {
	nb := b.getCopy()
	nb.hallway[j] = r
	nb.rooms[i][roomIndex] = 0

	distance := j - position
	if distance < 0 {
		distance *= -1
	}
	distance += roomIndex + 1

	factor := int(r - runeOffset)
	for k := 0; k < factor; k++ {
		distance *= 10
	}
	nb.cost += distance
	nb.generateName()
	return nb
}

func FindMinimum(upper, lower string) int {
	b := newBoard(upper, lower)
	pq := PriorityQueue{b}
	boards := map[string]*Board{b.name: b}
	heap.Init(&pq)

	cost := 0
	for pq.Len() != 0 {
		current := heap.Pop(&pq).(*Board)
		//current.print()
		if current.IsOrganised() {
			cost = current.cost
			break
		}
		current.visited = true
		moves := current.allowedMoves()
		for _, m := range moves {
			if m.visited {
				continue
			}
			e, ok := boards[m.name]
			if !ok {
				boards[m.name] = m
				heap.Push(&pq, m)
			} else {
				if m.cost < e.cost {
					boards[m.name] = m
					heap.Push(&pq, m)
				}
			}
		}
	}
	return cost
}

func Challenge(path string) (int, int) {
	lines, err := utils.ReadStrings(path)
	if err != nil {
		log.Fatal(err)
	}
	a := FindMinimum(lines[2], lines[3])
	return a, 0
}
