package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/albertorestifo/dijkstra"
)

func replace(s string, i int, c rune) string {
	rr := []rune(s)
	rr[i] = c
	return string(rr)
}

func pos(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := []string{}
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	graph := dijkstra.Graph{}
	var start, end string
	var starts []string
	for y, row := range grid {
		for x, height := range row {
			switch height {
			case 'S':
				start = pos(x, y)
				grid[y] = replace(grid[y], x, 'a')
			case 'E':
				end = pos(x, y)
				grid[y] = replace(grid[y], x, 'z')
			case 'a':
				starts = append(starts, pos(x, y))
			}
		}
	}
	for y, row := range grid {
		for x, height := range row {
			if _, found := graph[pos(x, y)]; !found {
				vertex := map[string]int{}
				edge := func(x, y int) {
					if x >= 0 && x < len(row) &&
						y >= 0 && y < len(grid) &&
						int(grid[y][x])-int(height) < 2 {
						vertex[pos(x, y)] = 1
					}
				}
				edge(x+1, y)
				edge(x-1, y)
				edge(x, y+1)
				edge(x, y-1)
				graph[pos(x, y)] = vertex
			}
		}
	}
	_, n1, _ := graph.Path(start, end)
	n2 := 100000000000
	for _, s := range starts {
		_, cost, _ := graph.Path(s, end)
		if cost > 0 && cost < n2 {
			n2 = cost
		}
	}
	fmt.Println(n1, n2)
}
