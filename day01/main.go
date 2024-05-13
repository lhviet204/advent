package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input string

func puzzle1(input string) int {
	results := parseInputs(input)
	total := 0
	for _, result := range results {
		total += result
	}
	return total

}

func parseInputs(input string) (results []int) {
	for _, line := range strings.Split(input, "\n") {
		digits := make([]int, 0)
		for i := range line {
			if line[i] >= '0' && line[i] <= '9' {
				digit, err := strconv.Atoi(string(line[i]))
				if err != nil {
					log.Fatal(err)
				}
				digits = append(digits, digit)
				continue
			}
		}
		if len(digits) == 0 {
			log.Fatalf("No digit are found on the line %s", line)
		}
		result := digits[0] * 10
		result += digits[len(digits)-1]

		results = append(results, result)
	}
	return
}

func main() {
	// return the map of byte type
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	//string type conversion for byte.
	input = strings.TrimRight(string(dat), "\n")
	fmt.Println("Solving problems ....")

	solution := 0
	solution = puzzle1(input)

	fmt.Println("Solution is: ", solution)

}
