package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type monkey struct {
	n    int64
	wait []string
	op   string
}

func (m monkey) yell(monkeys map[string]monkey) int64 {
	if m.n > 0 {
		return m.n
	}
	switch m.op {
	case "+":
		return monkeys[m.wait[0]].yell(monkeys) + monkeys[m.wait[1]].yell(monkeys)
	case "-":
		return monkeys[m.wait[0]].yell(monkeys) - monkeys[m.wait[1]].yell(monkeys)
	case "*":
		return monkeys[m.wait[0]].yell(monkeys) * monkeys[m.wait[1]].yell(monkeys)
	case "/":
		return monkeys[m.wait[0]].yell(monkeys) / monkeys[m.wait[1]].yell(monkeys)
	}
	panic("unknown operation" + fmt.Sprintf("%v", m))
}

func (m monkey) compare(monkeys map[string]monkey) int {
	if monkeys[m.wait[0]].yell(monkeys) < monkeys[m.wait[1]].yell(monkeys) {
		return -1
	}
	if monkeys[m.wait[0]].yell(monkeys) > monkeys[m.wait[1]].yell(monkeys) {
		return 1
	}
	return 0
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
		if n == 0 && op == "" {
			panic("wrong monkey " + scanner.Text() + " " + string(n) + " op " + op)
		}
		monkeys[name[:4]] = monkey{n, []string{w1, w2}, op}
	}
	n1 := monkeys["root"].yell(monkeys)
	left, right := int64(1), int64(1000000000000000)
	monkeys["humn"] = monkey{int64(left + (right-left)/2), []string{}, ""}
	root := monkeys["root"]
	for {
		cmp := root.compare(monkeys)
		if cmp > 0 {
			left += (right - left) / 2
		} else if cmp < 0 {
			right -= (right - left) / 2
		} else if cmp == 0 {
			fmt.Println(n1, monkeys["humn"].n)
			for n := left; n <= right; n++ {
				fmt.Println(n, monkeys[root.wait[0]].yell(monkeys), monkeys[root.wait[1]].yell(monkeys))
			}
			return
		}
		monkeys["humn"] = monkey{int64(left + (right-left)/2), []string{}, ""}
	}
}
