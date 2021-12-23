package reactorreboot

import (
	"fmt"
	"regexp"
	"strconv"
)

const instructionPattern = `(\w{2,3}) x=(\-?\d{1,5})\.\.(\-?\d{1,5})\,y=(\-?\d{1,5})\.\.(\-?\d{1,5})\,z=(\-?\d{1,5})\.\.(\-?\d{1,5})`

type Cuboid struct {
	x0, x1, y0, y1, z0, z1 int
	on                     bool
}

func (c *Cuboid) volume() int {
	return (c.x1 + 1 - c.x0) * (c.y1 + 1 - c.y0) * (c.z1 + 1 - c.z0)
}

func newCuboid(s string, pattern *regexp.Regexp) (*Cuboid, error) {
	components := pattern.FindStringSubmatch(s)
	values := make([]int, 6)
	for i, c := range components[2:] {
		val, err := strconv.Atoi(c)
		if err != nil {
			return nil, err
		}
		values[i] = val
	}
	c := &Cuboid{
		on: components[1] == "on",
	}
	if values[0] < values[1] {
		c.x0, c.x1 = values[0], values[1]
	} else {
		c.x0, c.x1 = values[1], values[0]
	}
	if values[2] < values[3] {
		c.y0, c.y1 = values[2], values[3]
	} else {
		c.y0, c.y1 = values[3], values[2]
	}
	if values[4] < values[5] {
		c.z0, c.z1 = values[4], values[5]
	} else {
		c.z0, c.z1 = values[5], values[4]
	}
	return c, nil
}

func cuboidsOverlap(a, b *Cuboid) bool {
	return (overlappingRange(a.x0, a.x1, b.x0, b.x1) &&
		overlappingRange(a.y0, a.y1, b.y0, b.y1) &&
		overlappingRange(a.z0, a.z1, b.z0, b.z1))
}

func totalOverlap(a, b *Cuboid) bool {
	return (totalOverlappingRange(a.x0, a.x1, b.x0, b.x1) &&
		totalOverlappingRange(a.y0, a.y1, b.y0, b.y1) &&
		totalOverlappingRange(a.z0, a.z1, b.z0, b.z1))
}

func overlappingRange(a, b, c, d int) bool {
	return (c <= a && a <= d) || (c <= b && b <= d)
}

func totalOverlappingRange(a, b, c, d int) bool {
	return (c <= a && a <= d) && (c <= b && b <= d)
}

func overlappingFaces(a, smaller *Cuboid) []bool {
	overlaps := make([]bool, 6)
	if a.x0 < smaller.x0 {
		overlaps[0] = true
	}
	if a.x1 > smaller.x1 {
		overlaps[1] = true
	}
	if a.y0 < smaller.y0 {
		overlaps[2] = true
	}
	if a.y1 > smaller.y1 {
		overlaps[3] = true
	}
	if a.z0 < smaller.z0 {
		overlaps[4] = true
	}
	if a.z1 > smaller.z1 {
		overlaps[5] = true
	}
	return overlaps
}

func findOverlaps(a *Cuboid, existing []*Cuboid) []*Cuboid {
	if len(existing) == 0 {
		return []*Cuboid{a}
	}
	newCuboids := []*Cuboid{}
	overlapped := false
	for i, on := range existing {
		if cuboidsOverlap(a, on) {
			overlapped = true
			if a.on {
				overlapCuboids := []*Cuboid{}
				overlaps := buildNewOnCuboids(a, on)
				for _, c := range overlaps {
					overlapCuboids = append(overlapCuboids, findOverlaps(c, existing[i+1:])...)
				}
				return combineCuboids(append(newCuboids, overlapCuboids...))
			} else {
				newCuboids = append(newCuboids, splitOldOnCuboid(a, on)...)
			}
		} else {
			newCuboids = append(newCuboids, on)
		}
	}
	if !overlapped {
		newCuboids = append(newCuboids, a)
	}
	return newCuboids
}

func buildNewOnCuboids(a, original *Cuboid) []*Cuboid {
	if totalOverlap(a, original) {
		if a.x0 < original.x0 {
			return []*Cuboid{a}
		}
		return []*Cuboid{original}
	}
	overlap := findOverlapDimensions(a, original)
	overlappingDimensions := overlappingFaces(a, overlap)
	newCuboids := expandOverlapFaces(overlappingDimensions, a, overlap)
	return combineCuboids(append([]*Cuboid{original}, newCuboids...))
}

