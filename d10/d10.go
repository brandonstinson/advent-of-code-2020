package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	lines := utils.SplitFileByLine("input.txt")

	var nums []int

	for _, line := range lines {
		num, err := strconv.Atoi(line)
		utils.ErrorCheck(err)
		nums = append(nums, num)
	}

	// sort the list and add 0 and the max+3
	nums = append(nums, 0)
	sort.Ints(nums)
	nums = append(nums, nums[len(nums)-1]+3)

	jd := joltDiff(nums)
	jc := joltConfigs(nums, 0, map[int]int{})

	fmt.Printf("Day 10, Part 1: %d\nDay 10, Part 2: %d\n", jd, jc)
}

func joltDiff(nums []int) int {
	ones := 0
	threes := 0
	for i := 0; i < len(nums)-1; i++ {
		num1 := nums[i]
		num2 := nums[i+1]
		diff := num2 - num1
		if diff == 1 {
			ones++
		}
		if diff == 3 {
			threes++
		}
	}
	return ones * threes
}

func joltConfigs(nums []int, startInd int, cache map[int]int) int {
	if startInd == len(nums)-1 {
		return 1
	}
	if v, ok := cache[startInd]; ok {
		return v
	}
	configs := 0
	for i := startInd + 1; i < len(nums); i++ {
		if nums[i]-nums[startInd] <= 3 {
			configs += joltConfigs(nums, i, cache)
		}
	}
	cache[startInd] = configs
	return configs
}
