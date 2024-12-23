package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func day3_1(input string) {
	// r, _ := regexp.Compile("mul\\(\\d{1,3}\\,\\d{1,3}\\)")
	r, _ := regexp.Compile(`mul\(\d{1,3}\,\d{1,3}\)`) //using `` to escape characters
	r2, _ := regexp.Compile(`mul\((\d{1,3})\,(\d{1,3})\)`)
	results := r.FindAllString(input, -1)
	output := 0
	for _, result := range results {
		tmp := r2.FindAllStringSubmatch(result, -1)
		for _, item := range tmp {
			num1, _ := strconv.Atoi(item[1])
			num2, _ := strconv.Atoi(item[2])
			multiple := num1 * num2
			output += multiple
		}
	}
	fmt.Println("Output of task 1 is ", output)
}

func day3_2(input string) {
	r, _ := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	results := r.FindAllStringSubmatch(input, -1)
	output := 0
	cont := true
	for _, result := range results {
		if result[0] == "do()" {
			cont = true
			continue
		} else if result[0] == "don't()" {
			cont = false
			continue
		} else if !cont {
			continue
		}
		num1, _ := strconv.Atoi(result[1])
		num2, _ := strconv.Atoi(result[2])
		multiple := num1 * num2
		output += multiple
	}
	fmt.Println("Output of task 2 is ", output)
}
