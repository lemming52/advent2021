package passagepathing

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

type Cave struct {
	name      string
	large     bool
	connected []*Cave
}

func newCave(a string) *Cave {
	return &Cave{
		name:      a,
		large:     isUpper(a),
		connected: []*Cave{},
	}
}

func isUpper(s string) bool {
	for _, a := range s {
		if !unicode.IsUpper(a) {
			return false
		}
	}
	return true
}

func (c *Cave) addPath(cave *Cave) {
	c.connected = append(c.connected, cave)
}

func (c *Cave) exploreCave(visited string) int {
	paths := 0
	if c.name == "end" {
		return 1
	}
	visited = visited + c.name
	for _, cave := range c.connected {
		if cave.large {
			paths += cave.exploreCave(visited)
			continue
		}
		if !strings.Contains(visited, cave.name) {
			paths += cave.exploreCave(visited)
		}
	}
	return paths
}

func (c *Cave) exploreCaveTwice(visited string, smallCaveVisited bool) int {
	paths := 0
	if c.name == "end" {
		return 1
	}
	if c.name == "start" && visited != "" {
		return 0
	}
	visited = visited + c.name
	for _, cave := range c.connected {
		if cave.large {
			paths += cave.exploreCaveTwice(visited, smallCaveVisited)
			continue
		}
		if !strings.Contains(visited, cave.name) {
			paths += cave.exploreCaveTwice(visited, smallCaveVisited)
			continue
		}
		if !smallCaveVisited {
			paths += cave.exploreCaveTwice(visited, true)
		}
	}
	return paths
}

type CaveSystem struct {
	caves map[string]*Cave
}

func (c *CaveSystem) addPath(a, b string) {
	aCave, ok := c.caves[a]
	if !ok {
		aCave = newCave(a)
		c.caves[a] = aCave
	}
	bCave, ok := c.caves[b]
	if !ok {
		bCave = newCave(b)
		c.caves[b] = bCave
	}
	aCave.addPath(bCave)
	bCave.addPath(aCave)
}

func MapCaves(paths []string) (int, int) {
	caves := &CaveSystem{
		caves: map[string]*Cave{},
	}
	for _, connection := range paths {
		ends := strings.Split(connection, "-")
		caves.addPath(ends[0], ends[1])
	}
	return caves.caves["start"].exploreCave(""), caves.caves["start"].exploreCaveTwice("", false)
}

func Challenge(path string) (int, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	paths := []string{}
	for scanner.Scan() {
		paths = append(paths, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return MapCaves(paths)
}
