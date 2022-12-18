package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct{ x, y, z int }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func contains(poss map[pos]bool, p pos) bool {
	_, found := poss[p]
	return found
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cubes := map[pos]bool{}
	for scanner.Scan() {
		var cube pos
		fmt.Sscanf(scanner.Text(), "%d,%d,%d", &cube.x, &cube.y, &cube.z)
		cubes[cube] = true
	}
	surface := 0
	for cube, _ := range cubes {
		surface += 6
		if _, found := cubes[pos{cube.x + 1, cube.y, cube.z}]; found {
			surface -= 1
		}
		if _, found := cubes[pos{cube.x - 1, cube.y, cube.z}]; found {
			surface -= 1
		}
		if _, found := cubes[pos{cube.x, cube.y + 1, cube.z}]; found {
			surface -= 1
		}
		if _, found := cubes[pos{cube.x, cube.y - 1, cube.z}]; found {
			surface -= 1
		}
		if _, found := cubes[pos{cube.x, cube.y, cube.z + 1}]; found {
			surface -= 1
		}
		if _, found := cubes[pos{cube.x, cube.y, cube.z - 1}]; found {
			surface -= 1
		}
	}
	pockets := map[pos]bool{}
	for x := 0; x <= 30; x++ {
		for y := 0; y <= 30; y++ {
			for z := 0; z <= 30; z++ {
				if !contains(cubes, pos{x, y, z}) &&
					contains(cubes, pos{x + 1, y, z}) && contains(cubes, pos{x - 1, y, z}) &&
					contains(cubes, pos{x, y + 1, z}) && contains(cubes, pos{x, y - 1, z}) &&
					contains(cubes, pos{x, y, z - 1}) && contains(cubes, pos{x, y, z + 1}) &&
					contains(cubes, pos{x + 1, y + 1, z}) && contains(cubes, pos{x, y + 1, z + 1}) &&
					contains(cubes, pos{x - 1, y - 1, z}) && contains(cubes, pos{x, y - 1, z - 1}) &&
					contains(cubes, pos{x + 1, y, z + 1}) && contains(cubes, pos{x - 1, y, z - 1}) {
					pockets[pos{x, y, z}] = true
				}
			}
		}
	}
	fmt.Println(surface, external)
}
