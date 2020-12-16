package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/brandonstinson/advent-of-code/utils"
)

type movement struct {
	instruction string
	distance    int
}

func main() {
	lines := utils.SplitFileByLine("input.txt")

	var movements []movement

	for _, line := range lines {
		i := line[:1]
		d, e := strconv.Atoi(line[1:])
		utils.ErrorCheck(e)
		m := movement{}
		m.instruction = i
		m.distance = d
		movements = append(movements, m)
	}

	md1 := calcPos1(movements)

	md2 := calcPos2(movements)

	fmt.Printf("Day 12, Part 1: %d\nDay 12, Part 2: %d\n", md1, md2)
}

func calcPos1(movements []movement) int {
	x := 0
	y := 0
	dirInd := 1
	for _, m := range movements {
		switch m.instruction {
		case "N":
			y += m.distance
		case "S":
			y -= m.distance
		case "E":
			x += m.distance
		case "W":
			x -= m.distance
		case "L":
			dirInd = (dirInd - (m.distance / 90) + 4) % 4
		case "R":
			dirInd = (dirInd + (m.distance / 90) + 4) % 4
		case "F":
			if dirInd == 0 {
				y += m.distance
			}
			if dirInd == 1 {
				x += m.distance
			}
			if dirInd == 2 {
				y -= m.distance
			}
			if dirInd == 3 {
				x -= m.distance
			}

		}
	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func calcPos2(movements []movement) int {
	px := 0
	py := 0
	wx := 10
	wy := 1
	for _, m := range movements {
		switch m.instruction {
		case "N":
			wy += m.distance
		case "S":
			wy -= m.distance
		case "E":
			wx += m.distance
		case "W":
			wx -= m.distance
		case "L":
			for i := 0; i < m.distance/90; i++ {
				wx, wy = -wy, wx
			}
		case "R":
			for i := 0; i < m.distance/90; i++ {
				wx, wy = wy, -wx
			}
		case "F":
			px += wx * m.distance
			py += wy * m.distance
		}
	}
	return int(math.Abs(float64(px)) + math.Abs(float64(py)))
}
