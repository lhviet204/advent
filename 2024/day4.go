package main

import (
	"fmt"
	"strings"
)

func day4_1(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	output := 0

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	word := "XMAS"

	// Idea is to check relevant points for directions to
	// calculate points to iterate the required char. Check the index within
	// the scope of range, and compare the char
	directions := [][2]int{
		{1, 1},
		{1, -1},
		{1, 0},
		{0, 1},
		{0, -1},
		{-1, 0},
		{-1, 1},
		{-1, -1},
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			for _, dir := range directions {
				rowIndex := dir[0]
				colIndex := dir[1]
				isXmax := true
				for charIndex := 0; charIndex < len(word); charIndex++ {
					rowOffset := row + (rowIndex * charIndex)
					colOffset := col + (colIndex * charIndex)

					if rowOffset < 0 || rowOffset >= len(grid) || colOffset < 0 || colOffset >= len(grid[row]) {
						isXmax = false
						break
					}

					if grid[rowOffset][colOffset] != rune(word[charIndex]) {
						isXmax = false
						break
					}
				}
				if isXmax {
					output++
				}
			}
		}
	}
	fmt.Println("Output of task 1 is ", output)
}

func day4_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	output := 0

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	directions := [][2][2]int{
		{
			{-1, -1},
			{1, 1},
		},
		{
			{-1, 1},
			{1, -1},
		},
	}

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			//the char must be A to cont with the logic for the index
			//between start and end of lines with the condition flag to check
			if grid[row][col] != 'A' {
				continue
			}
			isXmax := true
			for _, dir := range directions {
				row1Index := row + dir[0][0]
				col1Index := col + dir[0][1]

				row2Index := row + dir[1][0]
				col2Index := col + dir[1][1]

				if row1Index < 0 || row1Index >= len(grid) || col1Index < 0 || col1Index >= len(grid[row]) || row2Index < 0 || row2Index >= len(grid) || col2Index < 0 || col2Index >= len(grid[row]) {
					isXmax = false
					break
				}

				if (grid[row1Index][col1Index]) == 'M' && (grid[row2Index][col2Index]) == 'S' || (grid[row1Index][col1Index]) == 'S' && (grid[row2Index][col2Index]) == 'M' {
					continue
				}

				isXmax = false
				break

			}
			if isXmax {
				output++
			}
		}
	}
	fmt.Println("Output of task 2 is ", output)
}
