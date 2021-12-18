package snailmaths

import (
	"advent/solutions/utils"
	"fmt"
	"log"
	"strconv"
)

type Number struct {
	value               int
	depth               int
	left, right, parent *Number
}

func newNumber(s string, depth int, parent *Number) *Number {
	n := &Number{
		value:  -1,
		depth:  depth,
		parent: parent,
	}
	breakIndex := findMidpoint(s)
	l, r := s[1:breakIndex], s[breakIndex+1:len(s)-1]
	n.left = parseNumberComponent(l, depth, n)
	n.right = parseNumberComponent(r, depth, n)
	return n
}

func findMidpoint(s string) int {
	openCount := 0
	for i, r := range s[1 : len(s)-1] {
		switch r {
		case '[':
			openCount++
		case ']':
			openCount--
		case ',':
			if openCount == 0 {
				return i + 1
			}

		default:
		}
	}
	return 0
}

func parseNumberComponent(s string, depth int, n *Number) *Number {
	val, err := strconv.Atoi(s)
	if err != nil {
		return newNumber(s, depth+1, n)
	} else {
		return &Number{
			value:  val,
			depth:  depth + 1,
			parent: n,
		}
	}
}

func (n *Number) nest() {
	n.depth++
	if n.value != -1 {
		return
	}
	n.left.nest()
	n.right.nest()
}

func (n *Number) reduce() bool {
	exploded := n.reduceExplode()
	if exploded {
		return true
	}
	return n.reduceSplit()
}

func (n *Number) reduceExplode() bool {
	if n.depth > 3 && n.value == -1 {
		n.explode()
		return true
	}
	if n.value != -1 {
		return false
	}
	leftAct := n.left.reduceExplode()
	if leftAct {
		return leftAct
	}
	return n.right.reduceExplode()
}

func (n *Number) reduceSplit() bool {
	if n.value >= 10 {
		n.split()
		return true
	}
	if n.value != -1 {
		return false
	}
	leftAct := n.left.reduceSplit()
	if leftAct {
		return leftAct
	}
	return n.right.reduceSplit()
}

func (n *Number) explode() {
	n.incrementAdjacentNodes(true, n.left.value)
	n.incrementAdjacentNodes(false, n.right.value)
	n.value = 0
	n.left = nil
	n.right = nil
}

func (n *Number) incrementAdjacentNodes(left bool, value int) {
	if n.parent == nil {
		return
	}
	if left {
		if n == n.parent.right {
			n.parent.left.increment(false, value)
			return
		}
		n.parent.incrementAdjacentNodes(left, value)
		return
	}
	if n == n.parent.left {
		n.parent.right.increment(true, value)
		return
	}
	n.parent.incrementAdjacentNodes(left, value)
	return
}

func (n *Number) increment(incrementLeft bool, value int) {
	if n.value != -1 {
		n.value += value
		return
	}
	if incrementLeft {
		n.left.increment(true, value)
	} else {
		n.right.increment(false, value)
	}
}

func (n *Number) split() {
	val := n.value / 2
	rightVal := n.value - val
	n.left = &Number{
		value:  val,
		depth:  n.depth + 1,
		parent: n,
	}
	n.right = &Number{
		value:  rightVal,
		depth:  n.depth + 1,
		parent: n,
	}
	n.value = -1
}

func (n *Number) magnitude() int {
	if n.value != -1 {
		return n.value
	}
	return 3*n.left.magnitude() + 2*n.right.magnitude()
}

func (n *Number) print() string {
	if n.value != -1 {
		return fmt.Sprintf("%d{%d}", n.value, n.depth)
	}
	return fmt.Sprintf("[%s,%s]", n.left.print(), n.right.print())

}

func add(a, b *Number) *Number {
	n := &Number{
		value: -1,
		depth: 0,
		left:  a,
		right: b,
	}
	a.parent = n
	b.parent = n
	a.nest()
	b.nest()
	for n.reduce() {
	}
	return n
}

func SnailSum(numbers []string) (int, int) {
	n := newNumber(numbers[0], 0, nil)
	for _, num := range numbers[1:] {
		newNum := newNumber(num, 0, nil)
		n = add(n, newNum)
	}
	maxMag := 0
	for i, a := range numbers {
		for _, b := range numbers[i:] {
			aNum, bNum := newNumber(a, 0, nil), newNumber(b, 0, nil)
			mag := add(aNum, bNum).magnitude()
			if mag > maxMag {
				maxMag = mag
			}
			aNum, bNum = newNumber(a, 0, nil), newNumber(b, 0, nil)
			mag = add(bNum, aNum).magnitude()
			if mag > maxMag {
				maxMag = mag
			}
		}
	}
	return n.magnitude(), maxMag
}

func Challenge(path string) (int, int) {
	numbers, err := utils.ReadStrings(path)
	if err != nil {
		log.Fatal(err)
	}
	return SnailSum(numbers)
}
