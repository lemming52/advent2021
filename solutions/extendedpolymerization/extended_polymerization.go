package extendedpolymerization

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

const instructionPattern = `(\w{2}) -> (\w)`

type PolymerFactory struct {
	polymer    []rune
	insertions map[string]rune
	counts     map[rune]int
}

func newPolymerFactory(polymer string, insertions []string) (*PolymerFactory, error) {
	pattern, err := regexp.Compile(instructionPattern)
	if err != nil {
		return nil, err
	}
	counts := map[rune]int{}
	for _, r := range polymer {
		counts[r]++
	}
	p := &PolymerFactory{
		polymer:    []rune(polymer),
		insertions: map[string]rune{},
		counts:     counts,
	}
	for _, i := range insertions {
		p.parseInsertion(i, pattern)
	}
	return p, nil
}

func (p *PolymerFactory) parseInsertion(s string, pattern *regexp.Regexp) {
	components := pattern.FindStringSubmatch(s)
	p.insertions[components[1]] = rune(components[2][0])
}

func (p *PolymerFactory) extend() {
	newPolymer := []rune{}
	for i := 0; i < len(p.polymer)-1; i++ {
		pair := string(p.polymer[i : i+2])
		r, ok := p.insertions[pair]
		if ok {
			newPolymer = append(newPolymer, p.polymer[i], r)
			p.counts[r]++
		} else {
			newPolymer = append(newPolymer, p.polymer[i])
		}
	}
	newPolymer = append(newPolymer, p.polymer[len(p.polymer)-1])
	p.polymer = newPolymer
}

func (p *PolymerFactory) calculate() int {
	min, max := 0, 0
	for _, v := range p.counts {
		if v > max {
			max = v
			continue
		}
		if v < min || min == 0 {
			min = v
		}
	}
	return max - min
}

func ExtendPolymer(polymer string, subs []string, steps int) (int, error) {
	p, err := newPolymerFactory(polymer, subs)
	if err != nil {
		return 0, err
	}
	for i := 0; i < steps; i++ {
		p.extend()
	}
	return p.calculate(), nil
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	polymer := scanner.Text()
	scanner.Scan()

	subs := []string{}
	for scanner.Scan() {
		subs = append(subs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	a, err := ExtendPolymer(polymer, subs, 10)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ExtendPolymer(polymer, subs, 40)
	if err != nil {
		log.Fatal(err)
	}
	return a, b
}
