package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items []int64
	op    struct {
		name string
		arg  int64
		self bool
	}
	divisor   int
	dest      map[bool]int
	inspected int
}

func (m *monkey) inspect(monkeys []monkey, relief bool) {
	d := 1
	for _, mm := range monkeys {
		d *= mm.divisor
	}
	for _, item := range m.items {
		var level int64
		if m.op.self {
			level = item * item
		} else {
			switch m.op.name {
			case "*":
				level = item * m.op.arg
			case "+":
				level = item + m.op.arg
			}
		}
		if relief {
			level = level / 3
		}
		level = level % int64(d)
		var to = &monkeys[m.dest[level%int64(m.divisor) == 0]]
		to.items = append(to.items, level)
	}
	m.inspected += len(m.items)
	m.items = []int64{}
}

func business(monkeys []monkey) int {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspected > monkeys[j].inspected
	})
	return monkeys[0].inspected * monkeys[1].inspected
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	monkeys := []monkey{}
	monkeys2 := []monkey{}
	for scanner.Scan() {
		if len(scanner.Text()) > 0 && scanner.Text()[:6] == "Monkey" {
			m := monkey{items: []int64{}, dest: map[bool]int{}}
			scanner.Scan()
			re := regexp.MustCompile("[0-9]+")
			for _, s := range re.FindAllString(scanner.Text(), -1) {
				n, _ := strconv.ParseInt(s, 0, 0)
				m.items = append(m.items, n)
			}
			scanner.Scan()
			ss := strings.Split(scanner.Text(), "=")
			var arg string
			fmt.Sscanf(ss[1], " old %s %s", &m.op.name, &arg)
			if arg == "old" {
				m.op.self = true
			} else {
				n, _ := strconv.ParseInt(arg, 0, 0)
				m.op.arg = n
			}
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "  Test: divisible by %d", &m.divisor)
			scanner.Scan()
			var to int
			fmt.Sscanf(scanner.Text(), "    If true: throw to monkey %d", &to)
			m.dest[true] = to
			scanner.Scan()
			fmt.Sscanf(scanner.Text(), "    If false: throw to monkey %d", &to)
			m.dest[false] = to
			monkeys = append(monkeys, m)
			monkeys2 = append(monkeys2, m)
		}
	}
	for i := 0; i < 20; i++ {
		for m := 0; m < len(monkeys); m++ {
			(&monkeys[m]).inspect(monkeys, true)
		}
	}
	for i := 0; i < 10000; i++ {
		for m := 0; m < len(monkeys2); m++ {
			(&monkeys2[m]).inspect(monkeys2, false)
		}
	}
	fmt.Println(business(monkeys), business(monkeys2))
}
