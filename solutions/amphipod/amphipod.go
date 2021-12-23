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

func newEmptyBoard(size int) *Board {
	b := &Board{
		rooms:   make([][]rune, 4),
		hallway: make([]rune, 11),
	}
	for i := 0; i < 4; i++ {
		b.rooms[i] = make([]rune, size)
	}
	return b
}

func newBoard(rooms []string) *Board {
	b := &Board{
		hallway:  make([]rune, 11),
		rooms:    make([][]rune, 4),
		previous: []*Board{},
	}
	for i := range b.hallway {
		b.hallway[i] = 0
	}
	for i := range b.rooms {
		b.rooms[i] = make([]rune, len(rooms))
	}
	for i, s := range rooms {
		for j, r := range s[3:11] {
			if r != '#' {
				b.rooms[j/2][i] = r
			}
		}
	}
	b.generateName()
	return b
}

func (b *Board) IsOrganised() bool {
	for i, rm := range b.rooms {
		desired := rune(i) + runeOffset
		for _, r := range rm {
			if r != desired {
				return false
			}
		}
	}
	return true
}

func (b *Board) generateName() {
	b.name = ""
	for _, r := range b.hallway {
		if r >= runeOffset && r <= runeOffset+4 {
			b.name += string(r)
		} else {
			b.name += " "
		}
	}
	for _, rm := range b.rooms {
		for _, r := range rm {
			if r >= runeOffset && r <= runeOffset+4 {
				b.name += string(r)
			} else {
				b.name += " "
			}
		}
	}
}

func (b *Board) getCopy() *Board {
	nb := newEmptyBoard(len(b.rooms[0]))
	copy(nb.hallway, b.hallway)
	for i, r := range b.rooms {
		copy(nb.rooms[i], r)
	}
	nb.cost = b.cost
	nb.previous = append(b.previous, b)
	return nb
}

func (b *Board) print() {
	fmt.Println("Cost: ", b.cost, b.name)
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
	fmt.Println(s)
	roomStrings := []string{"###"}
	for _ = range b.rooms[0][1:] {
		roomStrings = append(roomStrings, "  #")
	}
	for _, rm := range b.rooms {
		for j, r := range rm {
			if r >= runeOffset && r < runeOffset+4 {
				roomStrings[j] += fmt.Sprintf("%s#", string(r))
			} else {
				roomStrings[j] += " #"
			}
		}
	}
	for i, s := range roomStrings {
		if i == 0 {
			fmt.Println(s + "##")
		} else {
			fmt.Println(s + "  ")
		}
	}
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
	for i, rm := range b.rooms {
		for j, r := range rm {
			if r != 0 {
				newMoves := b.findRoomMoves(i, j, r)
				if newMoves != nil {
					moves = append(moves, newMoves...)
				}
				moves = append(moves)
				break
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
	roomIndex := -1
	for j := len(b.rooms[destination]) - 1; j >= 0; j-- {
		if b.rooms[destination][j] == 0 {
			roomIndex = j
			distance += j + 1
			break
		} else if b.rooms[destination][j] != r {
			return nil
		}
	}
	if roomIndex == -1 {
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

func FindMinimum(rooms []string) int {
	b := newBoard(rooms)
	b.print()
	pq := PriorityQueue{b}
	boards := map[string]*Board{b.name: b}
	heap.Init(&pq)

	cost := 0
	for pq.Len() != 0 {
		current := heap.Pop(&pq).(*Board)
		if current.IsOrganised() {
			cost = current.cost
			for _, v := range current.previous {
				v.print()
			}
			break
		}
		current.visited = true
		moves := current.allowedMoves()
		for _, m := range moves {
			if current.cost == 2691 {
				current.print()
				m.print()
			}
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
	return FindMinimum([]string{lines[2], lines[3]}), FindMinimum([]string{lines[2], "  #D#C#B#A#  ", "  #D#B#A#C#  ", lines[3]})
}
