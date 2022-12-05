package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack []string

func (s *stack) move(to *stack, n int) {
	for i := 0; i < n; i++ {
		*to = append(*to, (*s)[len(*s)-1])
		*s = (*s)[:len(*s)-1]
	}
}

func tops(ss []stack) string {
	result := ""
	for _, s := range ss {
		result = result + s[len(s)-1]
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := []string{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		lines = append(lines, scanner.Text())
	}
	stacks := []stack{stack{}, stack{}, stack{}, stack{}, stack{}, stack{}, stack{}, stack{}, stack{}}
	stacks2 := []stack{stack{}, stack{}, stack{}, stack{}, stack{}, stack{}, stack{}, stack{}, stack{}}
	for i := len(lines) - 2; i >= 0; i-- {
		for j := 1; j < len(lines[i]); j += 4 {
			if lines[i][j] != ' ' {
				n := (j - 1) / 4
				stacks[n] = append(stacks[n], string(lines[i][j]))
				stacks2[n] = append(stacks2[n], string(lines[i][j]))
			}
		}
	}
	for scanner.Scan() {
		var count, from, to int
		_, err := fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &count, &from, &to)
		if err != nil {
			continue
		}
		from = from - 1
		to = to - 1
		(&stacks[from]).move(&stacks[to], count)
		tmp := stack{}
		(&stacks2[from]).move(&tmp, count)
		(&tmp).move(&stacks2[to], count)
	}
	fmt.Println(tops(stacks), tops(stacks2))
}
