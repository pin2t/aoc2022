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
		propositions := map[pos]pos{}
		for e := range elves {
			step := false
			for _, d := range []pos{{0, -1}, {0, 1}, {-1, 0}, {1, 0}, {-1, -1}, {1, 1}, {-1, 1}, {1, -1}} {
				if elves[pos{e.x + d.x, e.y + d.y}] {
					step = true
					break
				}
			}
			if step {
				moves := []struct {
					to   pos
					scan []pos
				}{
					{pos{e.x, e.y - 1}, []pos{{e.x, e.y - 1}, {e.x - 1, e.y - 1}, {e.x + 1, e.y - 1}}},
					{pos{e.x, e.y + 1}, []pos{{e.x, e.y + 1}, {e.x - 1, e.y + 1}, {e.x + 1, e.y + 1}}},
					{pos{e.x - 1, e.y}, []pos{{e.x - 1, e.y}, {e.x - 1, e.y - 1}, {e.x - 1, e.y + 1}}},
					{pos{e.x + 1, e.y}, []pos{{e.x + 1, e.y}, {e.x + 1, e.y - 1}, {e.x + 1, e.y + 1}}},
				}
				for j := 0; j < 4; j++ {
					to := (i + j) % 4
					scan := moves[to].scan
					if !elves[scan[0]] && !elves[scan[1]] && !elves[scan[2]] {
						propositions[e] = moves[to].to
						break
					}
				}
			}
		}
		nproposition := map[pos]int{}
		for _, to := range propositions {
			nproposition[to]++
		}
		moved := map[pos]bool{}
		for p, to := range propositions {
			if nproposition[to] == 1 {
				moved[to] = true
				delete(elves, p)
			}
		}
		if len(moved) == 0 {
			n2 = i + 1
			break
		}
		for e := range elves {
			moved[e] = true
		}
		elves = moved
		if i == 9 {
			top, left, bottom, right := 1000000, 1000000, -1000000, -1000000
			for e := range elves {
				top, left, bottom, right = min(top, e.x), min(left, e.y), max(bottom, e.x), max(right, e.y)
			}
			n1 = (bottom-top+1)*(right-left+1) - len(elves)
		}
	}
	fmt.Println(n1, n2)
}
