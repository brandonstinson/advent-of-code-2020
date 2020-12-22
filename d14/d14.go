package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	lines := utils.SplitFileByLine("input.txt")

	p1 := p1(lines)

	p2 := p2(lines)

	fmt.Printf("Day 14, Part 1: %d\nDay 14, Part 2: %d\n", p1, p2)
}

func p1(lines []string) int {
	memory := make(map[string]int)
	m := ""
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		val := parts[1]
		if line[:4] == "mask" {
			m = val
		} else {
			a := strings.TrimSuffix((strings.Split(parts[0], "[")[1]), "]")
			memory[a] = maskp1(m, val)
		}

	}
	total := 0
	for _, v := range memory {
		total += v
	}
	return total
}

func p2(lines []string) int {
	memory := make(map[int]int)
	m := ""
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		val := parts[1]
		if line[:4] == "mask" {
			m = val
		} else {
			a := strings.TrimSuffix((strings.Split(parts[0], "[")[1]), "]")
			ms := maskp2(m, a)
			for _, k := range ms {
				v, e := strconv.Atoi(val)
				utils.ErrorCheck(e)
				memory[k] = v
			}
		}

	}
	total := 0
	for _, v := range memory {
		total += v
	}
	return total
}

func maskp1(mask string, val string) int {
	res := ""
	b := numToBinary(val)
	for i, r := range mask {
		switch string(r) {
		case "1":
			res += "1"
		case "0":
			res += "0"
		default:
			res += string(b[i])
		}
	}
	return binaryToInt(res)
}

func maskp2(mask string, val string) []int {
	res := ""
	b := numToBinary(val)
	for i, r := range mask {
		switch string(r) {
		case "1":
			res += "1"
		case "X":
			res += "X"
		default:
			res += string(b[i])
		}
	}
	return getPermutations(res)
}

func numToBinary(s string) string {
	n, e := strconv.Atoi(s)
	utils.ErrorCheck(e)
	return fmt.Sprintf("%036b", n)
}

func binaryToInt(b string) int {
	n, e := strconv.ParseInt(b, 2, 64)
	utils.ErrorCheck(e)
	return int(n)
}

func getPermutations(s string) []int {
	xi := []int{}
	var bt int
	for i, r := range s {
		if string(r) == "X" {
			xi = append(xi, 35-i)
		}
		if string(r) == "1" {
			bt += int(math.Pow(2, 35-float64(i)))
		}
	}
	p := []int{bt}
	for _, v := range xi {
		for _, i := range p {
			p = append(p, i+int(math.Pow(2, float64(v))))
		}
	}
	return p
}
