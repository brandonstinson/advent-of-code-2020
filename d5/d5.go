package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	strs := utils.SplitFileByLine("input.txt")

	var ids []float64
	var highest float64
	for _, line := range strs {
		id := getSeatID(line, 128, 8)
		ids = append(ids, id)
		if id > highest {
			highest = id
		}
	}

	sort.Float64s(ids)
	var prev float64
	var myID float64
	for i, v := range ids {
		if i != 0 {
			prev = ids[i-1]
			if prev != v-1 {
				myID = v - 1
			}
		}
	}

	fmt.Printf("Day 5 Part 1: %v\nDay 5 Part 2: %v\n", highest, myID)
}

func getSeatID(bsp string, rows float64, cols float64) float64 {
	rowInfo := bsp[:7]
	colInfo := bsp[7:]
	var rowMin float64 = 1
	rowMax := rows
	for i, letter := range rowInfo {
		delta := rows / (math.Pow(2, float64(i)) * 2)
		if string(letter) == "F" {
			rowMax = rowMax - delta
		} else {
			rowMin = rowMin + delta
		}
	}
	var colMin float64 = 1
	colMax := cols
	for i, letter := range colInfo {
		delta := cols / (math.Pow(2, float64(i)) * 2)
		if string(letter) == "L" {
			colMax = colMax - delta
		} else {
			colMin = colMin + delta
		}
	}
	return (rowMin-1)*8 + (colMin - 1)
}
