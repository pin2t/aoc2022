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

func simulate(field [][]byte) (n int) {
	n = 0
	for {
		x, y := 500, 0
		for y < 999 {
			if field[y+1][x] == '.' {
				y += 1
			} else if field[y+1][x-1] == '.' {
				y += 1
				x -= 1
			} else if field[y+1][x+1] == '.' {
				y += 1
				x += 1
			} else {
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	field := make([][]byte, 1000)
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
	field2 := make([][]byte, 1000)
	for i := 0; i < 1000; i++ {
		field2[i] = make([]byte, 1000)
		copy(field2[i], field[i])
	}
	floory := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if field2[y][x] != '.' && y > floory {
				floory = y
			}
		}
	}
	floory += 2
	for x := 0; x < 1000; x++ {
		field2[floory][x] = '#'
	}
	fmt.Println(simulate(field), simulate(field2)+1)
}
