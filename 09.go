package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct{ x, y int }

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func direction(from, to int) int {
	if to == from {
		return 0
	} else if to > from {
		return 1
	}
	return -1
}

type rope struct {
	knots   []pos
	visited map[pos]bool
}

func (r rope) move(dir string) {
	switch dir {
	case "R":
		r.knots[0].x += 1
	case "L":
		r.knots[0].x -= 1
	case "D":
		r.knots[0].y += 1
	case "U":
		r.knots[0].y -= 1
	}
	for i := 1; i < len(r.knots); i++ {
		if abs(r.knots[i-1].x-r.knots[i].x) > 1 || abs(r.knots[i-1].y-r.knots[i].y) > 1 {
			r.knots[i] = pos{
				r.knots[i].x + direction(r.knots[i].x, r.knots[i-1].x),
				r.knots[i].y + direction(r.knots[i].y, r.knots[i-1].y),
			}
		}
	}
	r.visited[r.knots[len(r.knots)-1]] = true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rope1 := rope{[]pos{{0, 0}, {0, 0}}, map[pos]bool{}}
	rope2 := rope{[]pos{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}, map[pos]bool{}}
	for scanner.Scan() {
		var dir string
		var n int
		fmt.Sscanf(scanner.Text(), "%s %d", &dir, &n)
		for i := 0; i < n; i++ {
			rope1.move(dir)
			rope2.move(dir)
		}
	}
	fmt.Println(len(rope1.visited), len(rope2.visited))
}
