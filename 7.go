package main

import (
	"bufio"
	"fmt"
	"os"
)

type file struct {
	name string
	size int
}

type dir struct {
	name   string
	parent *dir
	files  []file
	dirs   []*dir
}

func (d *dir) size() int {
	result := 0
	for _, file := range d.files {
		result = result + file.size
	}
	for _, dir := range d.dirs {
		result = result + dir.size()
	}
	return result
}

func (d *dir) total() int {
	result := 0
	if d.size() <= 100000 {
		result = result + d.size()
	}
	for _, dir := range d.dirs {
		result = result + dir.total()
	}
	return result
}

func (d *dir) smallest(threshold int) int {
	result := 10000000000
	for _, dir := range d.dirs {
		size := dir.smallest(threshold)
		if size >= threshold && size < result {
			result = size
		}
	}
	size := d.size()
	if size >= threshold && size < result {
		result = size
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	root := dir{name: "/", files: []file{}, dirs: []*dir{}}
	current := &root
	for scanner.Scan() {
		line := scanner.Text()
		if line[:4] == "$ ls" {
			// skip
		} else if line[:5] == "$ cd " {
			switch line[5:] {
			case "/":
				current = &root
			case "..":
				current = current.parent
			default:
				for _, dir := range current.dirs {
					if dir.name == line[5:] {
						current = dir
					}
				}
			}
		} else if line[:3] == "dir" {
			current.dirs = append(current.dirs, &dir{name: line[4:], files: []file{}, dirs: []*dir{}, parent: current})
		} else {
			var file file
			fmt.Sscanf(line, "%d %s", &file.size, &file.name)
			current.files = append(current.files, file)
		}
	}
	fmt.Println((&root).total(), (&root).smallest(30000000-(70000000-(&root).size())))
}
