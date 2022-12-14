package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func direction(from, to int) int {
	if from == to {
		return 0
	}
	if from < to {
		return 1
	}
	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	field := make([][]byte, 1000)
	simulate := func() (n int) {
		n = 0
		for {
			x, y := 500, 0
			for y < 999 {
				fall := func(dx int) bool {
					if field[y+1][x+dx] == '.' {
						y += 1
						x += dx
						return true
					}
					return false
				}
				if !fall(0) && !fall(-1) && !fall(1) {
					field[y][x] = 'o'
					break
				}
			}
			if y == 999 || x == 500 && y == 0 {
				break
			}
			n += 1
		}
		return n
	}
	for i := 0; i < 1000; i++ {
		field[i] = make([]byte, 1000)
		for j := 0; j < 1000; j++ {
			field[i][j] = '.'
		}
	}
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), " -> ")
		var x, y int
		fmt.Sscanf(pairs[0], "%d,%d", &x, &y)
		for i := 1; i < len(pairs); i++ {
			var tox, toy int
			fmt.Sscanf(pairs[i], "%d,%d", &tox, &toy)
			dx, dy := direction(x, tox), direction(y, toy)
			for ; x != tox || y != toy; x, y = x+dx, y+dy {
				field[y][x] = '#'
			}
			field[y][x] = '#'
		}
	}
	n1 := simulate()
	floory := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if field[y][x] != '.' && y > floory {
				floory = y
			}
		}
	}
	floory += 2
	for x := 0; x < 1000; x++ {
		field[floory][x] = '#'
	}
	fmt.Println(n1, n1+simulate()+1)
}
