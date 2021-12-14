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
	created    map[string][]string
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
		created:    map[string][]string{},
	}
	for _, i := range insertions {
		p.parseInsertion(i, pattern)
	}
	return p, nil
}

func (p *PolymerFactory) parseInsertion(s string, pattern *regexp.Regexp) {
	components := pattern.FindStringSubmatch(s)
	p.insertions[components[1]] = rune(components[2][0])
	p.created[components[1]] = []string{
		string(components[1][0]) + components[2],
		components[2] + string(components[1][1]),
	}
}

func (p *PolymerFactory) extend(pairs map[string]int, steps int) {
	for steps > 0 {
		newPairs := make(map[string]int, len(pairs))
		for pair, count := range pairs {
			newPairs[pair] += count
			if count == 0 {
				continue
			}
			r, ok := p.insertions[pair]
			if !ok {
				continue
			}
			p.counts[r] += count
			newPairs[pair] -= count
			c, ok := p.created[pair]
			if !ok {
				continue
			}
			newPairs[c[0]] += count
			newPairs[c[1]] += +count
		}
		steps--
		pairs = newPairs
	}
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
	pairs := map[string]int{}
	for i := 0; i < len(polymer)-1; i++ {
		pairs[string(polymer[i:i+2])]++
	}
	p.extend(pairs, steps)
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
