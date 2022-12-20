package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	file := []int{}
	file2 := []int64{}
	for scanner.Scan() {
		n, _ := strconv.ParseInt(scanner.Text(), 0, 0)
		file = append(file, int(n))
		file2 = append(file2, n*811589153)
	}
	indices := make([]int, len(file))
	indices2 := make([]int, len(file))
	for i, _ := range file {
		indices[i] = i
		indices2[i] = i
	}
	for i := 0; i < len(file); i++ {
		from := 0
		for from < len(indices) {
			if indices[from] == i {
				break
			}
			from += 1
		}
		to := from + file[from]
		for to <= 0 {
			to += len(file) - 1
		}
		for to >= len(file) {
			to -= len(file) - 1
		}
		n := file[from]
		nidx := indices[from]
		if to > from {
			copy(file[from:], file[from+1:to+1])
			copy(indices[from:], indices[from+1:to+1])
		} else {
			copy(file[to+1:], file[to:from])
			copy(indices[to+1:], indices[to:from])
		}
		file[to] = n
		indices[to] = nidx
	}
	n1 := 0
	for j := 0; j < len(file); j++ {
		if file[j] == 0 {
			n1 = file[(j+1000)%len(file)] + file[(j+2000)%len(file)] + file[(j+3000)%len(file)]
			break
		}
	}
	for r := 0; r < 10; r++ {
		for i := 0; i < len(file2); i++ {
			from := int64(0)
			for from < int64(len(indices2)) {
				if indices2[from] == i {
					break
				}
				from += 1
			}
			to := from + file2[from]
			if to < 0 {
				to = int64(abs(to)%int64(len(file2)-1) + 1)
			}
			if to >= int64(len(file2)) {
				to = int64(to % int64(len(file2)-1))
			}
			n := file2[from]
			nidx := indices2[from]
			if to > from {
				copy(file2[from:], file2[from+1:to+1])
				copy(indices2[from:], indices2[from+1:to+1])
			} else {
				copy(file2[to+1:], file2[to:from])
				copy(indices2[to+1:], indices2[to:from])
			}
			file2[to] = n
			indices2[to] = nidx
		}
	}
	for j := 0; j < len(file2); j++ {
		if file2[j] == 0 {
			fmt.Println(n1,
				file2[(j+1000)%len(file2)]+
					file2[(j+2000)%len(file2)]+
					file2[(j+3000)%len(file2)])
		}
	}
}
