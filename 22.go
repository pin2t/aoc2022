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

func bounds(y int) (int, int) {
	left, right := 10000, 0
	for p, _ := range walls {
		if p.y == y {
			left, right = min(left, p.x), max(right, p.x)
		}
	}
	for p, _ := range tiles {
		if p.y == y {
			left, right = min(left, p.x), max(right, p.x)
		}
	}
	return left, right
}

func boundsy(x int) (int, int) {
	top, bottom := 10000, 0
	for p, _ := range walls {
		if p.x == x {
			top, bottom = min(top, p.y), max(bottom, p.y)
		}
	}
	for p, _ := range tiles {
		if p.x == x {
			top, bottom = min(top, p.y), max(bottom, p.y)
		}
	}
	return top, bottom
}

var player struct {
	pos    point
	dx, dy int
}

func step() {
	np := point{player.pos.x + player.dx, player.pos.y + player.dy}
	left, right := bounds(player.pos.y)
	top, bottom := boundsy(player.pos.x)
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
		player.pos = np
		return
	}
	panic("wrong position " + fmt.Sprintf("%v", player))
}

func step2() {
	np := point{player.pos.x + player.dx, player.pos.y + player.dy}
	left, right := bounds(player.pos.y)
	top, bottom := boundsy(player.pos.x)
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
		player.pos = np
		return
	}
	panic("wrong position " + fmt.Sprintf("%v", player))
}

func left() {
	direction := point{player.dx, player.dy}
	switch direction {
	case point{1, 0}:
		player.dx, player.dy = 0, -1
	case point{0, 1}:
		player.dx, player.dy = 1, 0
	case point{-1, 0}:
		player.dx, player.dy = 0, 1
	case point{0, -1}:
		player.dx, player.dy = -1, 0
	}
}

func right() {
	direction := point{player.dx, player.dy}
	switch direction {
	case point{1, 0}:
		player.dx, player.dy = 0, 1
	case point{0, 1}:
		player.dx, player.dy = -1, 0
	case point{-1, 0}:
		player.dx, player.dy = 0, -1
	case point{0, -1}:
		player.dx, player.dy = 1, 0
	}
}

func face() int {
	direction := point{player.dx, player.dy}
	switch direction {
	case point{1, 0}:
		return 0
	case point{0, 1}:
		return 1
	case point{-1, 0}:
		return 2
	case point{0, -1}:
		return 3
	}
	return -1
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
	l, _ := bounds(0)
	player.pos = point{l, 0}
	player.dx, player.dy = 1, 0
	scanner.Scan()
	re := regexp.MustCompile("L|R|[0-9]+")
	for _, cmd := range re.FindAllString(scanner.Text(), -1) {
		switch cmd {
		case "R":
			right()
		case "L":
			left()
		default:
			n, _ := strconv.ParseInt(cmd, 0, 0)
			for i := 0; i < int(n); i++ {
				step()
			}
		}
	}
	password1 := (player.pos.y+1)*1000 + (player.pos.x+1)*4 + face()
	player.pos = point{l, 0}
	player.dx, player.dy = 1, 0
	scanner.Scan()
	for _, cmd := range re.FindAllString(scanner.Text(), -1) {
		switch cmd {
		case "R":
			right()
		case "L":
			left()
		default:
			n, _ := strconv.ParseInt(cmd, 0, 0)
			for i := 0; i < int(n); i++ {
				step2()
			}
		}
	}
	fmt.Println(password1, (player.pos.y+1)*1000+(player.pos.x+1)*4+face())
}
