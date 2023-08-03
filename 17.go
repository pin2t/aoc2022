package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct{ x, y int64 }
type key struct {
	ijet, irock int
	depths      [7]int
}

var (
	shapes = [][]string{
		{"####"},
		{".#.", "###", ".#."},
		{"..#", "..#", "###"},
		{"#", "#", "#", "#"},
		{"##", "##"},
	}
	directions = map[byte]int{'<': -1, '>': 1}
	states     map[key]struct{ rocks, height int64 }
	stones     map[pos]bool
	jets       string
	ijet       int
	maxheight  int64
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	jets = scanner.Text()
	fmt.Println(height(2022), height(1000000000000))
}

func drop(rock []string) {
	pp := pos{maxheight + 3 + int64(len(rock)), 2}
	for {
		dy := directions[jets[ijet%len(jets)]]
		ijet++
		hitStone := false
		for x, row := range rock {
			for y, ch := range row {
				if ch == '#' {
					if stones[pos{pp.x - int64(x), int64(y) + pp.y + int64(dy)}] {
						hitStone = true
						break
					}
				}
			}
		}
		if (dy > 0 && pp.y+int64(len(rock[0])) < 7 || dy < 0 && pp.y > 0) && !hitStone {
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
	for x, row := range rock {
		for y, ch := range row {
			if ch == '#' {
				_pp := pos{pp.x - int64(x), int64(y) + pp.y}
				stones[_pp] = true
				maxheight = max(maxheight, _pp.x)
			}
		}
	}
}

func height(threshold int64) int64 {
	maxheight = int64(-1)
	skipRocks, skipHeight := int64(0), int64(0)
	states = make(map[key]struct{ rocks, height int64 })
	stones = make(map[pos]bool)
	ijet = 0
	for r := int64(0); r+skipRocks < threshold; r++ {
		drop(shapes[r%int64(len(shapes))])
		depths := [7]int{-100, -100, -100, -100, -100, -100, -100}
		for y := 0; y < 7; y++ {
			for x := 0; x > -100; x-- {
				if stones[pos{maxheight + int64(x), int64(y)}] {
					depths[y] = x
					break
				}
			}
		}
		k := key{ijet % len(jets), int(r % int64(len(shapes))), depths}
		if prev, ok := states[k]; ok {
			idiff := r - prev.rocks
			remaining := threshold - r
			hdiff := maxheight - prev.height
			cycles := remaining / idiff
			states = make(map[key]struct{ rocks, height int64 })
			skipRocks, skipHeight = cycles*idiff, cycles*hdiff
		} else {
			states[k] = struct{ rocks, height int64 }{r, maxheight}
		}
	}
	return maxheight + 1 + skipHeight
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
