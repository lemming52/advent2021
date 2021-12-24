package alu

import (
	"advent/solutions/utils"
	"fmt"
	"log"
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

func FindModelNumber(instructions []string) (int, error) {
	var model int
	for i := 99999999999999; i > 11111111111111; i-- {
		s := fmt.Sprintf("%d", i)
		for _, r := range s {
			if r == '0' {
				continue
			}
		}
		processor := newALU([]string{"w", "x", "y", "z"}, s)
		for _, i := range instructions {
			err := processor.Execute(i)
			if err != nil {
				return 0, err
			}
		}
		if processor.registers["z"].value == 0 {
			model = i
		}
	}
	return model, nil
}

func Challenge(path string) (int, int) {
	lines, err := utils.ReadStrings(path)
	if err != nil {
		log.Fatal(err)
	}
	val, err := FindModelNumber(lines)
	if err != nil {
		log.Fatal(err)
	}
	return val, 0
}
