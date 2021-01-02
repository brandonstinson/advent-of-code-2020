package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	strs := utils.SplitFileByLine("input.txt")

	part1 := checkPassports(strs, false)

	part2 := checkPassports(strs, true)

	fmt.Printf("Day 04, Part 1: %d\nDay 04, Part 2: %d\n", part1, part2)
}

func checkPassports(lines []string, validate bool) int {
	var valid int
	var hasBYR, hasIYR, hasEYR, hasHGT, hasHCL, hasECL, hasPID bool

	for _, line := range lines {
		if line == "" {
			hasBYR = false
			hasIYR = false
			hasEYR = false
			hasHGT = false
			hasHCL = false
			hasECL = false
			hasPID = false
		} else {
			parts := strings.Split(line, " ")
			for _, field := range parts {
				fieldParts := strings.Split(field, ":")
				key := fieldParts[0]
				value := fieldParts[1]
				switch key {
				case "byr":
					if validate {
						birthYear, err := strconv.Atoi(value)
						utils.ErrorCheck(err)
						if len(value) == 4 && birthYear >= 1920 && birthYear <= 2002 {
							hasBYR = true
						}
					} else {
						hasBYR = true
					}
				case "iyr":
					if validate {
						issueYear, err := strconv.Atoi(value)
						utils.ErrorCheck(err)
						if len(value) == 4 && issueYear >= 2010 && issueYear <= 2020 {
							hasIYR = true
						}
					} else {
						hasIYR = true
					}
				case "eyr":
					if validate {
						expYear, err := strconv.Atoi(value)
						utils.ErrorCheck(err)
						if len(value) == 4 && expYear >= 2020 && expYear <= 2030 {
							hasEYR = true
						}
					} else {
						hasEYR = true
					}
				case "hgt":
					if validate {
						if len(value) >= 4 {
							num, err := strconv.Atoi(value[:len(value)-2])
							utils.ErrorCheck(err)
							unit := value[len(value)-2:]
							if unit == "cm" {
								if num >= 150 && num <= 193 {
									hasHGT = true
								}
							}
							if unit == "in" {
								if num >= 59 && num <= 76 {
									hasHGT = true
								}
							}
						}
					} else {
						hasHGT = true
					}
				case "hcl":
					if validate {
						match, err := regexp.Match(`^#[0-9a-f]{6}$`, []byte(value))
						utils.ErrorCheck(err)
						if match {
							hasHCL = true
						}
					} else {
						hasHCL = true
					}
				case "ecl":
					if validate {
						if value == "amb" || value == "blu" || value == "brn" || value == "gry" || value == "grn" || value == "hzl" || value == "oth" {
							hasECL = true
						}
					} else {
						hasECL = true
					}
				case "pid":
					if validate {
						if len(value) == 9 {
							hasPID = true
						}
					} else {
						hasPID = true
					}
				}
			}
			if hasBYR && hasIYR && hasEYR && hasHGT && hasHCL && hasECL && hasPID {
				valid++
				hasBYR = false
				hasIYR = false
				hasEYR = false
				hasHGT = false
				hasHCL = false
				hasECL = false
				hasPID = false
			}
		}
	}

	return valid
}
