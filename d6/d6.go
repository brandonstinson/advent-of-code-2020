package main

import (
	"fmt"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	strs := utils.SplitFileByLine("input.txt")

	anyTotal := 0
	allTotal := 0
	var group []string

	for _, line := range strs {
		if line == "" {
			any, all := sumGroup(group)
			anyTotal += any
			allTotal += all
			group = []string{}
		} else {
			group = append(group, line)
		}
	}

	fmt.Printf("Day 6 Part 1: %d\nDay 6 Part 2: %d\n", anyTotal, allTotal)
}

func sumGroup(lines []string) (int, int) {
	count := make(map[string]int)
	for _, line := range lines {
		for _, letter := range line {
			if _, ok := count[string(letter)]; !ok {
				count[string(letter)] = 1
			} else {
				count[string(letter)]++
			}
		}
	}

	any := len(count)
	all := 0

	for _, v := range count {
		if v == len(lines) {
			all++
		}
	}
	return any, all
}
