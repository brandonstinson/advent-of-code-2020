package main

import (
	"fmt"
	"strconv"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	lines := utils.SplitFileByLine("input.txt")

	var numbers []int

	for _, line := range lines {
		num, err := strconv.Atoi(line)
		utils.ErrorCheck(err)
		numbers = append(numbers, num)
	}

	var available []int
	var wrongNumber int

	for _, num := range numbers[:25] {
		available = append(available, num)
	}

	for _, num := range numbers[25:] {
		found := findTwoThatSum(available, num)
		if !found {
			wrongNumber = num
		}
		available = shift(available, num)
	}

	sum := findNumsThatSum(numbers, wrongNumber)

	fmt.Printf("Day 09, Part 1: %d\nDay 09, Part 2: %d\n", wrongNumber, sum)
}

func shift(slice []int, n int) []int {
	new := slice[1:]
	new = append(new, n)
	return new
}

func findTwoThatSum(avail []int, sum int) bool {
	for i, num := range avail {
		diff := sum - num
		ind := findIndex(avail[i+1:], diff)
		if ind != -1 {
			return true
		}
	}
	return false
}

func findIndex(slice []int, n int) int {
	for i, v := range slice {
		if n == v {
			return i
		}
	}
	return -1
}

func findNumsThatSum(slice []int, n int) int {
	for i, v := range slice {
		iter := 1
		sum := v
		nums := []int{v}
		for sum <= n {
			sum += slice[i+iter]
			nums = append(nums, slice[i+iter])
			if sum == n {
				return sumMinMax(nums)
			}
			iter++
		}
	}
	return 0
}

func sumMinMax(nums []int) int {
	min := nums[0]
	max := 0
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min + max
}
