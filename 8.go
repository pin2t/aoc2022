package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := []string{}
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	n := len(grid[0])*2 + len(grid)*2 - 4 // tress on edges
	highest := 0
	for i := 1; i < len(grid[0])-1; i++ {
		for j := 1; j < len(grid)-1; j++ {
			score, k, visible := 1, 0, false
			for k = i - 1; k > 0 && grid[j][k] < grid[j][i]; k-- {
			}
			visible = visible || k == 0 && grid[j][k] < grid[j][i]
			score *= (i - k)
			for k = i + 1; k < len(grid[j])-1 && grid[j][k] < grid[j][i]; k++ {
			}
			visible = visible || k == len(grid[j])-1 && grid[j][k] < grid[j][i]
			score *= (k - i)
			for k = j - 1; k > 0 && grid[k][i] < grid[j][i]; k-- {
			}
			visible = visible || k == 0 && grid[k][i] < grid[j][i]
			score *= (j - k)
			for k = j + 1; k < len(grid)-1 && grid[k][i] < grid[j][i]; k++ {
			}
			visible = visible || k == len(grid)-1 && grid[k][i] < grid[j][i]
			score *= (k - j)
			if score > highest {
				highest = score
			}
			if visible {
				n += 1
			}
		}
	}
	fmt.Println(n, highest)
}
