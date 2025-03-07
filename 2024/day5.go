package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day5_1(input string) {
	// Idea is to iterate via the list of numbers to take part of numbers using delimiter |,
	// then append number 2 as index == number 1 of array
	//
	// 2nd attempts: progress the rules to create slice of slice of int,
	// with index is number 1, and vlaues as from multiple number 2
	// then set flags to progress the questions by using boolen flags as true/false when meet
	// blank line. Then iterate two new windows or two new loops from start/end,
	// if number from start loop is appeared from the values of array[index] so set flags to flase and
	// break
	//

	//To split two parts as requirements
	output := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rules := make(map[int][]int)

	updates := [][]int{}
	isAddRule := true

	for _, line := range lines {
		if len(line) == 0 {
			isAddRule = false
			continue
		}

		if isAddRule {
			nums := strings.Split(line, "|")
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			rules[num1] = append(rules[num1], num2)
		} else {
			currentUpdates := []int{}
			parts := strings.Split(line, ",")
			for _, part := range parts {
				tmp, _ := strconv.Atoi(part)
				currentUpdates = append(currentUpdates, tmp)

			}
			updates = append(updates, currentUpdates)
			//forget to put outside of loop, and conditions of loop is wrong
		}
	}
	for _, update := range updates {
		isValid := true
		for index := (len(update) - 1); index > 0; index-- {
			currentNum := update[index]
			for _, num := range update[:index] {
				for _, ruleNum := range rules[currentNum] {
					if num == ruleNum {
						isValid = false
						break
					}
				}
				if !isValid {
					break
				}
			}
		}
		if isValid {
			middle := update[len(update)/2]
			output += middle
		}
	}
	fmt.Println("Output of task 1 of day 5 is", output)
}

func day5_2(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rules := make(map[int][]int)
	output := 0

	updates := [][]int{}
	isAddRule := true

	for _, line := range lines {
		if len(line) == 0 {
			isAddRule = false
			continue
		}
		if isAddRule {
			parts := strings.Split(line, "|")
			num1, _ := strconv.Atoi(parts[0])
			num2, _ := strconv.Atoi(parts[1])

			rules[num1] = append(rules[num1], num2)
		} else {
			currentUpdate := []int{}
			nums := strings.Split(line, ",")
			for _, n := range nums {
				num, _ := strconv.Atoi(n)
				currentUpdate = append(currentUpdate, num)
			}
			updates = append(updates, currentUpdate)
		}

	}

	for _, update := range updates {
		isValid := true
		for index := len(update) - 1; index > 0; index-- {
			currentNum := update[index]
			for _, num := range update[:index] {
				for _, ruleNum := range rules[currentNum] {
					if num == ruleNum {
						isValid = false
						break
					}
				}
				if !isValid {
					break
				}
			}
		}
		if !isValid {
			ordered := []int{}
			remaining := map[int]bool{}
			dependencies := make(map[int][]int)

			for _, num := range update {
				remaining[num] = true

				for _, dep := range rules[num] {
					dependencies[dep] = append(dependencies[dep], num)
				}
			}

			for len(remaining) > 0 {
				for num := range remaining {
					if len(dependencies[num]) == 0 {
						ordered = append(ordered, num)
						delete(remaining, num)

						for key, val := range dependencies {
							newList := []int{}

							for _, n := range val {
								if n != num { // make new list without removed number.
									newList = append(newList, n)
								}
							}

							dependencies[key] = newList
						}
					}

				}
			}

			middle := ordered[len(ordered)/2]
			output += middle
		}
	}
	fmt.Println("Output of day 5 task 2 is ", output)
}
