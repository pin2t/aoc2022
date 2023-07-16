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

type state struct {
	pos         pos
	time, stage int // 0 - down, 1 - up, 2 - down again
}

func (s state) move(to pos) state {
	return state{to, s.time + 1, s.stage}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := map[pos]rune{}
	blizzards := map[pos]pos{}
	bottomright := pos{0, 0}
	directions := map[rune]pos{'<': {0, -1}, '>': {0, 1}, '^': {-1, 0}, 'v': {1, 0}}
	for x := 0; scanner.Scan(); x++ {
		for y, c := range scanner.Text() {
			grid[pos{x, y}] = c
			bottomright = bottomright.max(pos{x, y})
			if c != '#' && c != '.' {
				blizzards[pos{x, y}] = directions[c]
			}
		}
	}
	queue := make([]state, 1000000)
	queue = append(queue, state{pos{0, 1}, 0, 0})
	processed := make(map[state]bool, 1000)
	n1, maxtime, step := 0, 0, 0
	for {
		step++
		s := queue[0]
		queue = queue[1:]
		if s.time > maxtime {
			processed = make(map[state]bool, 1000)
			maxtime = s.time
		}
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
		if s.pos.x == 0 && s.stage == 1 {
			s.stage++
		}
		if s.pos.x == bottomright.x {
			if s.stage == 0 {
				if n1 == 0 {
					n1 = s.time
				}
				s.stage = 1
			}
			if s.stage > 1 {
				fmt.Println(n1, s.time)
				return
			}
		}
		for _, d := range directions {
			to := s.pos.move(d)
			if grid[to] == '#' {
				continue
			}
			if _, ok := grid[to]; !ok {
				continue
			}
			queue = append(queue, s.move(to))
		}
		queue = append(queue, s.move(s.pos))
	}
}
