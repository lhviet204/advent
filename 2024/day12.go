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
		//Due to "modern" business practices, the price of fence required for a region is found by multiplying that region's area by its perimeter.
		// The total price of fencing all regions on a map is found by adding together the price of fence for every region on the map

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

func day12_2(input string) {
	output := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	directions := [][2]int{
		{-1, 0}, // up
		{0, -1}, // left
		{1, 0},  // down
		{0, 1},  // right
	}

	outerCorners := []int{
		0, 0, 0, 1, 0, 0, 1,
		2, 0, 1, 0, 2,
		1, 2, 2, 4,
	}

	checkInnerCorners := [][][]int{
		{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}},
		{{1, -1}, {1, 1}},
		{{-1, 1}, {1, 1}},
		{{1, 1}},
		{{-1, -1}, {-1, 1}},
		{},
		{{-1, 1}},

		{},
		{{-1, -1}, {1, -1}},
		{{1, -1}},
		{},
		{},

		{{-1, -1}},
		{},
		{},
		{},
	}

	visited := make(map[[2]int]bool)

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if visited[[2]int{row, col}] {
				continue
			}

			plant := grid[row][col]
			area := 0
			corners := 0

			queue := [][2]int{{row, col}}

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				currentRow, currentCol := current[0], current[1]
				if visited[[2]int{currentRow, currentCol}] {
					continue
				}

				visited[[2]int{currentRow, currentCol}] = true

				area++
				cornerType := 0

				for i, direction := range directions {
					newRow := currentRow + direction[0]
					newCol := currentCol + direction[1]

					if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[row]) {
						cornerType += (1 << i)
					} else if grid[newRow][newCol] != plant {
						cornerType += (1 << i)
					} else if !visited[[2]int{newRow, newCol}] {
						queue = append(queue, [2]int{newRow, newCol})
					}
				}

				outerCornerCount := outerCorners[cornerType]
				innerCornerCount := 0

				for _, corner := range checkInnerCorners[cornerType] {
					newRow := currentRow + corner[0]
					newCol := currentCol + corner[1]

					if newRow < 0 || newRow >= len(grid) || newCol < 0 || newCol >= len(grid[row]) {
						continue
					} else if grid[newRow][newCol] != plant {
						innerCornerCount++
					}

				}

				corners += outerCornerCount + innerCornerCount
			}
			price := area * corners
			output += price
		}
	}

	fmt.Println("Output of task 2 day 12 is", output)
}
