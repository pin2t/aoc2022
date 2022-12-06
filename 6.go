package main

import (
	"bufio"
	"fmt"
	"os"
)

func unique(s string) bool {
	set := map[rune]bool{}
	for _, c := range s {
		set[c] = true
	}
	return len(set) == len(s)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	chars := scanner.Text()
	n1, n2 := 0, 0
	for i := 3; i < len(chars); i++ {
		if unique(chars[i-3:i+1]) && n1 == 0 {
			n1 = i + 1
		}
		if i > 12 && unique(chars[i-13:i+1]) && n2 == 0 {
			n2 = i + 1
		}
	}
	fmt.Println(n1, n2)
}
