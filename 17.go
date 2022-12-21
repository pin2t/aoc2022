package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	jets := []int{}
	scanner.Scan()
	for _, r := range scanner.Text() {
		if r == '<' {
			jets = append(jets, -1)
		} else {
			jets = append(jets, 1)
		}
	}
	fmt.Println(jets)
}
