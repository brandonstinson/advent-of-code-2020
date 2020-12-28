package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/brandonstinson/advent-of-code/utils"
)

type rule struct {
	min1 int
	max1 int
	min2 int
	max2 int
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	utils.ErrorCheck(err)
	parts := strings.Split(string(data), "\n\n")

	rules := strings.Split(parts[0], "\n")
	myTicket := strings.Split(parts[1], "\n")[1]
	nearby := strings.Split(parts[2], "\n")[1:]

	ruleMap := createRuleMap(rules)
	valids, p1 := getInvalids(ruleMap, nearby)

	p2 := p2(assignFields(decodeFields(ruleMap, valids, myTicket)), myTicket)

	fmt.Printf("Day 16, Part 1: %d\nDay 16, Part 2: %d\n", p1, p2)
}

func createRuleMap(rules []string) map[string]rule {
	rm := make(map[string]rule)
	for _, v := range rules {
		parts := strings.Split(v, ": ")
		name := parts[0]
		ranges := strings.Split(parts[1], " or ")
		r := rule{}
		for i, x := range ranges {
			strs := strings.Split(x, "-")
			a, err := strconv.Atoi(strs[0])
			utils.ErrorCheck(err)
			b, err := strconv.Atoi(strs[1])
			utils.ErrorCheck(err)
			if i == 0 {
				r.min1 = a
				r.max1 = b
			} else {
				r.min2 = a
				r.max2 = b
			}
		}
		rm[name] = r
	}
	return rm
}

func getPossibleNums(rm map[string]rule) map[int]bool {
	possible := make(map[int]bool)
	for _, r := range rm {
		for i := r.min1; i <= r.max1; i++ {
			possible[i] = true
		}
		for j := r.min2; j <= r.max2; j++ {
			possible[j] = true
		}
	}
	return possible
}

func getInvalids(rm map[string]rule, nearby []string) ([]string, int) {
	valids := []string{}
	invalidNums := []int{}
	p := getPossibleNums(rm)
	for _, n := range nearby {
		isValid := true
		vals := strings.Split(n, ",")
		for _, v := range vals {
			num, err := strconv.Atoi(v)
			utils.ErrorCheck(err)
			if _, ok := p[num]; !ok {
				invalidNums = append(invalidNums, num)
				isValid = false
			}
		}
		if isValid {
			valids = append(valids, n)
		}
	}
	total := 0
	for _, v := range invalidNums {
		total += v
	}
	return valids, total
}

func decodeFields(rm map[string]rule, valids []string, ticket string) map[string][]int {
	fields := make(map[string][]int)
	grid := createNearbyGrid(valids)
	for k, v := range rm {
		for i, r := range grid {
			valid := 0
			for _, c := range r {
				if (c >= v.min1 && c <= v.max1) || (c >= v.min2 && c <= v.max2) {
					valid++
				}
			}
			if valid == len(r) {
				fields[k] = append(fields[k], i)
			}
		}
	}
	return fields
}

func createNearbyGrid(valids []string) [][]int {
	ints := intGrid(valids)
	rl := len(ints)
	cl := len(ints[0])
	grid := make([][]int, cl)
	for i := range grid {
		grid[i] = make([]int, rl)
	}
	for i := 0; i < cl; i++ {
		for j := 0; j < rl; j++ {
			grid[i][j] = ints[j][i]
		}
	}
	return grid
}

func intGrid(strSlice []string) [][]int {
	var res [][]int
	for _, r := range strSlice {
		nums := []int{}
		slice := strings.Split(r, ",")
		for _, n := range slice {
			num, err := strconv.Atoi(n)
			utils.ErrorCheck(err)
			nums = append(nums, num)
		}
		res = append(res, nums)
	}
	return res
}

func assignFields(fields map[string][]int) []int {
	assigned := make(map[string]int)
	used := []int{}
	for i := 0; i < len(fields); i++ {
		for k, v := range fields {
			if len(v) == i+1 {
				x := filterUsed(v, used)
				assigned[k] = x
				used = append(used, x)
			}
		}
	}
	indexes := []int{}
	for k, v := range assigned {
		if k[:3] == "dep" {
			indexes = append(indexes, v)
		}
	}
	return indexes
}

func filterUsed(slice []int, used []int) int {
	for _, v := range slice {
		if !contains(used, v) {
			return v
		}
	}
	return 0
}

func contains(s []int, t int) bool {
	for _, v := range s {
		if v == t {
			return true
		}
	}
	return false
}

func p2(indexes []int, myTicket string) int {
	strs := strings.Split(myTicket, ",")
	nums := []int{}
	for _, str := range strs {
		n, err := strconv.Atoi(str)
		utils.ErrorCheck(err)
		nums = append(nums, n)
	}
	total := 1
	for _, v := range indexes {
		total *= nums[v]
	}
	return total
}
