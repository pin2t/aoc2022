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
	stacks := []stack{
		stack{"L", "N", "W", "T", "D"},
		stack{"C", "P", "H"},
		stack{"W", "P", "H", "N", "D", "G", "M", "J"},
		stack{"C", "W", "S", "N", "T", "Q", "L"},
		stack{"P", "H", "C", "N"},
		stack{"T", "H", "N", "D", "M", "W", "Q", "B"},
		stack{"M", "B", "R", "J", "G", "S", "L"},
		stack{"Z", "N", "W", "G", "V", "B", "R", "T"},
		stack{"W", "G", "D", "N", "P", "L"},
	}
	stacks2 := []stack{
		stack{"L", "N", "W", "T", "D"},
		stack{"C", "P", "H"},
		stack{"W", "P", "H", "N", "D", "G", "M", "J"},
		stack{"C", "W", "S", "N", "T", "Q", "L"},
		stack{"P", "H", "C", "N"},
		stack{"T", "H", "N", "D", "M", "W", "Q", "B"},
		stack{"M", "B", "R", "J", "G", "S", "L"},
		stack{"Z", "N", "W", "G", "V", "B", "R", "T"},
		stack{"W", "G", "D", "N", "P", "L"},
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
