package main

import (
	"fmt"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	strs := utils.SplitFileByLine("input.txt")

	part1 := checkSlopes(strs, 3, 1)

	part2a := checkSlopes(strs, 1, 1)
	part2b := part1
	part2c := checkSlopes(strs, 5, 1)
	part2d := checkSlopes(strs, 7, 1)
	part2e := checkSlopes(strs, 1, 2)

	part2 := part2a * part2b * part2c * part2d * part2e

	fmt.Printf("Day 03, Part 1: %d\nDay 03, Part 2: %d\n", part1, part2)
}

func checkSlopes(slopes []string, rightInc int, downInc int) int {
	var trees int
	index := rightInc
	for i, line := range slopes[downInc:] {
		if i%downInc == 0 {
			if line[index] == []byte("#")[0] {
				trees++
			}
			index += rightInc
			if index >= len(line) {
				index -= len(line)
			}
		}
	}
	return trees
}
