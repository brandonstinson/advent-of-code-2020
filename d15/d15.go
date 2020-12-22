package main

import "fmt"

func main() {
	input := []int{16, 11, 15, 0, 1, 7}

	p1 := getNthNumber(input, 2020)

	p2 := getNthNumber(input, 30000000)

	fmt.Printf("Day 15, Part 1: %v\nDay 15, Part 2: %d\n", p1, p2)
}

func getNthNumber(input []int, number int) int {
	ls := make(map[int]int)
	for i, n := range input {
		if i != len(input)-1 {
			ls[n] = i
		}
	}
	for len(input) < number {
		pn := input[len(input)-1]
		pi := -1
		if v, ok := ls[pn]; ok {
			pi = v
		}
		ls[pn] = len(input) - 1
		if pi == -1 {
			input = append(input, 0)
		} else {
			n := len(input) - 1 - pi
			input = append(input, n)
		}
	}
	return input[number-1]
}
