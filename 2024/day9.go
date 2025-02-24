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

func day9_2(input string) {
	blocks := []rune{}
	isBlock := true
	fileId := 0
	for _, numRune := range input {
		num, _ := strconv.Atoi(string(numRune))

		if isBlock {
			for i := 0; i < num; i++ {
				blocks = append(blocks, rune('0'+fileId))
			}
			fileId++
		} else {
			for i := 0; i < num; i++ {
				blocks = append(blocks, '.')
			}
		}
		isBlock = !isBlock
	}

	for currentFile := fileId - 1; currentFile >= 0; currentFile-- {
		fileBlocks := []int{}
		// input := "00...111...2...333.44.5555.6666.777.888899" //examples 2

		for i, block := range blocks {
			if block == rune('0'+currentFile) {
				fileBlocks = append(fileBlocks, i)
			}
		}

		freeStart := -1 //set outside the range
		freeLength := 0

		for i := 0; i < fileBlocks[0]; i++ {
			if blocks[i] != '.' {
				freeStart = -1
				freeLength = 0
				continue
			}

			if freeStart == -1 {
				freeStart = i // start the second pointer from left to right (free space)) when i == '.'
			}

			freeLength += 1 // keep counting free space

			if freeLength == len(fileBlocks) {
				break
			}
		}

		if freeLength == len(fileBlocks) { //free space == file block (number of files)
			for j := 0; j < freeLength; j++ {
				blocks[fileBlocks[j]] = '.'                   // assign to the right most file block
				blocks[freeStart+j] = rune('0' + currentFile) // start of freespace + number of item from length of freespace to assgin the value of blockid
			}
		}
	}

	output := 0

	for blockId := 0; blockId < len(blocks); blockId++ {
		if blocks[blockId] != '.' {
			output += blockId * int(blocks[blockId]-'0')
		}
	}

	fmt.Println("Output of task 2 day 9 is", output)
}
