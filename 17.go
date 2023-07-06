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
var njet = 0

type pos struct{ x, y int64 }

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

type state struct {
	iter   int64
	height int64
}

var stones = make(map[pos]bool)

func drop(r int64, hstones int64, h int64, threshold int64) (int64, int64, int64) {
	rock := shapes[r%int64(len(shapes))]
	pp := pos{hstones + 3 + int64(len(rock)), 2}
	for {
		ch := jets[njet%len(jets)]
		njet++
		dy := dx[ch]
		hitStone := false
		for x, row := range rock {
			for y, ch := range row {
				if ch == '#' {
					pp := pos{pp.x - int64(x), int64(y) + pp.y + int64(dy)}
					if stones[pp] {
						hitStone = true
						break
					}
				}
			}
		}
		if (dy > 0 && pp.y+int64(len(rock[0])) < 7 ||
			dy < 0 && pp.y > 0) && !hitStone {
			ch := jets[(njet-1)%len(jets)]
			dy := dx[ch]
			pp.y += int64(dy)
		}
		if pp.x == 0 {
			break
		}
		if int64(len(rock))-pp.x-1 == 0 {
			break
		}
		pp.x--
		overlap := false
		for x, row := range rock {
			for y, ch := range row {
				if ch == '#' {
					if stones[pos{pp.x - int64(x), int64(y) + pp.y}] {
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
	prevh := hstones
	for x, row := range rock {
		for y, ch := range row {
			if ch == '#' {
				_pp := pos{pp.x - int64(x), int64(y) + pp.y}
				stones[_pp] = true
				hstones = max(hstones, _pp.x)
			}
		}
	}
	if r > 3000 && r < 10000 {
		key := ""
		for x := hstones; x > hstones-100; x-- {
			for y := 0; y < 7; y++ {
				if stones[pos{x, int64(y)}] {
					key = key + "#"
				} else {
					key = key + "."
				}
			}
		}
		if prev, ok := seen[key]; ok {
			idiff := r - prev.iter
			remaining := threshold - r
			hdiff := hstones - prev.height
			cycles := remaining / idiff
			seen = map[string]state{}
			return r + 1 + (cycles+1)*idiff, hstones, h + cycles*hdiff + (hstones - prevh)
		} else {
			seen[key] = state{r, hstones}
		}
	}
	return r + 1, hstones, h + (hstones - prevh)
}

func height(threshold int64) int64 {
	r, hstones, h := int64(0), int64(-1), int64(0)
	for r < threshold {
		r, hstones, h = drop(r, hstones, h, threshold)
	}
	return h
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	jets = scanner.Text()
	fmt.Println(height(2022), height(1000000000000))
}
