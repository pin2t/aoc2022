package main

import (
	"bufio"
	"fmt"
	"os"
)

var shapes = [][]string{
	{"####"},
	{".#.", "###", ".#."},
	{"..#", "..#", "###"},
	{"#", "#", "#", "#"},
	{"##", "##"},
}
var jets string
var dx = map[byte]int{'<': -1, '>': 1}
var seen = map[string]state{}
var maxheight int = 0
var njet int = 0

type pos struct{ x, y int }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type state struct {
	iter   int64
	height int
}

func drop(r int64, threshold int64) int64 {
	stones := make(map[pos]bool)
	rock := shapes[r%int64(len(shapes))]
	if r == 0 {
		maxheight = -1
	}
	pp := pos{maxheight + 3 + len(rock), 2}
	for {
		ch := jets[njet%len(jets)]
		njet++
		dy := dx[ch]
		hitStone := false
		for x, row := range rock {
			for y, ch := range row {
				if ch == '#' {
					pp := pos{pp.x - x, y + pp.y + dy}
					if stones[pp] {
						hitStone = true
						break
					}
				}
			}
		}
		if (dy > 0 && pp.y+len(rock[0]) < 7 ||
			dy < 0 && pp.y > 0) && !hitStone {
			ch := jets[(njet-1)%len(jets)]
			dy := dx[ch]
			pp.y += dy
		}
		if pp.x == 0 {
			break
		}
		if len(rock)-pp.x-1 == 0 {
			break
		}
		pp.x--
		overlap := false
		for x, row := range rock {
			for y, ch := range row {
				if ch == '#' {
					if stones[pos{pp.x - x, y + pp.y}] {
						overlap = true
						break
					}
				}
			}
		}
		pp.x++
		if overlap {
			break
		}
		pp.x--
	}
	for x, row := range rock {
		for y, ch := range row {
			if ch == '#' {
				_pp := pos{pp.x - x, y + pp.y}
				stones[_pp] = true
				maxheight = max(maxheight, _pp.x)
			}
		}
	}
	if r > 3000 {
		key := ""
		for x := maxheight; x > maxheight-100; x-- {
			for y := 0; y < 7; y++ {
				if stones[pos{x, y}] {
					key = key + "#"
				} else {
					key = key + "."
				}
			}
		}
		if prev, ok := seen[key]; ok {
			idiff := r - prev.iter
			remaining := threshold - r
			if remaining%idiff == 0 {
				height := int64(maxheight + 1)
				hdiff := height - int64(prev.height)
				remcycles := (remaining / idiff) + 1
				return int64(prev.height) + int64(remcycles)*hdiff - 1
			}
		}
		seen[key] = state{r, maxheight + 1}
	}
	return -1
}

func height(threshold int64) int64 {
	for r := int64(0); r < threshold; r++ {
		if x := drop(r, threshold); x > 0 {
			return x
		}
	}
	return int64(maxheight + 1)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	jets = scanner.Text()
	fmt.Println(height(2022), height(1000000000000))
}
