package transparentorigami

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const foldPattern = `fold along ([x,y])=(\d{1,3})`

type Origami struct {
	dots        map[string][]int
	foldPattern *regexp.Regexp
}

func (o *Origami) fold(instruction string) error {
	components := o.foldPattern.FindStringSubmatch(instruction)
	val, err := strconv.Atoi(components[2])
	if err != nil {
		return err
	}
	if components[1] == "x" {
		o.foldX(val)
	} else {
		o.foldY(val)
	}
	return nil
}

func (o *Origami) foldX(val int) {
	for key, coords := range o.dots {
		if coords[0] == val {
			delete(o.dots, key)
		} else if coords[0] > val {
			newCoords := []int{val - (coords[0] - val), coords[1]}
			o.dots[fmt.Sprintf("%d,%d", newCoords[0], newCoords[1])] = newCoords
			delete(o.dots, key)
		}
	}
}

func (o *Origami) foldY(val int) {
	for key, coords := range o.dots {
		if coords[1] == val {
			delete(o.dots, key)
		} else if coords[1] > val {
			newCoords := []int{coords[0], val - (coords[1] - val)}
			o.dots[fmt.Sprintf("%d,%d", newCoords[0], newCoords[1])] = newCoords
			delete(o.dots, key)
		}
	}
}

func (o *Origami) Print() {
	xMax, yMax := 0, 0
	for _, v := range o.dots {
		if v[0] > xMax {
			xMax = v[0]
		}
		if v[1] > yMax {
			yMax = v[1]
		}
	}
	grid := make([][]string, yMax+1)
	for i := 0; i <= yMax; i++ {
		grid[i] = make([]string, xMax+1)
	}
	for _, v := range o.dots {
		grid[v[1]][v[0]] = "#"
	}
	delimiter := ""
	for i := 0; i <= xMax; i++ {
		delimiter += "-"
	}
	fmt.Println(delimiter)
	for _, y := range grid {
		s := ""
		for _, x := range y {
			if x == "" {
				s += " "
			} else {
				s += x
			}
		}
		fmt.Println(s)
	}
	fmt.Println(delimiter)
}

func FoldOrigami(dots, folds []string) (int, int, error) {
	pattern, err := regexp.Compile(foldPattern)
	if err != nil {
		return 0, 0, err
	}
	origami := &Origami{
		dots:        make(map[string][]int, len(dots)),
		foldPattern: pattern,
	}
	for _, d := range dots {
		components := strings.Split(d, ",")
		coords := make([]int, len(components))
		for i, c := range components {
			val, err := strconv.Atoi(c)
			if err != nil {
				return 0, 0, err
			}
			coords[i] = val
			origami.dots[d] = coords
		}
	}

	var first int
	for i, f := range folds {
		err = origami.fold(f)
		if err != nil {
			return 0, 0, err
		}
		if i == 0 {
			first = len(origami.dots)
		}
	}
	origami.Print()
	return first, len(origami.dots), nil
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	dots := []string{}
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			break
		}
		dots = append(dots, l)
	}

	folds := []string{}
	for scanner.Scan() {
		folds = append(folds, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	first, total, err := FoldOrigami(dots, folds)
	if err != nil {
		log.Fatal(err)
	}
	return first, total
}
