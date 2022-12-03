package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func priority(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	}
	if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 27
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	p1, p2 := 0, 0
	items := make([]string, 0)
	for scanner.Scan() {
		item := scanner.Text()
		first, second := item[:len(item)/2], item[len(item)/2:]
		for _, c := range first {
			if strings.Contains(second, string(c)) {
				p1 += priority(c)
				break
			}
		}
		items = append(items, item)
		if len(items) == 3 {
			for _, r := range items[0] {
				if strings.Contains(items[1], string(r)) && strings.Contains(items[2], string(r)) {
					p2 += priority(r)
					break
				}
			}
			items = make([]string, 0)
		}
	}
	fmt.Println(p1, p2)
}
