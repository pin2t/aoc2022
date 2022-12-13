package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type item interface {
	compare(with item) int
}

type number int

func (n number) compare(i item) int {
	if nitem, ok := i.(number); ok {
		if n < nitem {
			return -1
		}
		if n > nitem {
			return 1
		}
		return 0
	}
	if litem, ok := i.(list); ok {
		return n.toList().compare(litem)
	}
	panic("unexpected value in compare")
}

func (n number) toList() list {
	return list([]item{number(n)})
}

type list []item

func (l list) compare(i item) int {
	if n, ok := i.(number); ok {
		return l.compare(n.toList())
	}
	litem, _ := i.(list)
	for j, v := range l {
		if j >= len(litem) {
			return 1
		}
		vresult := v.compare(litem[j])
		if vresult != 0 {
			return vresult
		}
	}
	if len(litem) > len(l) {
		return -1
	}
	return 0
}

func (l list) divider() bool {
	if len(l) != 1 {
		return false
	}
	ll, ok := l[0].(list)
	if !ok || len(ll) != 1 {
		return false
	}
	var n number
	n, ok = ll[0].(number)
	if !ok || (n != 2 && n != 6) {
		return false
	}
	return true
}

func parse(input string, pos int) (item, int) {
	if input[pos] >= '0' && input[pos] <= '9' {
		ss := ""
		for ; pos < len(input) && input[pos] >= '0' && input[pos] <= '9'; pos++ {
			ss = ss + string(input[pos])
		}
		n, _ := strconv.ParseInt(ss, 0, 0)
		return number(int(n)), pos
	}
	if input[pos] == '[' {
		result := list([]item{})
		if input[pos+1] == ']' {
			return result, pos + 2
		}
		for pos += 1; pos < len(input); {
			var it item
			it, pos = parse(input, pos)
			result = append(result, it)
			if input[pos] == ',' {
				pos += 1
				continue
			}
			if input[pos] == ']' {
				pos += 1
				break
			}
		}
		return result, pos
	}
	panic(fmt.Sprintf("unexpected symbol %c at pos %d in %s", input[pos], pos, input))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n1, packets := 0, []item{}
	for idx := 1; scanner.Scan(); idx++ {
		p1, _ := parse(scanner.Text(), 0)
		scanner.Scan()
		p2, _ := parse(scanner.Text(), 0)
		scanner.Scan()
		if p1.compare(p2) < 0 {
			n1 += idx
		}
		packets = append(packets, p1)
		packets = append(packets, p2)
	}
	packets = append(packets, list([]item{list([]item{number(2)})}))
	packets = append(packets, list([]item{list([]item{number(6)})}))
	sort.Slice(packets, func(i, j int) bool {
		return packets[i].compare(packets[j]) < 0
	})
	n2 := 1
	for i := 0; i < len(packets); i++ {
		if ll, ok := packets[i].(list); ok && ll.divider() {
			n2 *= i + 1
		}
	}
	fmt.Println(n1, n2)
}
