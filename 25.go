package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	digits := map[rune]int{'=': -2, '-': -1, '0': 0, '1': 1, '2': 2}
	snafuDigits := []string{"0", "1", "2", "=", "-"}
	for scanner.Scan() {
		snafu := scanner.Text()
		n := 0
		for _, digit := range snafu {
			n = n*5 + digits[digit]
		}
		sum += n
	}
	snafu := ""
	for sum > 0 {
		snafu = snafuDigits[sum%5] + snafu
		if (sum % 5) > 2 {
			sum += 5
		}
		sum = sum / 5
	}
	fmt.Println(snafu)
}
