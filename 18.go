package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type pos struct{ x, y, z int }

func (a pos) max(b pos) pos {
	return pos{max(a.x, b.x), max(a.y, b.y), max(a.z, b.z)}
}

func (a pos) min(b pos) pos {
	return pos{min(a.x, b.x), min(a.y, b.y), min(a.z, b.z)}
}

type posset map[pos]bool

func (ps posset) contains(p pos) bool {
	_, found := ps[p]
	return found
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cubes := posset{}
	for scanner.Scan() {
		var cube pos
		fmt.Sscanf(scanner.Text(), "%d,%d,%d", &cube.x, &cube.y, &cube.z)
		cubes[cube] = true
	}
	surface := 0
	bounds := []pos{pos{}, pos{}}
	for cube, _ := range cubes {
		if !cubes.contains(pos{cube.x + 1, cube.y, cube.z}) {
			surface += 1
		}
		if !cubes.contains(pos{cube.x - 1, cube.y, cube.z}) {
			surface += 1
		}
		if !cubes.contains(pos{cube.x, cube.y - 1, cube.z}) {
			surface += 1
		}
		if !cubes.contains(pos{cube.x, cube.y + 1, cube.z}) {
			surface += 1
		}
		if !cubes.contains(pos{cube.x, cube.y, cube.z - 1}) {
			surface += 1
		}
		if !cubes.contains(pos{cube.x, cube.y, cube.z + 1}) {
			surface += 1
		}
		bounds[0] = bounds[0].min(cube)
		bounds[1] = bounds[1].max(cube)
	}
	bounds[0].x -= 1
	bounds[0].y -= 1
	bounds[0].z -= 1
	bounds[1].x += 1
	bounds[1].y += 1
	bounds[1].z += 1
	external := posset{}
	queue := []pos{bounds[0]}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		if external.contains(c) || cubes.contains(c) ||
			c.x < bounds[0].x || c.y < bounds[0].y || c.z < bounds[0].z ||
			c.x > bounds[1].x || c.y > bounds[1].y || c.z > bounds[1].z {
			continue
		}
		external[c] = true
		queue = append(queue, pos{c.x - 1, c.y, c.z})
		queue = append(queue, pos{c.x + 1, c.y, c.z})
		queue = append(queue, pos{c.x, c.y - 1, c.z})
		queue = append(queue, pos{c.x, c.y + 1, c.z})
		queue = append(queue, pos{c.x, c.y, c.z - 1})
		queue = append(queue, pos{c.x, c.y, c.z + 1})
	}
	esurface := 0
	for c, _ := range cubes {
		if external.contains(pos{c.x + 1, c.y, c.z}) {
			esurface += 1
		}
		if external.contains(pos{c.x - 1, c.y, c.z}) {
			esurface += 1
		}
		if external.contains(pos{c.x, c.y + 1, c.z}) {
			esurface += 1
		}
		if external.contains(pos{c.x, c.y - 1, c.z}) {
			esurface += 1
		}
		if external.contains(pos{c.x, c.y, c.z + 1}) {
			esurface += 1
		}
		if external.contains(pos{c.x, c.y, c.z - 1}) {
			esurface += 1
		}
	}
	fmt.Println(surface, esurface)
}
