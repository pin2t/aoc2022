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
	if _, result := set[key]; result {
		return true
	}
	return false
}

var flow = map[string]int{}

type state struct {
	valve    string
	opened   sset
	timeLeft int
	pressure int
}

func (s state) open() state {
	opened := sset{}
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
	opened := sset{}
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tunnels := sset{}
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
	n1, n2 := 0, 0
	{
		processed := sset{}
		queue := []state{state{"AA", sset{}, 30, 0}}
		for len(queue) > 0 {
			s := queue[0]
			queue = queue[1:]
			key := fmt.Sprintf("%s,%d", s.valve, s.pressure)
			if processed.contains(key) || s.timeLeft == 0 {
				continue
			}
			processed[key] = true
			n1 = max(n1, s.pressure)
			if !s.opened.contains(s.valve) && flow[s.valve] > 0 {
				queue = append(queue, s.open())
			}
			for to, _ := range flow {
				if to != s.valve && tunnels.contains(s.valve+to) {
					queue = append(queue, s.step(to))
				}
			}
		}
	}
	{
		processed := sset{}
		queue := []state2{state2{"AA", "AA", sset{}, 26, 0}}
		for len(queue) > 0 {
			s := queue[0]
			queue = queue[1:]
			if s.timeLeft == 0 {
				continue
			}
			key := fmt.Sprintf("%s%s,%d", s.valve, s.evalve, s.pressure)
			if processed.contains(key) {
				continue
			}
			processed[key] = true
			n2 = max(n2, s.pressure)
			leftPressure := 0
			for v, f := range flow {
				if !s.opened.contains(v) {
					leftPressure += f
				}
			}
			if n2 > s.pressure+leftPressure*(s.timeLeft-1) {
				continue
			}
			if !s.opened.contains(s.valve) && flow[s.valve] > 0 ||
				!s.opened.contains(s.evalve) && flow[s.evalve] > 0 {
				queue = append(queue, s.open())
			}
			for to, _ := range flow {
				for eto, _ := range flow {
					if to != eto && to != s.valve && tunnels.contains(s.valve+to) &&
						eto != s.evalve && tunnels.contains(s.evalve+eto) {
						queue = append(queue, s.step(to, eto))
					}
				}
			}
			if len(processed)%10000 == 0 {
				fmt.Println(n2, len(processed), len(queue), s.timeLeft, s)
			}
		}
	}
	fmt.Println(n1, n2)
}
