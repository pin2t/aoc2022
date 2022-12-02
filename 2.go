package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scores := map[string]struct{ f, s int }{
		"AX": {1 + 3, 3 + 0}, "AY": {2 + 6, 1 + 3}, "AZ": {3 + 0, 2 + 6},
		"BX": {1 + 0, 1 + 0}, "BY": {2 + 3, 2 + 3}, "BZ": {3 + 6, 3 + 6},
		"CX": {1 + 6, 2 + 0}, "CY": {2 + 0, 3 + 3}, "CZ": {3 + 3, 1 + 6},
	}
	score, score2 := 0, 0
	for scanner.Scan() {
		s := scores[string(scanner.Text()[0])+string(scanner.Text()[2])]
		score += s.f
		score2 += s.s
	}
	fmt.Println(score, score2)
}
