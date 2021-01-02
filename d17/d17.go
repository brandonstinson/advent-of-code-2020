package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/brandonstinson/advent-of-code/utils"
)

type set []string

func (s *set) add(str string) {
	if !contains(*s, str) {
		*s = append(*s, str)
	}
}

func (s *set) remove(str string) {
	new := []string{}
	for _, v := range *s {
		if v != str {
			new = append(new, v)
		}
	}
	*s = new
}

func contains(s []string, t string) bool {
	for _, v := range s {
		if t == v {
			return true
		}
	}
	return false
}

func createStr(x, y, z, w int) string {
	return fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)
}

func main() {
	lines := utils.SplitFileByLine("input.txt")

	p1, p1Time := p1(lines, 6)

	p2, p2Time := p2(lines, 6)

	fmt.Printf("Day 17, Part 1: %d (in %v)\nDay 17, Part 2: %d (in %v)\n", p1, p1Time, p2, p2Time)
}

func p1(lines []string, iterations int) (int, time.Duration) {
	start := time.Now()
	var on set
	for r, l := range lines {
		for c, ch := range l {
			if string(ch) == "#" {
				on.add(createStr(r, c, 0, 0))
			}
		}
	}
	for i := 0; i < iterations; i++ {
		var new set
		for x := min(on, 0); x < max(on, 0); x++ {
			for y := min(on, 1); y < max(on, 1); y++ {
				for z := min(on, 2); z < max(on, 2); z++ {
					num := 0
					for _, dx := range []int{-1, 0, 1} {
						for _, dy := range []int{-1, 0, 1} {
							for _, dz := range []int{-1, 0, 1} {
								if dx != 0 || dy != 0 || dz != 0 {
									if contains(on, createStr(x+dx, y+dy, z+dz, 0)) {
										num++
									}
								}
							}
						}
					}
					if !contains(on, createStr(x, y, z, 0)) && num == 3 {
						new.add(createStr(x, y, z, 0))
					}
					if contains(on, createStr(x, y, z, 0)) && (num == 2 || num == 3) {
						new.add(createStr(x, y, z, 0))
					}
				}
			}
		}
		on = new
	}
	return len(on), time.Since(start)
}

func p2(lines []string, iterations int) (int, time.Duration) {
	start := time.Now()
	var on set
	for r, l := range lines {
		for c, ch := range l {
			if string(ch) == "#" {
				on.add(createStr(r, c, 0, 0))
			}
		}
	}
	for i := 0; i < iterations; i++ {
		var new set
		for x := min(on, 0); x < max(on, 0); x++ {
			for y := min(on, 1); y < max(on, 1); y++ {
				for z := min(on, 2); z < max(on, 2); z++ {
					for w := min(on, 3); w < max(on, 3); w++ {
						num := 0
						for _, dx := range []int{-1, 0, 1} {
							for _, dy := range []int{-1, 0, 1} {
								for _, dz := range []int{-1, 0, 1} {
									for _, dw := range []int{-1, 0, 1} {
										if dx != 0 || dy != 0 || dz != 0 || dw != 0 {
											if contains(on, createStr(x+dx, y+dy, z+dz, w+dw)) {
												num++
											}
										}
									}
								}
							}
						}
						if !contains(on, createStr(x, y, z, w)) && num == 3 {
							new.add(createStr(x, y, z, w))
						}
						if contains(on, createStr(x, y, z, w)) && (num == 2 || num == 3) {
							new.add(createStr(x, y, z, w))
						}
					}
				}
			}
		}
		on = new
	}
	return len(on), time.Since(start)
}

func min(on set, ind int) int {
	min := 100
	for _, v := range on {
		strs := strings.Split(v, ",")
		n, err := strconv.Atoi(strs[ind])
		utils.ErrorCheck(err)
		if n < min {
			min = n
		}
	}
	return min - 1
}

func max(on set, ind int) int {
	max := -100
	for _, v := range on {
		strs := strings.Split(v, ",")
		n, err := strconv.Atoi(strs[ind])
		utils.ErrorCheck(err)
		if n > max {
			max = n
		}
	}
	return max + 2
}
