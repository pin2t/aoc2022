package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/albertorestifo/dijkstra"
)

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
			if height == 'S' {
				start = fmt.Sprintf("%d,%d", x, y)
				rr := []rune(grid[y])
				rr[x] = 'a'
				grid[y] = string(rr)
			} else if height == 'E' {
				end = fmt.Sprintf("%d,%d", x, y)
				rr := []rune(grid[y])
				rr[x] = 'a'
				grid[y] = string(rr)
			} else if height == 'a' {
				starts = append(starts, fmt.Sprintf("%d,%d", x, y))
			}
		}
	}
	for y, row := range grid {
		for x, height := range row {
			if _, found := graph[fmt.Sprintf("%d,%d", x, y)]; !found {
				vertex := map[string]int{}
				if x+1 < len(row) && int(row[x+1])-int(height) < 2 {
					vertex[fmt.Sprintf("%d,%d", x+1, y)] = 1
				}
				if x-1 >= 0 && int(row[x-1])-int(height) < 2 {
					vertex[fmt.Sprintf("%d,%d", x-1, y)] = 1
				}
				if y+1 < len(grid) && int(grid[y+1][x])-int(height) < 2 {
					vertex[fmt.Sprintf("%d,%d", x, y+1)] = 1
				}
				if y-1 >= 0 && int(grid[y-1][x])-int(height) < 2 {
					vertex[fmt.Sprintf("%d,%d", x, y-1)] = 1
				}
				graph[fmt.Sprintf("%d,%d", x, y)] = vertex
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
