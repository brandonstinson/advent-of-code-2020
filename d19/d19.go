package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	text, err := ioutil.ReadFile("ex.txt")
	utils.ErrorCheck(err)
	parts := strings.Split(string(text), "\n\n")
	ruleList := strings.Split(parts[0], "\n")
	messages := strings.Split(parts[1], "\n")
	ruleMap := generateRuleMap(ruleList)

	p1 := p1(ruleMap, messages)

	fmt.Printf("Day 19, Part 1: %v\n", p1)
}

func p1(rm map[string][][]string, messages []string) int {
	total := 0
	for _, m := range messages {
		if match(rm, m) {
			total++
		}
	}
	return total
}

func match(grammer map[string][][]string, s string) bool {
	return false
}

func generateRuleMap(rl []string) map[string][][]string {
	rm := make(map[string][][]string)
	for _, rule := range rl {
		parts := strings.Split(rule, ": ")
		num := parts[0]
		rule := parts[1]
		new := [][]string{}
		if string(rule[0]) == "\"" {
			rule = string(rule[1])
			new = append(new, []string{rule})
			rm[num] = new
		} else {
			options := strings.Split(rule, " | ")
			for _, v := range options {
				o := strings.Split(v, " ")
				new = append(new, o)
			}
			rm[num] = new
		}
	}
	return rm
}