func splitOldOnCuboid(a, original *Cuboid) []*Cuboid {
	if totalOverlap(a, original) {
		if a.x0 < original.x0 {
			return []*Cuboid{}
		}
	}
	overlap := findOverlapDimensions(a, original)
	overlappingDimensions := overlappingFaces(original, overlap)
	newCuboids := expandOverlapFaces(overlappingDimensions, original, overlap)
	return combineCuboids(newCuboids)
}

func combineCuboids(cuboids []*Cuboid) []*Cuboid {
	for i, c := range cuboids[:len(cuboids)-1] {
		for j, cx := range cuboids[i+1:] {
			if shareFace(c, cx) {
				cn := combine(c, cx)
				newArray := append([]*Cuboid{}, cuboids[:i]...)
				newArray = append(newArray, cn)
				newArray = append(newArray, cuboids[i+1:i+j+1]...)
				if j+i != len(cuboids)-1 {
					newArray = append(newArray, cuboids[i+j+2:]...)
				}
				return combineCuboids(newArray)
			}
		}
	}
	return cuboids
}

func shareFace(a, b *Cuboid) bool {
	sharedCount := 0
	xAxis, yAxis, zAxis := false, false, false
	if a.x0 == b.x0 {
		xAxis = true
		sharedCount++
	}
	if a.x1 == b.x1 {
		xAxis = true
		sharedCount++
	}
	if a.y0 == b.y0 {
		yAxis = true
		sharedCount++
	}
	if a.y1 == b.y1 {
		yAxis = true
		sharedCount++
	}
	if a.z0 == b.z0 {
		zAxis = true
		sharedCount++
	}
	if a.z1 == b.z1 {
		zAxis = true
		sharedCount++
	}
	if sharedCount >= 4 {
		if !xAxis {
			return a.x1 == b.x0-1 || b.x1 == a.x0-1
		}
		if !yAxis {
			return a.y1 == b.y0-1 || b.y1 == a.y0-1
		}
		if !zAxis {
			return a.z1 == b.z0-1 || b.z1 == a.z0-1
		}
	}
	return false
}

func combine(a, b *Cuboid) *Cuboid {
	c := &Cuboid{
		x0: a.x0,
		x1: a.x1,
		y0: a.y0,
		y1: a.y1,
		z0: a.z0,
		z1: a.z1,
		on: a.on,
	}
	if a.x1 == b.x0-1 {
		c.x1 = b.x1
	}
	if b.x1 == a.x0-1 {
		c.x0 = b.x0
	}
	if a.y1 == b.y0-1 {
		c.y1 = b.y1
	}
	if b.y1 == a.y0-1 {
		c.y0 = b.y0
	}
	if a.z1 == b.z0-1 {
		c.z1 = b.z1
	}
	if b.z1 == a.z0-1 {
		c.z0 = b.z0
	}
	return c
}

func findOverlapDimensions(a, b *Cuboid) *Cuboid {
	overlap := &Cuboid{}
	if a.x0 >= b.x0 {
		if a.x1 <= b.x1 {
			// a entirely within
			overlap.x0, overlap.x1 = a.x0, a.x1
		} else {
			overlap.x0, overlap.x1 = a.x0, b.x1
		}
	} else {
		overlap.x0, overlap.x1 = b.x0, a.x1
	}
	if a.y0 >= b.y0 {
		if a.y1 <= b.y1 {
			// a entirely within
			overlap.y0, overlap.y1 = a.y0, a.y1
		} else {
			overlap.y0, overlap.y1 = a.y0, b.y1
		}
	} else {
		overlap.y0, overlap.y1 = b.y0, a.y1
	}
	if a.z0 >= b.z0 {
		if a.z1 <= b.z1 {
			// a entirely within
			overlap.z0, overlap.z1 = a.z0, a.z1
		} else {
			overlap.z0, overlap.z1 = a.z0, b.z1
		}
	} else {
		overlap.z0, overlap.z1 = b.z0, a.z1
	}
	return overlap
}

func Reboot(instructions []string) (int, error) {
	pattern, err := regexp.Compile(instructionPattern)
	if err != nil {
		return 0, err
	}
	var onInstructions []*Cuboid
	for _, i := range instructions {
		cuboid, err := newCuboid(i, pattern)
		if err != nil {
			return 0, err
		}
		if onInstructions == nil {
			if cuboid.on {
				onInstructions = []*Cuboid{cuboid}
			}
			continue
		}
		onInstructions = combineCuboids(findOverlaps(cuboid, onInstructions))
		fmt.Println("++++", len(onInstructions))
		for j, i := range onInstructions {
			fmt.Println(i, j)
		}
	}
	total := 0
	for _, i := range onInstructions {
		total += i.volume()
	}
	return total, nil
}
