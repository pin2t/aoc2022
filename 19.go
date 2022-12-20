package main

import (
	"bufio"
	"fmt"
	"os"
)

type blueprint struct {
	n         int
	ore, clay int
	obsidian  struct{ ore, clay int }
	geode     struct{ ore, obsidian int }
}

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
	fmt.Println(blueprints)
}
