package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brandonstinson/advent-of-code/utils"
)

type bagSet []string

func (bs *bagSet) add(bags []string) {
	for _, b := range bags {
		if !contains(*bs, b, 0) {
			*bs = append(*bs, b)
		}
	}
}

func main() {
	strs := utils.SplitFileByLine("input.txt")

	target := "shiny gold"
	rules := make(map[string][]string)

	for _, line := range strs {
		parts := strings.Split(line, " bags contain ")
		bag := parts[0]
		var contents []string
		for _, i := range strings.Split(parts[1], ", ") {
			words := strings.Split(i, " ")
			if words[0] == "no" {
				contents = append(contents, "none")
			} else {
				color := strings.Join(words[:len(words)-1], " ")
				contents = append(contents, color)
			}
		}
		rules[bag] = contents
	}

	var canHoldTarget bagSet
	new := canHold(rules, canHoldTarget, []string{target})
	canHoldTarget.add(new)

	count := countBags(rules, "1 shiny gold")

	fmt.Printf("Day 07, Part 1: %d\nDay 07, Part 2: %d\n", len(canHoldTarget), count)
}

func contains(slice []string, str string, start int) bool {
	for _, v := range slice {
		if v[start:] == str {
			return true
		}
	}
	return false
}

func canHold(rules map[string][]string, initial []string, targets []string) []string {
	var new []string
	for k, v := range rules {
		for _, t := range targets {
			if contains(v, t, 2) {
				if !contains(initial, k, 0) {
					new = append(new, k)
				}
			}
		}
	}
	new = append(new, initial...)
	if len(new) == len(initial) {
		return new
	}
	return canHold(rules, new, new)
}

func countBags(rules map[string][]string, bag string) int {
	contains := rules[bag[2:]]

	if len(contains) == 1 && contains[0] == "none" {
		return 0
	}

	count := 0

	for _, c := range contains {
		num, err := strconv.Atoi(c[:1])
		utils.ErrorCheck(err)
		count += num + (num * countBags(rules, c))
	}

	return count
}
