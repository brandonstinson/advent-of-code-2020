package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	lines := utils.SplitFileByLine("input.txt")

	acc1, _ := accumulator(lines)

	acc2 := bruteForce(lines)

	fmt.Printf("Day 08, Part 1: %d\nDay 08, Part 2: %d\n", acc1, acc2)
}

func accumulator(lines []string) (int, int) {
	seen := make(map[int]int)
	acc := 0
	index := 0

	for seen[index] < 1 && index < len(lines) {
		parts := strings.Split(lines[index], " ")
		cmd := parts[0]
		inc, err := strconv.Atoi(parts[1])
		utils.ErrorCheck(err)
		seen[index] = 1
		switch cmd {
		case "nop":
			index++
		case "acc":
			acc += inc
			index++
		case "jmp":
			index += inc
		}
	}
	return acc, index
}

func bruteForce(lines []string) int {
	for i := 0; i < len(lines)-1; i++ {
		cmd := strings.Split(lines[i], " ")[0]
		if cmd != "acc" {
			var newLines []string
			for ind, l := range lines {
				if i == ind {
					if cmd == "nop" {
						newLines = append(newLines, "jmp"+l[3:])
					} else {
						newLines = append(newLines, "nop"+l[3:])
					}
				} else {
					newLines = append(newLines, l)
				}
			}
			acc, lastIndex := accumulator(newLines)
			if lastIndex == len(lines) {
				return acc
			}
		}
	}
	return 0
}
