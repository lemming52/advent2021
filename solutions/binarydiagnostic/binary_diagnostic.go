package binarydiagnostic

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type DiagnosticTreeNode struct {
	number     string
	childCount int
	sum        int
	one        *DiagnosticTreeNode
	zero       *DiagnosticTreeNode
}

func newDiagnosticTreeNode(number string) *DiagnosticTreeNode {
	return &DiagnosticTreeNode{
		childCount: 0,
		sum:        0,
		number:     number,
	}
}

func Diagnose(numbers []string) (int64, int64) {
	counts := map[int]int{}
	rootNode := &DiagnosticTreeNode{}
	for _, number := range numbers {
		currentNode := rootNode
		for i, b := range number {
			currentNode.childCount++
			if b == '1' {
				counts[i]++
				currentNode.sum++
				if currentNode.one == nil {
					currentNode.one = newDiagnosticTreeNode(currentNode.number + "1")
				}
				currentNode = currentNode.one
			} else {
				if currentNode.zero == nil {
					currentNode.zero = newDiagnosticTreeNode(currentNode.number + "0")
				}
				currentNode = currentNode.zero
			}
		}

	}
	epsilon, gamma := make([]string, len(counts)), make([]string, len(counts))
	halfLength := len(numbers) / 2
	for k, v := range counts {
		if v > halfLength {
			gamma[k] = "1"
			epsilon[k] = "0"
		} else {
			gamma[k] = "0"
			epsilon[k] = "1"
		}
	}
	gammaRate, err := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	if err != nil {
		return 0, 0
	}
	epsilonRate, err := strconv.ParseInt(strings.Join(epsilon, ""), 2, 64)
	if err != nil {
		return 0, 0
	}
	return gammaRate * epsilonRate, evaluateTree(rootNode)
}

func evaluateTree(node *DiagnosticTreeNode) int64 {
	ox := computeNumber(node, true)
	co2 := computeNumber(node, false)
	return ox * co2
}

func computeNumber(node *DiagnosticTreeNode, flag bool) int64 {
	for node.childCount > 1 {
		if node.sum > node.childCount/2 {
			if flag {
				node = node.one
			} else {
				node = node.zero
			}
		} else if node.sum == node.childCount-node.sum {
			if flag {
				node = node.one
			} else {
				node = node.zero
			}
		} else {
			if flag {
				node = node.zero
			} else {
				node = node.one
			}
		}
	}
	for node.childCount != 0 {
		if node.one != nil {
			node = node.one
		} else {
			node = node.zero
		}
	}
	val, err := strconv.ParseInt(node.number, 2, 64)
	if err != nil {
		return 0
	}
	return val
}

func LoadBD(path string) (int64, int64) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []string{}
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return Diagnose(numbers)
}
