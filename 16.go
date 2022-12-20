package main

import (
	"bufio"
	"fmt"
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

type sset map[string]bool

func (set sset) contains(key string) bool {
	_, result := set[key]
	return result
}

var flow = map[string]int{}
var tunnels = sset{}

type state struct {
	valve    string
	opened   sset
	timeLeft int
	pressure int
}

func (s state) open() state {
	opened := sset(make(map[string]bool, 50))
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

func (s state) step(to string) state {
	return state{to, s.opened, s.timeLeft - 1, s.pressure}
}

type state2 struct {
	valve, evalve string
	opened        sset
	timeLeft      int
	pressure      int
}

func (s state2) open() state2 {
	opened := sset(make(map[string]bool, 50))
	for k, v := range s.opened {
		opened[k] = v
	}
	pressure := s.pressure
	if flow[s.valve] > 0 {
		opened[s.valve] = true
		pressure += flow[s.valve] * (s.timeLeft - 1)
	}
	if flow[s.evalve] > 0 {
		opened[s.evalve] = true
		pressure += flow[s.evalve] * (s.timeLeft - 1)
	}
	return state2{s.valve, s.evalve, opened, s.timeLeft - 1, pressure}
}

func (s state2) step(to, eto string) state2 {
	return state2{to, eto, s.opened, s.timeLeft - 1, s.pressure}
}

func pressure(initial state, forbidden sset) int {
	result := 0
	processed := sset(make(map[string]bool, 10000000))
	queue := make([]state, 10000000)
	queue = append(queue, initial)
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		key := fmt.Sprintf("%s,%d", s.valve, s.pressure)
		if processed.contains(key) || s.timeLeft == 0 {
			continue
		}
		processed[key] = true
		result = max(result, s.pressure)
		if !s.opened.contains(s.valve) && !forbidden.contains(s.valve) && flow[s.valve] > 0 {
			queue = append(queue, s.open())
		}
		for to, _ := range flow {
			if to != s.valve && !forbidden.contains(to) && tunnels.contains(s.valve+to) {
				queue = append(queue, s.step(to))
			}
		}
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
		flow[names[0]] = int(f)
		for i := 1; i < len(names); i++ {
			tunnels[names[0]+names[i]] = true
		}
	}
	n1 := pressure(state{"AA", sset{}, 30, 0}, sset{})
	first := sset{}
	second := sset{}
	for v, _ := range flow {
		second[v] = true
	}
	n2 := 0
	for v, _ := range second {
		n2 = max(n2, pressure(state{"AA", sset{}, 26, 0}, first)+pressure(state{"AA", sset{}, 26, 0}, second))
		delete(second, v)
		first[v] = true
	}
	fmt.Println(n1, n2)
}
