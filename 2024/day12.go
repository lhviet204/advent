package main

import (
	"fmt"
	"strings"
)

// This is a nice challenge
func day12_1(input string) {
	output := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	directions := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	visited := make(map[[2]int]bool)

	var explore func(row int, col int, plant rune) int
	explore = func(row int, col int, plant rune) int {
		area := 0
		perimeter := 0
		queue := [][2]int{{row, col}}
		visited[[2]int{row, col}] = true

		for len(queue) > 0 {
			currentPlant := queue[0]
			queue = queue[1:]
			area++

			for _, direction := range directions {
				newRow := currentPlant[0] + direction[0]
				newCol := currentPlant[1] + direction[1]

				if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[row]) {
					perimeter++ // reach to the edges
					continue
				}

				if grid[newRow][newCol] != plant {
					perimeter++
					continue
				}

				if !visited[[2]int{newRow, newCol}] {
					queue = append(queue, [2]int{newRow, newCol})
					visited[[2]int{newRow, newCol}] = true
				}
			}
		}

		return area * perimeter

	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if !visited[[2]int{row, col}] {
				output += explore(row, col, grid[row][col])
			}
		}
	}

	fmt.Println("Output of task 1 day 12 is", output)

}

// func day12_1(input string) {}
