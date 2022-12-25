package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type point struct{ x, y int }

var walls = map[point]bool{}
var tiles = map[point]bool{}
var rows = 0
var size = 50

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func bounds(x int, y int) (int, int, int, int) {
	left, right := 10000, 0
	top, bottom := 10000, 0
	for p, _ := range walls {
		if p.y == y || p.x == x {
			left, right = min(left, p.x), max(right, p.x)
			top, bottom = min(top, p.y), max(bottom, p.y)
		}
	}
	for p, _ := range tiles {
		if p.y == y || p.x == x {
			left, right = min(left, p.x), max(right, p.x)
			top, bottom = min(top, p.y), max(bottom, p.y)
		}
	}
	return left, right, top, bottom
}

var state struct {
	pos    point
	dx, dy int
}

func step1() {
	np := point{state.pos.x + state.dx, state.pos.y + state.dy}
	left, right, top, bottom := bounds(state.pos.x, state.pos.y)
	if np.x < left {
		np.x = right
	}
	if np.x > right {
		np.x = left
	}
	if np.y < top {
		np.y = bottom
	}
	if np.y > bottom {
		np.y = top
	}
	if _, found := walls[np]; found {
		return
	}
	if _, found := tiles[np]; found {
		state.pos = np
		return
	}
	panic("wrong position " + fmt.Sprintf("%v", state))
}

func step2() {
	left, right, top, bottom := bounds(state.pos.x, state.pos.y)
	np := point{state.pos.x + state.dx, state.pos.y + state.dy}
	nd := point{state.dx, state.dy}
	if np.x < left {
		if state.pos.x == 0 {
			if state.pos.y < 3*size {
				np, nd = point{size, size - state.pos.y}, point{0, 1}
			} else {
				np, nd = point{4*size - 1, state.pos.y - 2*size}, point{-1, 0}
			}
		}
		if state.pos.x == size {
			np, nd = point{0, 3*size - state.pos.y}, point{0, 1}
		}
		if state.pos.x == 2*size {
			np, nd = point{state.pos.y + size, size - 1}, point{0, 1}
		}
	}
	if np.x > right {
		if state.pos.x == size-1 {
			np, nd = point{state.pos.y - size, 2*size - 1}, point{0, -1}
		}
		if state.pos.x == 3*size-1 {
			np, nd = point{4*size - (state.pos.y - size) - 1, 2 * size}, point{0, 1}
		}
		if state.pos.x == 4*size-1 {
			np, nd = point{0, state.pos.y + 2*size}, point{1, 0}
		}
	}
	if np.y < top {
		if state.pos.y == size {
			if state.pos.x < size {
				np, nd = point{size - state.pos.x, 0}, point{0, 1}
			} else {
				np, nd = point{2 * size, state.pos.x - size}, point{1, 0}
			}
		}
		if state.pos.y == 0 {
			if state.pos.x < size*3-1 {
				np, nd = point{size - (state.pos.x - 2*size), size - 1}, point{0, 1}
			} else {
				np, nd = point{0, state.pos.x - 2*size}, point{0, 1}
			}
		}
	}
	if np.y > bottom {
		if state.pos.y == 3*size-1 {
			np, nd = point{3*size - 1 - state.pos.x, 2*size - 1}, point{0, -1}
		}
		if state.pos.y == 2*size-1 {
			if state.pos.x < 2*size {
				np, nd = point{size - 1, state.pos.x + size}, point{-1, 0}
			} else {
				np, nd = point{size - (state.pos.x - 2*size), 3*size - 1}, point{0, -1}
			}
		}
		if state.pos.y == size-1 {
			np, nd = point{3*size - 1, size + (state.pos.x - 3*size)}, point{-1, 0}
		}
	}
	if _, found := walls[np]; found {
		return
	}
	if _, found := tiles[np]; found {
		state.pos = np
		state.dx = nd.x
		state.dy = nd.y
		return
	}
	panic("wrong position " + fmt.Sprintf("%v", np) + " state " + fmt.Sprintf("%v", state))
}

func password(commands string, step func()) int {
	l, _, _, _ := bounds(50, 0)
	state.pos = point{l, 0}
	state.dx, state.dy = 1, 0
	re := regexp.MustCompile("L|R|[0-9]+")
	for _, cmd := range re.FindAllString(commands, -1) {
		switch cmd {
		case "R":
			d := point{state.dx, state.dy}
			switch d {
			case point{1, 0}:
				state.dx, state.dy = 0, 1
			case point{0, 1}:
				state.dx, state.dy = -1, 0
			case point{-1, 0}:
				state.dx, state.dy = 0, -1
			case point{0, -1}:
				state.dx, state.dy = 1, 0
			}
		case "L":
			d := point{state.dx, state.dy}
			switch d {
			case point{1, 0}:
				state.dx, state.dy = 0, -1
			case point{0, 1}:
				state.dx, state.dy = 1, 0
			case point{-1, 0}:
				state.dx, state.dy = 0, 1
			case point{0, -1}:
				state.dx, state.dy = -1, 0
			}
		default:
			n, _ := strconv.ParseInt(cmd, 0, 0)
			for i := 0; i < int(n); i++ {
				step()
			}
		}
	}
	d := point{state.dx, state.dy}
	face := 0
	switch d {
	case point{1, 0}:
		face = 0
	case point{0, 1}:
		face = 1
	case point{-1, 0}:
		face = 2
	case point{0, -1}:
		face = 3
	}
	return (state.pos.y+1)*1000 + (state.pos.x+1)*4 + face
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	for scanner.Scan() && scanner.Text() != "" {
		row := scanner.Text()
		for x, c := range row {
			if c == '#' {
				walls[point{x, y}] = true
			}
			if c == '.' {
				tiles[point{x, y}] = true
			}
		}
		y++
	}
	rows = y
	if rows == 12 {
		size = 4
	}
	scanner.Scan()
	fmt.Println(password(scanner.Text(), step1), password(scanner.Text(), step2))
}
