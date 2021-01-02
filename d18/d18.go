package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	lines := utils.SplitFileByLine("input.txt")

	p1, p2 := sumLines(lines)

	fmt.Printf("Day 18, Part 1: %d\nDay 18, Part 2: %d\n", p1, p2)
}

func sumLines(lines []string) (int, int) {
	p1total := 0
	p2total := 0
	for _, l := range lines {
		l = strings.ReplaceAll(l, "(", "( ")
		l = strings.ReplaceAll(l, ")", " )")
		parts := strings.Split(l, " ")
		p1total += totalSubString(eliminateParentheses(parts, false))
		p2 := totalSubString(eliminateAddition(eliminateParentheses(parts, true)))
		p2total += p2
	}
	return p1total, p2total
}

func totalSubString(ss []string) int {
	total, err := strconv.Atoi(ss[0])
	utils.ErrorCheck(err)
	for i := 1; i < len(ss); i += 2 {
		num, err := strconv.Atoi(ss[i+1])
		utils.ErrorCheck(err)
		if ss[i] == "+" {
			total += num
		}
		if ss[i] == "*" {
			total *= num
		}
	}
	return total
}

func eliminateParentheses(line []string, operatorPrecedence bool) []string {
	lastOpenIndex := -1
	for i, v := range line {
		if v == "(" {
			lastOpenIndex = i
		}
		if v == ")" {
			var subStrTotal int
			if operatorPrecedence {
				subStrTotal = totalSubString(eliminateAddition(line[lastOpenIndex+1 : i]))
			} else {
				subStrTotal = totalSubString(line[lastOpenIndex+1 : i])
			}
			new := []string{}
			new = append(new, line[:lastOpenIndex]...)
			new = append(new, fmt.Sprint(subStrTotal))
			new = append(new, line[i+1:]...)
			return eliminateParentheses(new, operatorPrecedence)
		}
	}
	return line
}

func eliminateAddition(ss []string) []string {
	for i := 1; i < len(ss); i += 2 {
		if ss[i] == "+" {
			subStrTotal := totalSubString(ss[i-1 : i+2])
			new := []string{}
			if i > 1 {
				new = append(new, ss[:i-1]...)
			}
			new = append(new, fmt.Sprint(subStrTotal))
			if i < len(ss)-1 {
				new = append(new, ss[i+2:]...)
			}
			return eliminateAddition(new)
		}
	}
	return ss
}
