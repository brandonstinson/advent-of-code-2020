package main

import (
	"fmt"

	"github.com/brandonstinson/advent-of-code/utils"
)

func main() {
	lines := utils.SplitFileByLine("input.txt")

	p1seats, p1changes, p1occupied := changeSeats(lines, true)

	p2seats, p2changes, p2occupied := changeSeats(lines, false)

	for p1changes != 0 {
		p1seats, p1changes, p1occupied = changeSeats(p1seats, true)
	}

	for p2changes != 0 {
		p2seats, p2changes, p2occupied = changeSeats(p2seats, false)
	}

	fmt.Printf("Day 11, Part 1: %d\nDay 11, Part 2: %d\n", p1occupied, p2occupied)
}

func changeSeats(seats []string, immediate bool) ([]string, int, int) {
	changes := 0
	occupied := 0
	var newSeats []string
	for i, row := range seats {
		var newRow string
		for j, seat := range row {
			switch string(seat) {
			case "L":
				adj := -1
				if immediate == true {
					adj = checkImmediatelyAdjacent(seats, i, j)
				} else {
					adj = checkExtendedAdjacent(seats, i, j)
				}
				if adj == 0 {
					newRow += "#"
					changes++
					occupied++
				} else {
					newRow += string(seat)
				}
			case "#":
				adj := -1
				var threshold int
				if immediate == true {
					adj = checkImmediatelyAdjacent(seats, i, j)
					threshold = 4
				} else {
					adj = checkExtendedAdjacent(seats, i, j)
					threshold = 5
				}
				if adj >= threshold {
					newRow += "L"
					changes++
				} else {
					newRow += string(seat)
					occupied++
				}
			default:
				newRow += string(seat)
			}
		}
		newSeats = append(newSeats, newRow)
	}
	return newSeats, changes, occupied
}

func checkImmediatelyAdjacent(seats []string, row int, col int) int {
	adj := 0
	// previous row
	if row > 0 {
		if col > 0 {
			if string(seats[row-1][col-1]) == "#" {
				adj++
			}
		}
		if string(seats[row-1][col]) == "#" {
			adj++
		}
		if col < len(seats[row])-1 {
			if string(seats[row-1][col+1]) == "#" {
				adj++
			}
		}
	}
	// current row
	if col > 0 {
		if string(seats[row][col-1]) == "#" {
			adj++
		}
	}
	if col < len(seats[row])-1 {
		if string(seats[row][col+1]) == "#" {
			adj++
		}
	}
	// next row
	if row < len(seats)-1 {
		if col > 0 {
			if string(seats[row+1][col-1]) == "#" {
				adj++
			}
		}
		if string(seats[row+1][col]) == "#" {
			adj++
		}
		if col < len(seats[row])-1 {
			if string(seats[row+1][col+1]) == "#" {
				adj++
			}
		}
	}
	return adj
}

func checkExtendedAdjacent(seats []string, row int, col int) int {
	adj := 0
	for _, dr := range []int{-1, 0, 1} {
		for _, dc := range []int{-1, 0, 1} {
			if dr == 0 && dc == 0 {
				continue
			}
			r := row + dr
			c := col + dc
			for r >= 0 && r < len(seats) && c >= 0 && c < len(seats[0]) && string(seats[r][c]) == "." {
				r += dr
				c += dc
			}
			if r >= 0 && r < len(seats) && c >= 0 && c < len(seats[0]) && string(seats[r][c]) == "#" {
				adj++
			}
		}
	}
	return adj
}
