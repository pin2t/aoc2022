package main

import (
	"bufio"
	"fmt"
	"os"
)

func mod(x, m int) int {
	return (x%m + m) % m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type pos struct {
	x, y int
}

func (p pos) max(b pos) pos {
	return pos{max(p.x, b.x), max(p.y, b.y)}
}

func (p pos) move(delta pos) pos {
	return pos{p.x + delta.x, p.y + delta.y}
}

var walls = map[pos]bool{}
var bottomright = pos{0, 0}

const (
	DOWN  = 0
	UP    = 1
	DOWN2 = 2
)

type state struct {
	pos         pos
	time, stage int
}

func (s state) move(to pos) state {
	var next = s.stage
	if s.pos.x == 0 && s.stage == UP || s.pos.x == bottomright.x && s.stage == DOWN {
		next++
	}
	return state{to, s.time + 1, next}
}

func (s state) canMove(direction pos) bool {
	to := s.pos.move(direction)
	if walls[to] || to.x < 0 || to.x > bottomright.x || to.y < 0 || to.y > bottomright.y {
		return false
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	blizzards := map[pos]pos{}
	directions := map[rune]pos{'<': {0, -1}, '>': {0, 1}, '^': {-1, 0}, 'v': {1, 0}}
	for x := 0; scanner.Scan(); x++ {
		for y, c := range scanner.Text() {
			if c == '#' {
				walls[pos{x, y}] = true
			} else if d, ok := directions[c]; ok {
				blizzards[pos{x, y}] = d
			}
			bottomright = bottomright.max(pos{x, y})
		}
	}
	queue := make([]state, 1000000)
	queue = append(queue, state{pos{0, 1}, 0, 0})
	processed := make(map[state]bool, 10000000)
	n1 := 0
	for {
		s := queue[0]
		queue = queue[1:]
		if processed[s] {
			continue
		}
		processed[s] = true
		stopped := false
		for sp, d := range blizzards {
			if sp.x == s.pos.x && d.y != 0 {
				sy := mod(sp.y+(d.y*s.time)-1, bottomright.y-1) + 1
				if sy == s.pos.y {
					stopped = true
					break
				}
			} else if sp.y == s.pos.y && d.x != 0 {
				sx := mod(sp.x+(d.x*s.time)-1, bottomright.x-1) + 1
				if sx == s.pos.x {
					stopped = true
					break
				}
			}
		}
		if stopped {
			continue
		}
		if s.pos.x == bottomright.x {
			if n1 == 0 {
				n1 = s.time
			}
			if s.stage == DOWN2 {
				fmt.Println(n1, s.time)
				return
			}
		}
		for _, d := range directions {
			if s.canMove(d) {
				queue = append(queue, s.move(s.pos.move(d)))
			}
		}
		queue = append(queue, s.move(s.pos))
	}
}
