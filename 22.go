package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type point struct{ x, y int }

var (
	walls, tiles          = map[point]bool{}, map[point]bool{}
	up, down, left, right = point{0, -1}, point{0, 1}, point{-1, 0}, point{1, 0}
	pos                   = point{1, 0}
	direction             = right
)

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

func bounds(po point) (int, int, int, int) {
	l, r := 10000, 0
	t, b := 10000, 0
	for p, _ := range walls {
		if p.y == po.y {
			l, r = min(l, p.x), max(r, p.x)
		}
		if p.x == po.x {
			t, b = min(t, p.y), max(b, p.y)
		}
	}
	for p, _ := range tiles {
		if p.y == po.y {
			l, r = min(l, p.x), max(r, p.x)
		}
		if p.x == po.x {
			t, b = min(t, p.y), max(b, p.y)
		}
	}
	return l, r, t, b
}

func step1() {
	l, r, t, b := bounds(pos)
	to := point{pos.x + direction.x, pos.y + direction.y}
	if to.x < l {
		to.x = r
	} else if to.x > r {
		to.x = l
	}
	if to.y < t {
		to.y = b
	} else if to.y > b {
		to.y = t
	}
	if _, found := walls[to]; found {
		return
	}
	if _, found := tiles[to]; found {
		pos = to
		return
	}
	panic("outside position " + fmt.Sprintf("%v", pos) + " direction " + fmt.Sprintf("%v", direction))
}

func step2() {
	np := point{pos.x + direction.x, pos.y + direction.y}
	nd := direction
	switch nd {
	case down:
		if pos.y == 49 {
			if np.x >= 100 && np.x < 150 {
				np, nd = point{99, np.x - 50}, left
			}
		} else if pos.y == 149 {
			if np.x >= 50 && np.x < 100 {
				np, nd = point{49, np.x + 100}, left
			}
		} else if np.y >= 200 && np.x >= 0 && np.x < 50 {
			np, nd = point{np.x + 100, 0}, down
		}
	case up:
		switch pos.y {
		case 0:
			if np.x >= 50 && np.x < 100 {
				np, nd = point{0, np.x + 100}, right
			} else if np.x >= 100 && np.x < 150 {
				np, nd = point{np.x - 100, 199}, up
			}
		case 100:
			if np.x >= 0 && np.x < 50 {
				np, nd = point{50, np.x + 50}, right
			}
		}
	case right:
		if np.x >= 150 && np.y >= 0 && np.y < 50 {
			np, nd = point{99, 149 - np.y}, left
		} else if pos.x == 49 {
			if np.y >= 150 && np.y < 200 {
				np, nd = point{np.y - 100, 149}, up
			}
		} else if pos.x == 99 {
			if np.y >= 50 && np.y < 100 {
				np, nd = point{np.y + 50, 49}, up
			} else if np.y >= 100 && np.y < 150 {
				np, nd = point{149, 149 - np.y}, left
			}
		}
	case left:
		if np.x < 0 && np.y >= 150 && np.y < 200 {
			np, nd = point{np.y - 100, 0}, down
		} else if np.x == 49 && np.y >= 50 && np.y < 100 {
			np, nd = point{np.y - 50, 100}, down
		} else if np.x == 49 && np.y >= 0 && np.y < 50 {
			np, nd = point{0, 149 - np.y}, right
		} else if np.x < 0 && np.y >= 100 && np.y < 150 {
			np, nd = point{50, 149 - np.y}, right
		}
	}
	if _, found := walls[np]; found {
		return
	}
	if _, found := tiles[np]; found {
		pos = np
		direction = nd
		return
	}
	panic("outside position " + fmt.Sprintf("%v", np) + " previous position " + fmt.Sprintf("%v", pos))
}

func step2_4() {
	l, r, t, b := bounds(pos)
	np := point{pos.x + direction.x, pos.y + direction.y}
	nd := direction
	if np.x < l {
		if pos.x == 0 {
			if pos.y < 12 {
				np, nd = point{4, 4 - pos.y}, point{0, 1}
			} else {
				np, nd = point{16 - 1, pos.y - 8}, point{-1, 0}
			}
		}
		if pos.x == 4 {
			np, nd = point{0, 12 - pos.y}, point{0, 1}
		}
		if pos.x == 8 {
			np, nd = point{pos.y + 4, 4 - 1}, point{0, 1}
		}
	}
	if np.x > r {
		if pos.x == 3 {
			np, nd = point{pos.y - 4, 7}, point{0, -1}
		}
		if pos.x == 11 {
			np, nd = point{16 - (pos.y - 4) - 1, 8}, point{0, 1}
		}
		if pos.x == 15 {
			np, nd = point{0, pos.y + 8}, point{1, 0}
		}
	}
	if np.y < t {
		if pos.y == 4 {
			if pos.x < 4 {
				np, nd = point{4 - pos.x, 0}, point{0, 1}
			} else {
				np, nd = point{8, pos.x - 4}, point{1, 0}
			}
		}
		if pos.y == 0 {
			if pos.x < 11 {
				np, nd = point{4 - (pos.x - 8), 3}, point{0, 1}
			} else {
				np, nd = point{0, pos.x - 8}, point{0, 1}
			}
		}
	}
	if np.y > b {
		if pos.y == 11 {
			np, nd = point{11 - pos.x, 7}, point{0, -1}
		}
		if pos.y == 7 {
			if pos.x < 8 {
				np, nd = point{3, pos.x + 4}, point{-1, 0}
			} else {
				np, nd = point{4 - (pos.x - 8), 11}, point{0, -1}
			}
		}
		if pos.y == 3 {
			np, nd = point{11, 4 + (pos.x - 12)}, point{-1, 0}
		}
	}
	if _, found := walls[np]; found {
		return
	}
	if _, found := tiles[np]; found {
		pos = np
		direction = nd
		return
	}
	panic("wrong position " + fmt.Sprintf("%v", np) + " previous position " + fmt.Sprintf("%v", pos))
}

func password(commands string, step func()) int {
	var (
		leftTurns  = map[point]point{{1, 0}: {0, 1}, {0, 1}: {-1, 0}, {-1, 0}: {0, -1}, {0, -1}: {1, 0}}
		rightTurns = map[point]point{{1, 0}: {0, -1}, {0, 1}: {1, 0}, {-1, 0}: {0, 1}, {0, -1}: {-1, 0}}
		faces      = map[point]int{{1, 0}: 0, {0, 1}: 1, {-1, 0}: 2, {0, -1}: 3}
	)
	l, _, _, _ := bounds(point{0, 0})
	pos = point{l, 0}
	direction = right
	re := regexp.MustCompile("L|R|[0-9]+")
	for _, cmd := range re.FindAllString(commands, -1) {
		switch cmd {
		case "R":
			direction = leftTurns[direction]
		case "L":
			direction = rightTurns[direction]
		default:
			n, _ := strconv.ParseInt(cmd, 0, 0)
			for i := 0; i < int(n); i++ {
				step()
			}
		}
	}
	return (pos.y+1)*1000 + (pos.x+1)*4 + faces[direction]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	y := 0
	for scanner.Scan() && scanner.Text() != "" {
		for x, c := range scanner.Text() {
			if c == '#' {
				walls[point{x, y}] = true
			}
			if c == '.' {
				tiles[point{x, y}] = true
			}
		}
		y++
	}
	var pass2 int
	scanner.Scan()
	pass1 := password(scanner.Text(), step1)
	if y == 12 { // test example
		pass2 = password(scanner.Text(), step2_4)
	} else {
		pass2 = password(scanner.Text(), step2)
	}
	fmt.Println(pass1, pass2)
}
