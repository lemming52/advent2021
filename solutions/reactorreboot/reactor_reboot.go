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
	name int
}

type ContestedRegion struct {
	parent *Cuboid
	region [][][]bool
}

func (c *Cuboid) volume() int {
	return (c.x1 + 1 - c.x0) * (c.y1 + 1 - c.y0) * (c.z1 + 1 - c.z0)
}

func newCuboid(s string, pattern *regexp.Regexp, name int) (*Cuboid, error) {
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
		name: name,
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
				return append(newCuboids, overlapCuboids...)
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
	return append([]*Cuboid{original}, newCuboids...)
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
	return newCuboids
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
	cuboids := []*Cuboid{}
	contestedRegions := map[int][]*ContestedRegion{}
	lightsOn := 0

	for i, ins := range instructions {
		cuboid, err := newCuboid(ins, pattern, i)
		if err != nil {
			return 0, err
		}
		for _, c := range cuboids {
			if cuboidsOverlap(cuboid, c) {
				overlap := findOverlapDimensions(cuboid, c)
				contested, ok := contestedRegions[c.name]
				if !ok {
					continue
				}
				for _, cn := range contested {
					if !cuboidsOverlap(cuboid, cn) {
						continue
					}
					if cn.on
				}
			}
		}
			lightsOn += cuboid.volume() - overlapCount
		} else {

		}

		cuboids = append(cuboids, cuboid)
	}
	total := 0
	for _, i := range onInstructions {
		total += i.volume()
	}
	return total, nil
}
