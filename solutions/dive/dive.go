package dive

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	depth int
	x     int
	aim   int
}

// Dive executes a set of commands to evaluate a final position
func Dive(commands []string, part2 bool) int {
	initial := &Position{
		depth: 0,
		x:     0,
	}
	for _, c := range commands {
		var err error
		if part2 {
			err = evaluateAlternateCommand(initial, c)
		} else {
			err = evaluateCommand(initial, c)
		}
		if err != nil {
			return 0
		}
	}
	return initial.depth * initial.x
}

func evaluateCommand(p *Position, s string) error {
	command := strings.Split(s, " ")
	value, err := strconv.Atoi(command[1])
	if err != nil {
		return err
	}
	switch command[0] {
	case "up":
		p.depth -= value
	case "down":
		p.depth += value
	case "forward":
		p.x += value
	}
	return nil
}

func evaluateAlternateCommand(p *Position, s string) error {
	command := strings.Split(s, " ")
	value, err := strconv.Atoi(command[1])
	if err != nil {
		return err
	}
	switch command[0] {
	case "up":
		p.aim -= value
	case "down":
		p.aim += value
	case "forward":
		p.x += value
		p.depth += p.aim * value
	}
	return nil
}

// LoadDive loads an input text file and executes the dive instructions
func LoadDive(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	commands := []string{}
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return Dive(commands, false), Dive(commands, true)
}
