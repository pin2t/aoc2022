package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	x := 1
	cycle := 1
	strength := 0
	scycle := 20
	crt := []string{}
	row := ""
	sprite := 0
	tick := func() {
		if cycle == scycle {
			strength += cycle * x
			scycle += 40
		}
		if (cycle-1)%40 == 0 {
			crt = append(crt, row)
			row = ""
			sprite = 0
		}
		if sprite >= x-1 && sprite <= x+1 {
			row += "#"
		} else {
			row += "."
		}
		cycle += 1
		sprite += 1
	}
	for scanner.Scan() {
		var instruction string
		var n int
		fmt.Sscanf(scanner.Text(), "%s %d", &instruction, &n)
		switch instruction {
		case "noop":
			tick()
		case "addx":
			tick()
			tick()
			x += n
		}
	}
	fmt.Println(strength)
	for _, r := range crt {
		fmt.Println(r)
	}
	fmt.Println(row)
}
