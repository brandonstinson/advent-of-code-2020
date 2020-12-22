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

	ts, err := strconv.Atoi(lines[0])
	utils.ErrorCheck(err)

	buses := strings.Split(lines[1], ",")

	var idsWithXs []int
	var idsWithoutXs []float64

	for _, bus := range buses {
		if bus != "x" {
			id, err := strconv.Atoi(bus)
			utils.ErrorCheck(err)
			idsWithXs = append(idsWithXs, id)
			idsWithoutXs = append(idsWithoutXs, float64(id))
		} else {
			idsWithXs = append(idsWithXs, 0)
		}
	}

	p1 := findBus(float64(ts), idsWithoutXs)

	p2 := findTimestamp(idsWithXs)

	fmt.Printf("Day 13, Part 1: %d\nDay 13, Part 2: %d\n", p1, p2)
}

func findBus(ts float64, ids []float64) int {
	lowest := ts
	lowInd := -1
	for i, id := range ids {
		n := math.Ceil(ts/id)*id - ts
		if n < lowest {
			lowest = n
			lowInd = i
		}
	}
	return int(ids[lowInd] * lowest)
}

func findTimestamp(ids []int) int {
	M := calcM(ids)
	constraints := getConstraints(ids)
	total := 0
	for _, v := range constraints {
		NI := M / v[1]
		MI := modInverse(NI, v[1])
		total += v[0] * NI * MI
	}
	return total % M
}

func calcM(ids []int) int {
	M := 1
	for _, v := range ids {
		if v != 0 {
			M *= v
		}
	}
	return M
}

func getConstraints(ids []int) [][]int {
	constraints := [][]int{}
	for i, k := range ids {
		if k != 0 {
			i %= k
			constraints = append(constraints, []int{(k - i) % k, k})
		}
	}
	return constraints
}

func modInverse(a int, m int) int {
	a %= m
	for i := 1; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}
	return 1
}
