package main

import (
	"fmt"
	"strings"
)

// Start from 0 to 9, count how many path can start from 0 to 9, start from any 0 number
func day10_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]int, len(lines))

	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, char := range line {
			grid[i][j] = int(char - '0')
		}
	}

	directions := [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	// Implement function to check
	var explore func(row int, col int, height int, reachableNine map[[2]int]bool)
	explore = func(row int, col int, height int, reachableNine map[[2]int]bool) {
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[row]) {
			return
		}

		if grid[row][col] != height {
			return // ?!
		}

		currentPosition := [2]int{row, col}
		if height == 9 {
			reachableNine[currentPosition] = true
			return
		}

		for _, direction := range directions {
			nextRow := row + direction[0] //direction is [2]int{}
			nextCol := col + direction[1]
			explore(nextRow, nextCol, height+1, reachableNine)
		}
	}

	output := 0

	//iterate via grid
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == 0 {
				reachableNine := make(map[[2]int]bool)
				explore(row, col, 0, reachableNine)
				output += len(reachableNine) // count many ways to travel
			}
		}
	}

	fmt.Println("Output of task 1 day 10 is ", output)
}

// func day10_2(input string){}
