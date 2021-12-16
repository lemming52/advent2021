package packetdecoder

import (
	"advent/solutions/utils"
	"log"
	"math"
	"strconv"
	"strings"
)

type opFunc func(a, b int) int

type Parser struct {
	hexMap     map[rune]string
	binaryMap  map[string]int
	versionSum int
}

func newParser() *Parser {
	return &Parser{
		hexMap: map[rune]string{
			'0': "0000",
			'1': "0001",
			'2': "0010",
			'3': "0011",
			'4': "0100",
			'5': "0101",
			'6': "0110",
			'7': "0111",
			'8': "1000",
			'9': "1001",
			'A': "1010",
			'B': "1011",
			'C': "1100",
			'D': "1101",
			'E': "1110",
			'F': "1111",
		},
		binaryMap:  binaryMap(),
		versionSum: 0,
	}
}

func binaryMap() map[string]int {
	return map[string]int{
		"000": 0,
		"001": 1,
		"010": 2,
		"011": 3,
		"100": 4,
		"101": 5,
		"110": 6,
		"111": 7,
	}
}

func (p *Parser) parse(s string) (int, int, error) {
	p.versionSum += p.binaryMap[s[:3]]
	switch s[3:6] {
	case "100": // literal value
		val, bitsRead, err := parseLiteral(s[6:])
		if err != nil {
			return 0, 0, err
		}
		return val, bitsRead + 6, nil
	default: // operator
		val, bitsRead, err := p.parseOperator(s[3:6], s[6:])
		if err != nil {
			return 0, 0, err
		}
		return val, bitsRead + 6, nil
	}
}

func parseLiteral(s string) (int, int, error) {
	index := 0
	var sb strings.Builder
	for s[index] != '0' {
		sb.WriteString(s[index+1 : index+5])
		index += 5
	}
	sb.WriteString(s[index+1 : index+5])
	val, err := strconv.ParseInt(sb.String(), 2, 64)
	return int(val), index + 5, err
}

func (p *Parser) parseOperator(op, s string) (int, int, error) {
	operator, initial := getOperator(op)
	switch s[0] {
	case '0': // bit length
		length, err := strconv.ParseInt(s[1:16], 2, 64)
		if err != nil {
			return 0, 0, err
		}
		index := 16
		toRead := int(length + 16)
		for index < toRead {
			val, bitsRead, err := p.parse(s[index:])
			if err != nil {
				return 0, 0, err
			}
			index += bitsRead
			initial = operator(initial, val)
		}
		return initial, index, nil
	case '1': // subpacket count
		count, err := strconv.ParseInt(s[1:12], 2, 64)
		if err != nil {
			return 0, 0, err
		}
		index := 12
		for i := 0; i < int(count); i++ {
			val, bitsRead, err := p.parse(s[index:])
			if err != nil {
				return 0, 0, err
			}
			index += bitsRead
			initial = operator(initial, val)
		}
		return initial, index, nil
	}
	return 0, 0, nil
}

func getOperator(op string) (opFunc, int) {
	switch op {
	case "000":
		return sum, 0
	case "001":
		return product, 1
	case "010":
		return minimum, math.MaxInt
	case "011":
		return maximum, 0
	case "101":
		return greater, -1
	case "110":
		return less, -1
	case "111":
		return equal, -1
	}
	return equal, 0
}

func sum(a, b int) int {
	return a + b
}

func product(a, b int) int {
	return a * b
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func greater(a, b int) int {
	if a == -1 {
		return b
	}
	if a > b {
		return 1
	}
	return 0
}

func less(a, b int) int {
	if a == -1 {
		return b
	}
	if a < b {
		return 1
	}
	return 0
}

func equal(a, b int) int {
	if a == -1 {
		return b
	}
	if a == b {
		return 1
	}
	return 0
}

func ParsePackets(s string) (int, int, error) {
	parser := newParser()
	var sb strings.Builder
	for _, r := range s {
		sb.WriteString(parser.hexMap[r])
	}
	val, _, err := parser.parse(sb.String())
	return parser.versionSum, val, err
}

func Challenge(path string) (int, int) {
	packets, err := utils.ReadStrings(path)
	if err != nil {
		log.Fatal(err)
	}
	a, b, err := ParsePackets(packets[0])
	if err != nil {
		log.Fatal(err)
	}
	return a, b
}
