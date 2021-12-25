package alu

import (
	"advent/solutions/utils"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

const runeOffset = 48

type Register struct {
	name  string
	value int
}

func (r *Register) set(v int) {
	r.value = v
}

func (r *Register) add(v int) {
	r.value += v
}

func (r *Register) mul(v int) {
	r.value *= v
}

func (r *Register) div(v int) {
	r.value /= v
}

func (r *Register) mod(v int) {
	r.value = r.value % v
}

func (r *Register) eql(v int) {
	if r.value == v {
		r.value = 1
	} else {
		r.value = 0
	}
}

type ArithmeticLogicUnit struct {
	registers  map[string]*Register
	input      string
	inputIndex int
}

func newALU(registers []string, input string) *ArithmeticLogicUnit {
	processor := &ArithmeticLogicUnit{
		registers:  map[string]*Register{},
		input:      input,
		inputIndex: 0,
	}
	for _, r := range registers {
		processor.registers[r] = &Register{name: r, value: 0}
	}
	return processor
}

func (p *ArithmeticLogicUnit) Execute(instruction string) error {
	components := strings.Split(instruction, " ")
	switch components[0] {
	case "inp":
		p.set(components[1])
	case "add":
		return p.add(components[1], components[2])
	case "div":
		return p.div(components[1], components[2])
	case "mul":
		return p.mul(components[1], components[2])
	case "mod":
		return p.mod(components[1], components[2])
	case "eql":
		return p.eql(components[1], components[2])
	}
	return nil
}

func (p *ArithmeticLogicUnit) set(register string) {
	p.registers[register].set(int(p.input[p.inputIndex] - runeOffset))
	p.inputIndex++
}

func (p *ArithmeticLogicUnit) add(register, variable string) error {
	val, err := p.extractValue(variable)
	if err != nil {
		return err
	}
	p.registers[register].add(val)
	return nil
}

func (p *ArithmeticLogicUnit) mul(register, variable string) error {
	val, err := p.extractValue(variable)
	if err != nil {
		return err
	}
	p.registers[register].mul(val)
	return nil
}

func (p *ArithmeticLogicUnit) div(register, variable string) error {
	val, err := p.extractValue(variable)
	if err != nil {
		return err
	}
	p.registers[register].div(val)
	return nil
}

func (p *ArithmeticLogicUnit) mod(register, variable string) error {
	val, err := p.extractValue(variable)
	if err != nil {
		return err
	}
	p.registers[register].mod(val)
	return nil
}

func (p *ArithmeticLogicUnit) eql(register, variable string) error {
	val, err := p.extractValue(variable)
	if err != nil {
		return err
	}
	p.registers[register].eql(val)
	return nil
}

func (p *ArithmeticLogicUnit) extractValue(variable string) (int, error) {
	switch variable {
	case "w", "x", "y", "z": // could be not hardcoded
		return p.registers[variable].value, nil
	default:
		return strconv.Atoi(variable)
	}
}

func (p *ArithmeticLogicUnit) print() {
	fmt.Println(p.input, p.inputIndex)
	for _, v := range p.registers {
		fmt.Println(v.name, v.value)
	}
}

type digitMap map[int]bool

// courtesy of a mate and u/livluc https://www.reddit.com/r/adventofcode/comments/rnejv5/2021_day_24_solutions/hpu84cj/?context=3
func monad(a, b, c int) func(w, z int) int {
	return func(w, z int) int {
		var gam int
		if ((z % 26) + b) != w {
			gam = 1
		}
		return gam*((z/a)*25+w+c) + (z / a)
	}
}

func reverseMonad(a, b, c int) func(z int, results map[int][][]int) map[int][][]int {
	return func(z int, results map[int][][]int) map[int][][]int {
		for i := 0; i < a; i++ {
			for j := 1; j <= 9; j++ {
				zi := a*z + i
				if (zi%26 + b) == j {
					_, ok := results[zi]
					if !ok {
						results[zi] = [][]int{{z, j}}
					} else {
						results[zi] = append(results[zi], []int{z, j})
					}
				}
				zi = (z-j-c)/(26*a) + i
				if (zi%26 + b) != j {
					if zi/a*26+j+c == z {
						_, ok := results[zi]
						if !ok {
							results[zi] = [][]int{{z, j}}
						} else {
							results[zi] = append(results[zi], []int{z, j})
						}
					}
				}
			}
		}
		return results
	}
}

func determineConstants(instructions []string) ([][]int, error) {
	relativeIndex := 0
	constants := make([][]int, 14)
	var a, b, c, counter int
	var err error
	for _, l := range instructions {
		components := strings.Split(l, " ")
		if components[0] == "inp" {
			relativeIndex = 0
			continue
		}
		if relativeIndex == 4 {
			b, err = strconv.Atoi(components[2])
			if err != nil {
				return nil, err
			}
		}
		if relativeIndex == 14 {
			c, err = strconv.Atoi(components[2])
			if err != nil {
				return nil, err
			}
			if b < 0 {
				a = 26
			} else {
				a = 1
			}
			constants[counter] = []int{a, b, c}
			counter++
		}
		relativeIndex++
	}
	return constants, nil
}

func DetermineModelNumber(instructions []string) (int, int, error) {
	constants, err := determineConstants(instructions)
	if err != nil {
		return 0, 0, err
	}
	reverseFunctions := make([]func(z int, results map[int][][]int) map[int][][]int, 14)
	for i, c := range constants {
		reverseFunctions[i] = reverseMonad(c[0], c[1], c[2])
	}
	solutions := make([]map[int][][]int, 14)
	solutions[13] = reverseFunctions[13](0, map[int][][]int{})
	for i := 12; i >= 0; i-- {
		solutions[i] = map[int][][]int{}
		for z := range solutions[i+1] {
			reverseFunctions[i](z, solutions[i])
		}
	}

	minS, maxS := "", ""
	zMin, zMax := 0, 0
	for i := 0; i <= 13; i++ {
		minimum := solutions[i][zMin]
		sort.Slice(minimum, func(i, j int) bool {
			return minimum[i][1] < minimum[j][1]
		})
		zMin = minimum[0][0]
		minS += fmt.Sprintf("%d", minimum[0][1])

		maximum := solutions[i][zMax]
		sort.Slice(maximum, func(i, j int) bool {
			return maximum[i][1] > maximum[j][1]
		})
		zMax = maximum[0][0]
		maxS += fmt.Sprintf("%d", maximum[0][1])
	}

	processor := newALU([]string{"w", "x", "y", "z"}, minS)
	for _, i := range instructions {
		err := processor.Execute(i)
		if err != nil {
			return 0, 0, err
		}
	}
	if processor.registers["z"].value != 0 {
		return 0, 0, errors.New("minimum value not correct")
	}

	processor = newALU([]string{"w", "x", "y", "z"}, maxS)
	for _, i := range instructions {
		err := processor.Execute(i)
		if err != nil {
			return 0, 0, err
		}
	}
	if processor.registers["z"].value != 0 {
		return 0, 0, errors.New("minimum value not correct")
	}
	valMin, err := strconv.Atoi(minS)
	if err != nil {
		return 0, 0, err
	}
	valMax, err := strconv.Atoi(maxS)
	if err != nil {
		return 0, 0, err
	}
	return valMax, valMin, nil
}

func Challenge(path string) (int, int) {
	lines, err := utils.ReadStrings(path)
	if err != nil {
		log.Fatal(err)
	}
	a, b, err := DetermineModelNumber(lines)
	if err != nil {
		log.Fatal(err)
	}
	return a, b
}
