package chiton

import (
	"bufio"
	"container/heap"
	"log"
	"math"
	"os"
)

const runeOffset = 48

type Node struct {
	y, x, distance, value, index int
	visited                      bool
}

func newNode(y, x, value int) *Node {
	return &Node{
		y:        y,
		x:        x,
		distance: math.MaxInt,
		visited:  false,
		value:    value,
	}
}

type Cave struct {
	chiton   []string
	nodes    [][]*Node
	yMax     int
	xMax     int
	minimums [][]int // for the hack
}

func (c *Cave) dijkstra() {
	current := c.nodes[0][0]
	current.distance = 0

	pq := prioQueue{current}
	heap.Init(&pq)

	for pq.Len() != 0 {
		current := heap.Pop(&pq).(*Node)
		if current.x == c.xMax && current.y == c.yMax {
			break
		}
		current.visited = true
		neighbours := c.getCoords(current.y, current.x)
		for _, n := range neighbours {
			tentative := current.distance + n.value
			if tentative < n.distance {
				n.distance = tentative
				heap.Push(&pq, n)
			}
		}
	}
}

func (c *Cave) getCoords(y, x int) []*Node {
	coords := [][]int{
		{y - 1, x},
		{y, x - 1},
		{y + 1, x},
		{y, x + 1},
	}
	correct := []*Node{}
	for _, yx := range coords {
		if yx[0] >= 0 && yx[1] >= 0 && yx[0] <= c.yMax && yx[1] <= c.xMax && !c.nodes[yx[0]][yx[1]].visited {
			correct = append(correct, c.nodes[yx[0]][yx[1]])
		}
	}
	return correct
}

func val(r byte) int {
	return int(r) - runeOffset
}

func AvoidChiton(chiton []string) int {
	cave := &Cave{
		chiton: chiton,
		yMax:   len(chiton) - 1,
		xMax:   len(chiton[0]) - 1,
	}
	nodes := make([][]*Node, cave.yMax+1)
	for i := cave.yMax; i >= 0; i-- {
		nodes[i] = make([]*Node, (cave.xMax + 1))
		for j := cave.xMax; j >= 0; j-- {
			n := newNode(i, j, val(chiton[j][i]))
			nodes[i][j] = n
		}
	}
	cave.nodes = nodes
	cave.dijkstra()
	return cave.nodes[cave.yMax][cave.xMax].distance
}

func buildBiggerChiton(chiton []string, factor int) []string {
	newChiton := make([]string, len(chiton)*factor)
	for i := 0; i < len(chiton)*factor; i++ {
		layer := make([]rune, len(chiton[0])*factor)
		for j := 0; j < len(chiton[0])*factor; j++ {
			value := (val(chiton[i%len(chiton)][j%len(chiton[0])]) + i/len(chiton) + j/len(chiton)) % 9
			if value == 0 {
				value = 9
			}
			layer[j] = rune(value + runeOffset)
		}
		newChiton[i] = string(layer)
	}
	return newChiton
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	chiton := []string{}
	for scanner.Scan() {
		chiton = append(chiton, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	bigger := buildBiggerChiton(chiton, 5)
	return AvoidChiton(chiton), AvoidChiton(bigger)
}
