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
	rucksacks := make([]string, 0)
	for scanner.Scan() {
		rucksack := scanner.Text()
		first, second := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
		for _, c := range first {
			if strings.ContainsRune(second, c) {
				p1 += priority(c)
				break
			}
		}
		rucksacks = append(rucksacks, rucksack)
		if len(rucksacks) == 3 {
			for _, r := range rucksacks[0] {
				if strings.ContainsRune(rucksacks[1], r) && strings.ContainsRune(rucksacks[2], r) {
					p2 += priority(r)
					break
				}
			}
			rucksacks = make([]string, 0)
		}
	}
	fmt.Println(p1, p2)
}
