package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type monkey struct {
	n    float64
	wait []string
	op   string
}

func (m monkey) yell(monkeys map[string]monkey) float64 {
	if m.n > 0 {
		return m.n
	}
	yell1, yell2 := monkeys[m.wait[0]].yell(monkeys), monkeys[m.wait[1]].yell(monkeys)
	switch m.op {
	case "+":
		return yell1 + yell2
	case "-":
		return yell1 - yell2
	case "*":
		return yell1 * yell2
	case "/":
		return yell1 / yell2
	}
	panic("unknown operation" + fmt.Sprintf("%v", m))
}

func (m monkey) compare(monkeys map[string]monkey) float64 {
	return monkeys[m.wait[0]].yell(monkeys) - monkeys[m.wait[1]].yell(monkeys)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	monkeys := map[string]monkey{}
	for scanner.Scan() {
		var name string
		var n int64
		var w1, w2, op string
		if strings.ContainsAny(scanner.Text(), "0123456789") {
			fmt.Sscanf(scanner.Text(), "%s %d", &name, &n)
		} else {
			fmt.Sscanf(scanner.Text(), "%s %s %s %s", &name, &w1, &op, &w2)
		}
		monkeys[name[:4]] = monkey{float64(n), []string{w1, w2}, op}
	}
	n1 := monkeys["root"].yell(monkeys)
	low, high := int64(1), int64(1000000000000000)
	monkeys["humn"] = monkey{float64(low + (high-low)/2), []string{}, ""}
	root := monkeys["root"]
	for {
		cmp := root.compare(monkeys)
		if cmp > 0 {
			low += (high - low) / 2
		} else if cmp < 0 {
			high -= (high - low) / 2
		} else if cmp == 0 {
			fmt.Println(int64(n1), int64(monkeys["humn"].n))
			return
		}
		monkeys["humn"] = monkey{float64(low + (high-low)/2), []string{}, ""}
	}
}
