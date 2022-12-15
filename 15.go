package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type interval struct {
	left, right int
}

func (i interval) intersect(with interval) bool {
	return i.right >= with.left && i.left <= with.left ||
		with.right >= i.left && with.left <= i.left
}

func (i interval) merge(with interval) interval {
	if i.intersect(with) {
		return interval{min(i.left, with.left), max(i.right, with.right)}
	}
	return i
}

func (i interval) len() int {
	return i.right - i.left + 1
}

type pos struct{ x, y int }
type sensor struct {
	pos    pos
	beacon pos
}

func (s sensor) merge(coverage []interval, targety int) []interval {
	result := make([]interval, len(coverage))
	copy(result, coverage)
	distance := abs(s.beacon.x-s.pos.x) + abs(s.beacon.y-s.pos.y)
	if abs(targety-s.pos.y) <= distance {
		result = append(result, interval{
			s.pos.x - (distance - abs(targety-s.pos.y)),
			s.pos.x + (distance - abs(targety-s.pos.y)),
		})
		for {
			merged := []interval{}
			skip := map[int]bool{}
			for i := 0; i < len(result); i++ {
				if _, found := skip[i]; found {
					continue
				}
				var intersect bool = false
				for j := i + 1; j < len(result); j++ {
					if result[i].intersect(result[j]) {
						merged = append(merged, result[i].merge(result[j]))
						skip[j] = true
						intersect = true
					}
				}
				if !intersect {
					merged = append(merged, result[i])
				}
			}
			if len(merged) == len(result) {
				break
			}
			result = merged
		}
	}
	return result
}

func beacons(sensors []sensor, target int) int {
	result := map[int]bool{}
	for _, b := range sensors {
		if b.beacon.y == target {
			result[b.beacon.x] = true
		}
	}
	return len(result)
}

func coverage(sensors []sensor, target int) []interval {
	result := []interval{}
	for _, s := range sensors {
		result = s.merge(result, target)
	}
	return result
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	sensors := []sensor{}
	for input.Scan() {
		var s sensor
		fmt.Sscanf(input.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&s.pos.x, &s.pos.y, &s.beacon.x, &s.beacon.y)
		sensors = append(sensors, s)
	}
	n1 := 0
	for _, c := range coverage(sensors, 2000000) {
		n1 += c.len()
	}
	n1 -= beacons(sensors, 2000000)
	for y := 0; y <= 4000000; y++ {
		c := coverage(sensors, y)
		if len(c) > 1 {
			fmt.Println(n1, (c[0].right+1)*4000000+y)
			break
		}
	}
}
