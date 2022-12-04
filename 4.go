package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n1, n2 := 0, 0
	for scanner.Scan() {
		var pair1, pair2 struct{ left, right int }
		_, err := fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &pair1.left, &pair1.right, &pair2.left, &pair2.right)
		if err != nil {
			panic(err)
		}
		if pair1.left <= pair2.left && pair1.right >= pair2.right ||
			pair2.left <= pair1.left && pair2.right >= pair1.right {
			n1 += 1
		}
		if pair1.right >= pair2.left && pair2.right >= pair1.left {
			n2 += 1
		}
	}
	fmt.Println(n1, n2)
}
