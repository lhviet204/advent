package main

import (
	"fmt"
	"strconv"
)

func day9_1(input string) {
	blocks := []rune{}
	isBlock := true //tracking the block of file
	fileId := 0     // tracking the id of file

	//2333133121414131402
	//00...111...2...333.44.5555.6666.777.888899 Char
	//48 48 46 46 46 Dec
	for _, numRune := range input {
		num, _ := strconv.Atoi(string(numRune))

		if isBlock {
			for i := 0; i < num; i++ {
				blocks = append(blocks, rune('0'+fileId)) //This expression converts an integer (fileId) into its character representation in the ASCII table.
			}
			fileId++
		} else {
			for i := 0; i < num; i++ {
				blocks = append(blocks, '.') //46 - Decimal value of char .
			}
		}
		isBlock = !isBlock // set flag to check block file
	}
	// set up two pointers to swap
	emptyIndex := 0
	lastFileIndex := len(blocks) - 1

	for blocks[emptyIndex] != '.' {
		emptyIndex++
	}
	for blocks[lastFileIndex] == '.' {
		lastFileIndex--
	}
	for emptyIndex <= lastFileIndex {
		blocks[emptyIndex] = blocks[lastFileIndex] // swap tail and head values
		blocks[lastFileIndex] = '.'                // assign . back to tail values

		for blocks[emptyIndex] != '.' {
			emptyIndex++
		}
		for blocks[lastFileIndex] == '.' {
			lastFileIndex--
		}
	}

	blockId := 0
	output := 0

	for blocks[blockId] != '.' {
		output += blockId * int(blocks[blockId]-'0')
		blockId++
	}

	fmt.Println("Output of task 1 day 9 is", output)
}

// The problem is try to iterate through the numbers, and sort
