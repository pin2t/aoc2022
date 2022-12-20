package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type vset map[int]bool

func (set vset) contains(key int) bool {
	_, result := set[key]
	return result
}

var flow = map[int]int{}
var tunnels = vset{}

type state struct {
	valve    int
	opened   vset
	timeLeft int
	pressure int
}

func (s state) open() state {
	opened := vset{}
	for k, v := range s.opened {
		opened[k] = v
	}
	opened[s.valve] = true
	return state{
		s.valve,
		opened,
		s.timeLeft - 1,
		s.pressure + flow[s.valve]*(s.timeLeft-1),
	}
}

func (s state) step(to int) state {
	return state{to, s.opened, s.timeLeft - 1, s.pressure}
}

func pressure(initial state, forbidden vset) int {
	result := 0
	processed := vset{}
	queue := []state{}
	queue = append(queue, initial)
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		key := s.pressure*26*26 + s.valve
		if processed.contains(key) || s.timeLeft == 0 {
			continue
		}
		processed[key] = true
		result = max(result, s.pressure)
		if !s.opened.contains(s.valve) && !forbidden.contains(s.valve) && flow[s.valve] > 0 {
			queue = append(queue, s.open())
		}
		for to, _ := range flow {
			if to != s.valve && !forbidden.contains(to) && tunnels.contains(s.valve*26*26+to) {
				queue = append(queue, s.step(to))
			}
		}
	}
	return result
}

func key(valve string) int {
	return int(valve[0]-'A')*26 + int(valve[1]-'A')
}

func key2(v1, v2 string) int {
	return key(v1)*26*26 + key(v2)
}

func powerset(set []int) []vset {
	size := int(math.Pow(2, float64(len(set))))
	result := []vset{}
	var index int
	for index < size {
		subSet := vset{}
		for j, elem := range set {
			if index&(1<<uint(j)) > 0 {
				subSet[elem] = true
			}
		}
		result = append(result, subSet)
		index++
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	reFlow := regexp.MustCompile("[0-9]+")
	reNames := regexp.MustCompile("[A-Z]{2}")
	for scanner.Scan() {
		names := reNames.FindAllString(scanner.Text(), -1)
		f, _ := strconv.ParseInt(reFlow.FindString(scanner.Text()), 0, 0)
		flow[key(names[0])] = int(f)
		for i := 1; i < len(names); i++ {
			tunnels[key2(names[0], names[i])] = true
		}
	}
	valves := []int{}
	nonempty := []int{}
	for v, f := range flow {
		valves = append(valves, v)
		if f != 0 {
			nonempty = append(nonempty, v)
		}
	}
	n1 := pressure(state{0, vset{}, 30, 0}, vset{})
	n2 := 0
	for _, first := range powerset(nonempty) {
		second := vset{}
		for v, f := range flow {
			if !first.contains(v) && f > 0 {
				second[v] = true
			}
		}
		n2 = max(n2, pressure(state{0, vset{}, 26, 0}, first)+pressure(state{0, vset{}, 26, 0}, second))
	}
	fmt.Println(n1, n2)
}
