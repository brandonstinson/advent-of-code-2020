package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	strs := utils.SplitFileByLine("input.txt")

	var correct1 int
	var correct2 int

	for _, line := range strs {
		parts := strings.Split(line, " ")
		letter := parts[1][0:1]
		password := parts[2]
		min, err := strconv.Atoi(strings.Split(parts[0], "-")[0])
		utils.ErrorCheck(err)
		max, err := strconv.Atoi(strings.Split(parts[0], "-")[1])
		utils.ErrorCheck(err)
		var count1 int
		var count2 int
		for i, l := range strings.Split(password, "") {
			if l == letter {
				count1++
			}
			if i == min-1 || i == max-1 {
				if l == letter {
					count2++
				}
			}
		}
		if count1 <= max && count1 >= min {
			correct1++
		}
		if count2 == 1 {
			correct2++
		}
	}

	fmt.Printf("Day 2 Part 1: %d\nDay 2 Part 2: %d\n", correct1, correct2)
}
