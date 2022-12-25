package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	blueprints := []blueprint{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var b blueprint
		fmt.Sscanf(scanner.Text(),
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&b.n, &b.ore, &b.clay, &b.obsidian.ore, &b.obsidian.clay, &b.geode.ore, &b.geode.obsidian)
		blueprints = append(blueprints, b)
	}
	sum := 0
	for _, bp := range blueprints {
		sum += bp.n * maxgeodes(bp, 24, 5000000)
	}
	prod := 1
	for _, bp := range blueprints[:3] {
		prod *= maxgeodes(bp, 32, 5000000)
	}
	fmt.Println(sum, prod)
}

type blueprint struct {
	n         int
	ore, clay int
	obsidian  struct{ ore, clay int }
	geode     struct{ ore, obsidian int }
}

type items struct{ ore, clay, obsidian, geode int }

func (i items) add(ore, clay, obsidian, geode int) items {
	return items{i.ore + ore, i.clay + clay, i.obsidian + obsidian, i.geode + geode}
}

type state struct {
	minute int
	values items
	robots items
}

func (s state) val() int {
	return s.robots.ore + s.robots.clay*10 + s.robots.obsidian*100 + s.robots.geode*1000
}

func maxgeodes(bp blueprint, minutes, prune int) int {
	max := 0
	maxmin := 0
	processed := make(map[state]bool)
	queue := []state{{0, items{0, 0, 0, 0}, items{1, 0, 0, 0}}}
	for len(queue) > 0 {
		if queue[0].minute > maxmin {
			maxmin = queue[0].minute
			processed = map[state]bool{}
			if len(queue) > prune {
				sort.Slice(queue, func(i, j int) bool {
					return queue[i].val() > queue[j].val()
				})
				queue = queue[0:prune]
				continue
			}
		}
		s := queue[0]
		queue = queue[1:]
		if s.minute == minutes {
			if s.values.geode > max {
				max = s.values.geode
			}
			continue
		}
		if processed[s] {
			continue
		}
		processed[s] = true
		if s.values.ore >= bp.geode.ore && s.values.obsidian >= bp.geode.obsidian {
			queue = append(queue, state{
				s.minute + 1,
				s.values.add(s.robots.ore-bp.geode.ore, s.robots.clay, s.robots.obsidian-bp.geode.obsidian, s.robots.geode),
				s.robots.add(0, 0, 0, 1),
			})
		}
		if s.values.ore >= bp.obsidian.ore && s.values.clay >= bp.obsidian.clay {
			queue = append(queue, state{
				s.minute + 1,
				s.values.add(s.robots.ore-bp.obsidian.ore, s.robots.clay-bp.obsidian.clay, s.robots.obsidian, s.robots.geode),
				s.robots.add(0, 0, 1, 0),
			})
		}
		if s.values.ore >= bp.clay {
			queue = append(queue, state{
				s.minute + 1,
				s.values.add(s.robots.ore-bp.clay, s.robots.clay, s.robots.obsidian, s.robots.geode),
				s.robots.add(0, 1, 0, 0),
			})
		}
		if s.values.ore >= bp.ore {
			queue = append(queue, state{
				s.minute + 1,
				s.values.add(s.robots.ore-bp.ore, s.robots.clay, s.robots.obsidian, s.robots.geode),
				s.robots.add(1, 0, 0, 0),
			})
		}
		queue = append(queue, state{
			s.minute + 1,
			s.values.add(s.robots.ore, s.robots.clay, s.robots.obsidian, s.robots.geode),
			s.robots,
		})
	}
	return max
}
