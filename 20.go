package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mix(input []int, indices []int, multiplicator int) ([]int, []int) {
	file := make([]int, len(input))
	copy(file, input)
	ind := make([]int, len(indices))
	copy(ind, indices)
	for i := 0; i < len(file); i++ {
		from := 0
		for from < len(ind) {
			if ind[from] == i {
				break
			}
			from += 1
		}
		to := from + (file[from]*multiplicator)%(len(file)-1)
		for to <= 0 {
			to += len(file) - 1
		}
		for to >= len(file) {
			to -= len(file) - 1
		}
		n := file[from]
		nidx := ind[from]
		if to > from {
			copy(file[from:], file[from+1:to+1])
			copy(ind[from:], ind[from+1:to+1])
		} else {
			copy(file[to+1:], file[to:from])
			copy(ind[to+1:], ind[to:from])
		}
		file[to] = n
		ind[to] = nidx
	}
	return file, ind
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	file := []int{}
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 0, 0)
		file = append(file, int(n))
	}
	indices := make([]int, len(file))
	for i, _ := range file {
		indices[i] = i
	}
	file1, _ := mix(file, indices, 1)
	n1 := 0
	for j := 0; j < len(file1); j++ {
		if file1[j] == 0 {
			n1 = file1[(j+1000)%len(file1)] + file1[(j+2000)%len(file1)] + file1[(j+3000)%len(file1)]
			break
		}
	}
	file2, indices2 := mix(file, indices, 811589153)
	for r := 0; r < 9; r++ {
		file2, indices2 = mix(file2, indices2, 811589153)
	}
	for j := 0; j < len(file2); j++ {
		if file2[j] == 0 {
			fmt.Println(n1,
				file2[(j+1000)%len(file2)]*811589153+
					file2[(j+2000)%len(file2)]*811589153+
					file2[(j+3000)%len(file2)]*811589153)
			break
		}
	}
}
