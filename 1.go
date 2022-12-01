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
	calories := make([]int, 1)
	for scanner.Scan() {
		if scanner.Text() == "" {
			calories = append(calories, 0)
		} else {
			c, _ := strconv.ParseInt(scanner.Text(), 0, 0)
			calories[len(calories)-1] = calories[len(calories)-1] + int(c)
		}
	}
	sort.Ints(calories)
	fmt.Println(calories[len(calories)-1], calories[len(calories)-1]+calories[len(calories)-2]+calories[len(calories)-3])
}
