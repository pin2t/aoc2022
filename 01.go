package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	calories := make([]int, 1000)
	for scanner.Scan() {
		if scanner.Text() == "" {
			calories = append(calories, 0)
		} else {
			c, _ := strconv.ParseInt(scanner.Text(), 0, 0)
			calories[len(calories)-1] = calories[len(calories)-1] + int(c)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	fmt.Println(calories[0], calories[0]+calories[1]+calories[2])
}
