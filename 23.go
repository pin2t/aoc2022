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

type pos struct{ x, y int }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	elves := map[pos]bool{}
	for y := 0; scanner.Scan(); y++ {
		for x, c := range scanner.Text() {
			if c == '#' {
				elves[pos{x, y}] = true
			}
		}
	}
	n1, n2 := 0, 0
	for i := 0; true; i++ {
		moves := map[pos]pos{}
		for e := range elves {
			var north, south, west, east bool
			for _, d := range []pos{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
				if elves[pos{e.x + d.x, e.y + d.y}] {
					if d.y == -1 {
						north = true
					}
					if d.y == 1 {
						south = true
					}
					if d.x == -1 {
						west = true
					}
					if d.x == 1 {
						east = true
					}
				}
			}
			if !north && !south && !west && !east {
				continue
			}
			can := []struct {
				can bool
				pos pos
			}{
				{north, pos{e.y, e.y - 1}},
				{south, pos{e.y, e.y + 1}},
				{west, pos{e.x - 1, e.y}},
				{east, pos{e.x + 1, e.y}},
			}
			for j := 0; j < 4; j++ {
				to := (i + j) % 4
				if can[to].can {
					moves[e] = can[to].pos
					break
				}
			}
		}
		wills := map[pos]int{}
		for _, to := range moves {
			wills[to]++
		}
		moved := map[pos]bool{}
		for p, to := range moves {
			if wills[to] == 1 {
				moved[to] = true
				delete(elves, p)
			}
		}
		if len(moved) == 0 {
			n2 = i
			break
		}
		for e, _ := range elves {
			moved[e] = true
		}
		elves = moved
		if i == 9 {
			var topleft, bottomright = pos{1000000, 1000000}, pos{-1000000, -1000000}
			for e, _ := range elves {
				topleft.x, topleft.y, bottomright.x, bottomright.y = min(topleft.x, e.x), min(topleft.y, e.y), max(bottomright.x, e.x), max(bottomright.y, e.y)
			}
			var field int
			for x := topleft.x; x <= bottomright.x; x++ {
				for y := topleft.y; y <= bottomright.y; y++ {
					if _, found := elves[pos{x, y}]; found {
						field++
					}
				}
			}
			n1 = field
		}
	}
	fmt.Println(n1, n2)
}
