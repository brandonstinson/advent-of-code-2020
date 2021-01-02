package main

import (
	"fmt"
	"strconv"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	strs := utils.SplitFileByLine("input.txt")

	total1 := findTwoNumbersThatTotal(strs, 2020)

	total2 := findThreeNumbersThatTotal(strs, 2020)

	fmt.Printf("Day 01, Part 1: %d\nDay 01, Part 2: %d\n", total1, total2)
}

func findTwoNumbersThatTotal(strs []string, target int) int {
	for j := 0; j < len(strs)-1; j++ {
		for k := j + 1; k < len(strs)-1; k++ {
			n1, err := strconv.Atoi(strs[j])
			utils.ErrorCheck(err)
			n2, err := strconv.Atoi(strs[k])
			utils.ErrorCheck(err)
			if n1+n2 == target {
				return n1 * n2
			}
		}
	}
	return 0
}

func findThreeNumbersThatTotal(strs []string, target int) int {
	for j := 0; j < len(strs)-1; j++ {
		for k := j + 1; k < len(strs)-1; k++ {
			n1, err := strconv.Atoi(strs[j])
			utils.ErrorCheck(err)
			n2, err := strconv.Atoi(strs[k])
			utils.ErrorCheck(err)
			if n1+n2 <= target {
				for l := k + 1; l < len(strs)-1; l++ {
					n3, err := strconv.Atoi(strs[l])
					utils.ErrorCheck(err)
					if n1+n2+n3 == target {
						return n1 * n2 * n3
					}
				}
			}

		}
	}
	return 0
}
