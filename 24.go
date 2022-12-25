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
	pos     pos
	minutes int
	stage   int // 0 - down, 1 - up, 2 - down again
}

func (s state) next(to pos) state {
	return state{to, s.minutes + 1, s.stage}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := map[pos]rune{}
	blizzards := map[pos]pos{}
	bottomright := pos{0, 0}
	for x := 0; scanner.Scan(); x++ {
		for y, c := range scanner.Text() {
			grid[pos{x, y}] = c
			bottomright = bottomright.max(pos{x, y})
			switch c {
			case '<':
				blizzards[pos{x, y}] = pos{-1, 0}
			case '>':
				blizzards[pos{x, y}] = pos{1, 0}
			case '^':
				blizzards[pos{x, y}] = pos{0, -1}
			case 'v':
				blizzards[pos{x, y}] = pos{0, 1}
			}
		}
	}
	deltas := []pos{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	queue := []state{{pos{0, 1}, 0, 0}}
	processed := map[state]bool{}
	n1 := 0
	for {
		st := queue[0]
		queue = queue[1:]
		if processed[st] {
			continue
		}
		processed[st] = true
		stopped := false
		for sp, d := range blizzards {
			if sp.x == st.pos.x {
				sy := mod((sp.y+(d.y*st.minutes)-1), bottomright.y-1) + 1
				if sy == st.pos.y {
					stopped = true
				}
			}
			if sp.y == st.pos.y {
				sx := mod((sp.x+(d.x*st.minutes)-1), bottomright.x-1) + 1
				if sx == st.pos.x {
					stopped = true
				}
			}
		}
		if stopped {
			continue
		}
		if st.pos.x == 0 && st.stage > 1 {
			st.stage++
		}
		if st.pos.x == bottomright.x {
			if st.stage == 0 {
				n1 = st.minutes
				st.stage++
			}
			if st.stage > 1 {
				fmt.Println(n1, st.minutes)
				break
			}
		}
		for _, d := range deltas {
			to := st.pos.move(pos{d.x, d.y})
			if grid[to] == '#' {
				continue
			}
			if _, ok := grid[to]; !ok {
				continue
			}
			queue = append(queue, st.next(to))
		}
		queue = append(queue, st.next(st.pos))
	}
}
