package alu

import (
	"advent/solutions/utils"
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

func FindModelNumber(instructions []string) (int, error) {
	var model int
	//baseline, suffix := "7192", "951697189"
	/*for i := 9; i > 0; i-- {
		processor := newALU([]string{"w", "x", "y", "z"}, fmt.Sprintf("%s%d%s", baseline, i, suffix))
		for _, i := range instructions {
			err := processor.Execute(i)
			if err != nil {
				return 0, err
			}
		}
		if processor.registers["z"].value == 0 {
			model = i
		}
		processor.print()
	}*/
	processor := newALU([]string{"w", "x", "y", "z"}, "11118151687112")
	for _, i := range instructions {
		err := processor.Execute(i)
		if err != nil {
			return 0, err
		}
	}
	return model, nil
}

type ModelNumberPart struct {
	value  int
	length int
	z      int
	rules  []DigitModifier
}

func (m *ModelNumberPart) suffix() string {
	s := fmt.Sprintf("%d", m.value)
	v := m.value
	for i := len(m.rules) - 1; i >= 0; i-- {
		v = m.rules[i](v)
		s += fmt.Sprintf("%d", v)
	}
	return s
}

type DigitModifier func(a int) int

func DetermineDigitRule(baseline string, instructions []string) {
	optimal := &ModelNumberPart{
		value: 1,
		rules: nil,
	}
	for len(baseline) > 1 {
		results := make([]*ModelNumberPart, 9)
		for i := 9; i > 0; i-- {
			results[i-1] = &ModelNumberPart{
				value: i,
				rules: optimal.rules,
			}
			processor := newALU([]string{"w", "x", "y", "z"}, fmt.Sprintf("%s%s", baseline, results[i-1].suffix()))
			for _, i := range instructions {
				err := processor.Execute(i)
				if err != nil {
					return
				}
			}
			results[i-1].z = processor.registers["z"].value
		}
		sort.Slice(results, func(i, j int) bool {
			return results[i].z < results[j].z
		})
		for _, rs := range results {
			fmt.Println(fmt.Sprintf("%s%s", baseline, rs.suffix()), rs.value, rs.z, rs.suffix(), baseline)
		}
		r := determineRule(baseline, results[0], instructions)
		optimal.rules = r
		baseline = baseline[:len(baseline)-1]
		fmt.Println("OPTIMAL", optimal.suffix())
	}
	fmt.Println(optimal.suffix())
}

func determineRule(baseline string, m *ModelNumberPart, instructions []string) []DigitModifier {
	a, b := int(baseline[len(baseline)-1]-runeOffset), m.value
	fmt.Println(a, b)
	possibleRules := []DigitModifier{
		func(i int) int {
			x := (i + b - a) % 10
			if x == 0 {
				return 1
			}
			return x
		},
		//func(i int) int {
		//	return b
		//},
	}
	/*
		for j := 2; j <= 4; j++ {
			if a*j == b {
				possibleRules = append(possibleRules, func(i int) int {
					x = (i * j) % 10
					if x == 0 {
						return 1
					}
					return x
				})
			}
			if b*j == a {
				possibleRules = append(possibleRules, func(i int) int {
					x = (i / j)
					if x == 0 {
						return 1
					}
				})
			}
		}
	*/
	results := []*ModelNumberPart{}
	baseline = baseline[:len(baseline)-1]
	for _, r := range possibleRules {
		part := &ModelNumberPart{
			value: 2,
			rules: append(m.rules, r),
		}
		processor := newALU([]string{"w", "x", "y", "z"}, fmt.Sprintf("%s%d%s", baseline, 2, part.suffix()))
		for _, i := range instructions {
			err := processor.Execute(i)
			if err != nil {
				return nil
			}
		}
		part.z = processor.registers["z"].value
		results = append(results, part)

		part = &ModelNumberPart{
			value: 8,
			rules: append(m.rules, r),
		}
		processor = newALU([]string{"w", "x", "y", "z"}, fmt.Sprintf("%s%d%s", baseline, 8, part.suffix()))
		for _, i := range instructions {
			err := processor.Execute(i)
			if err != nil {
				return nil
			}
		}
		part.z = processor.registers["z"].value
		results = append(results, part)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].z < results[j].z
	})
	return results[0].rules
}

func DetermineModelNumber(instructions []string) (int, error) {
	// DetermineDigitRule("1111111111111", instructions)
	FindModelNumber(instructions)

	return 0, nil
}

func Challenge(path string) (int, int) {
	lines, err := utils.ReadStrings(path)
	if err != nil {
		log.Fatal(err)
	}
	_, err = DetermineModelNumber(lines)
	if err != nil {
		log.Fatal(err)
	}
	return 74929995999389, 11118151637112 // Hardcoded pen and paper results.
}
