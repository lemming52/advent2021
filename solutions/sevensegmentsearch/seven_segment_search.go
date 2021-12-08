package sevensegmentsearch

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

const maxValue = 510510

type CombinationManager struct {
	values           map[rune]int
	simpleDigitCount int
	fullDigitCount   int
	lexicon          map[int]int
	options          map[int][]int
	digits           map[int]int
}

func newCombinationManager() *CombinationManager {
	return &CombinationManager{
		values: map[rune]int{
			'a': 2,
			'b': 3,
			'c': 5,
			'd': 7,
			'e': 11,
			'f': 13,
			'g': 17,
		},
		lexicon: map[int]int{},
		options: map[int][]int{
			5: []int{},
			6: []int{},
		},
		digits: map[int]int{},
	}
}

func (c *CombinationManager) calculateValue(word string) int {
	key := 1
	if len(word) == 8 {
		return maxValue
	}
	for _, r := range word {
		value := c.values[r]
		key *= value
	}
	return key
}

func (c *CombinationManager) addDigit(value, digit int) {
	c.lexicon[value] = digit
	c.digits[digit] = value
}

func (c *CombinationManager) countResults(digits []string) {
	total := 0
	for i, d := range digits {
		val := c.lexicon[c.calculateValue(d)]
		switch val {
		case 1, 4, 7, 8:
			c.simpleDigitCount++
		}
		total += int(math.Pow(10, float64(4-i))) * val
	}
	c.fullDigitCount += total
}

func (c *CombinationManager) resetLexicon() {
	c.lexicon = map[int]int{}
	c.digits = map[int]int{}
	c.options = map[int][]int{}
}

func (c *CombinationManager) addOption(w string) {
	c.options[len(w)] = append(c.options[len(w)], c.calculateValue(w))
}

func (c *CombinationManager) determineFullLexicon() {
	cf := c.digits[1]
	a := c.digits[7] / cf
	bd := c.digits[4] / cf
	eg := maxValue / (cf * bd * a)
	for _, w := range c.options[5] {
		if w%eg == 0 {
			c.addDigit(w, 2)
		} else if w%bd == 0 {
			c.addDigit(w, 5)
		} else {
			c.addDigit(w, 3)
		}
	}
	for _, w := range c.options[6] {
		if w%eg != 0 {
			c.addDigit(w, 9)
		} else if w%bd != 0 {
			c.addDigit(w, 0)
		} else {
			c.addDigit(w, 6)
		}
	}
}

func CountDigits(signals []string) (int, int) {
	calculator := newCombinationManager()
	for _, signal := range signals {
		lexicon, digits := parseSignal(signal)
		for _, w := range lexicon {
			switch len(w) {
			case 2:
				calculator.addDigit(calculator.calculateValue(w), 1)
			case 3:
				calculator.addDigit(calculator.calculateValue(w), 7)
			case 4:
				calculator.addDigit(calculator.calculateValue(w), 4)
			case 7:
				calculator.addDigit(calculator.calculateValue(w), 8)
			default:
				calculator.addOption(w)
				continue
			}
		}
		calculator.determineFullLexicon()
		calculator.countResults(digits)
		calculator.resetLexicon()
	}
	return calculator.simpleDigitCount, calculator.fullDigitCount
}

func parseSignal(s string) ([]string, []string) {
	components := strings.Split(s, "|")
	lexicon := strings.Split(components[0], " ")
	digits := strings.Split(components[1], " ")
	return lexicon, digits
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	signals := []string{}
	for scanner.Scan() {
		signals = append(signals, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	res1, res2 := CountDigits(signals)
	return res1, res2
}
