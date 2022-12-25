package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct{ x, y int }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	elves := []pos{}
	y := 0
	for scanner.Scan() {
		for x, c := range scanner.Text() {
			if c == '#' {
				elves = append(elves, pos{x, y})
			}
		}
		y++
	}
	order := "NSWE"
	for i := 0; i < 10; i++ {

	}
	fmt.Println(elves)
}
