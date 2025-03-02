package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day11_1(input string) {
	strNums := strings.Fields(input)
	output := 0

	stones := []int{}

	for _, stone := range strNums {
		num, _ := strconv.Atoi(stone)
		stones = append(stones, num)
	}

	processBlink := func(stones []int) []int {
		newStones := []int{}

		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if len(strconv.Itoa(stone))%2 == 0 {
				numStr := strconv.Itoa(stone)
				half := len(numStr) / 2 // check number of digits of num
				left, _ := strconv.Atoi(numStr[:half])
				right, _ := strconv.Atoi(numStr[half:])
				newStones = append(newStones, left, right)
			} else {
				newStones = append(newStones, stone*2024)
			}
		}

		return newStones
	}

	for i := 0; i < 25; i++ {
		stones = processBlink(stones)
	}

	output = len(stones)
	fmt.Println("Output of task 1 day 11 is", output)
}

func day11_2(input string) {
	// Calculate how many type of stones based on the rules in a new variant way
	// like hash map to count how many times number shows up
	numStr := strings.Fields(input)
	stones := make(map[int]int)
	for _, stone := range numStr {
		num, _ := strconv.Atoi(stone)
		stones[num]++
	}

	for i := 0; i < 75; i++ { // for range 75
		newStones := map[int]int{}

		for stone, count := range stones {
			if stone == 0 {
				newStones[1] += count
			} else if len(strconv.Itoa(stone))%2 == 0 {
				numStr := strconv.Itoa(stone) // get the str of number
				half := len(numStr) / 2       // get the halt of number
				left, _ := strconv.Atoi(numStr[:half])
				right, _ := strconv.Atoi(numStr[half:])
				newStones[left] += count
				newStones[right] += count

			} else {
				newStones[stone*2024] += count
			}
		}
		stones = newStones
	}
	output := 0

	for _, count := range stones {
		output += count
	}

	fmt.Println("Output of task 2 day 11 is", output)
}
