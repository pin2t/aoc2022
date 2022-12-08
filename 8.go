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
	for i := 1; i < len(grid[0])-1; i++ {
		for j := 1; j < len(grid)-1; j++ {
			lower := true
			for k := 0; k < i && lower; k++ {
				if grid[j][k] >= grid[j][i] {
					lower = false
				}
			}
			if lower {
				n += 1
				continue
			}
			lower = true
			for k := i + 1; k < len(grid[0]) && lower; k++ {
				if grid[j][k] >= grid[j][i] {
					lower = false
				}
			}
			if lower {
				n += 1
				continue
			}
			lower = true
			for k := 0; k < j && lower; k++ {
				if grid[k][i] >= grid[j][i] {
					lower = false
				}
			}
			if lower {
				n += 1
				continue
			}
			lower = true
			for k := j + 1; k < len(grid) && lower; k++ {
				if grid[k][i] >= grid[j][i] {
					lower = false
				}
			}
			if lower {
				n += 1
				continue
			}
		}
	}
	highest := 0
	for i := 1; i < len(grid[0])-1; i++ {
		for j := 1; j < len(grid)-1; j++ {
			score, k := 1, 0
			for k = i - 1; k > 0 && grid[j][k] < grid[j][i]; k-- {
			}
			score *= (i - k)
			for k = i + 1; k < len(grid[j])-1 && grid[j][k] < grid[j][i]; k++ {
			}
			score *= (k - i)
			for k = j - 1; k > 0 && grid[k][i] < grid[j][i]; k-- {
			}
			score *= (j - k)
			for k = j + 1; k < len(grid)-1 && grid[k][i] < grid[j][i]; k++ {
			}
			score *= (k - j)
			if score > highest {
				highest = score
			}
		}
	}
	fmt.Println(n, highest)
}
