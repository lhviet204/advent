package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func FetchSliceOfIntsInString(line string) []int {
	nums := []int{}
	var build strings.Builder
	isNegative := false
	for _, char := range line {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}

		if char == '-' {
			isNegative = true
		}

		if (char == ' ' || char == ',' || char == '~' || char == '|') && build.Len() != 0 {
			localNum, err := strconv.ParseInt(build.String(), 10, 64)
			if err != nil {
				panic(err)
			}
			if isNegative {
				localNum *= -1
			}
			nums = append(nums, int(localNum))
			build.Reset()
			isNegative = false
		}
	}
	if build.Len() != 0 {
		localNum, err := strconv.ParseInt(build.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		if isNegative {
			localNum *= -1
		}
		nums = append(nums, int(localNum))
		build.Reset()
	}
	return nums
}

func day7_2plus(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	output := 0
	for _, line := range lines {
		// parts := strings.Split(line, ":")
		// target, _ := strconv.Atoi(parts[0])
		// numberStrings := strings.Fields(parts[1])
		// numbers := make([]int, len(numberStrings))

		// for i, numString := range numberStrings {
		// 	numbers[i], _ = strconv.Atoi(numString)
		// }
		nums := FetchSliceOfIntsInString(line)
		if isMatch(nums[0], 0, nums[1:]) {
			output += nums[0]
		}
	}
	fmt.Println("Output of day 7 task 2", output)
}
func calculate(a, b int, operator byte) int {
	result := 0
	switch operator {
	case '+':
		result = a + b
	case '*':
		result = a * b
	case '|':
		mul, q := 10, 10 // cont to divide 10
		for q != 0 {
			q = b / mul
			if q > 0 {
				mul *= 10
			}
		}
		result = (a * mul) + b
	}
	return result
}

func isMatch(target int, currentResult int, numbers []int) bool {
	if len(numbers) == 0 {
		return currentResult == target
	}

	if currentResult > target {
		return false
	}

	if isMatch(target, calculate(currentResult, numbers[0], '+'), numbers[1:]) {
		return true
	}
	if isMatch(target, calculate(currentResult, numbers[0], '|'), numbers[1:]) {
		return true
	}
	return isMatch(target, calculate(currentResult, numbers[0], '*'), numbers[1:])
}
