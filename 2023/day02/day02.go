package 2023

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func puzzle01(input string) (sums int) {
	games := ParseInput(input)

	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sums = 0

	for _, game := range games {
		goodGame := true
		for color, max := range maxCubes {
			//[color][0] is number of game
			if game.cube[color][1] > max {
				goodGame = false
				break
			}
		}
		if goodGame {
			sums += game.gameNum
		}
	}
	return sums
}

// map[blue:[2 1 1 1 2] green:[4 9 5 2 14 11] red:[6 3 14 2 9 8] total:[10 14 20 5 24 21]]
// map[blue:[1 1 1 2 2] green:[2 4 5 9 11 14] red:[2 3 6 8 9 14] total:[5 10 14 20 21 24]]
// 2486

func puzzle02(input string) (multiple int) {
	// find the fewest number to make the game possible (in other way is to make that )
	games := ParseInput(input)

	multiple = 0

	for _, game := range games {
		multiple += game.cube["green"][1] * game.cube["blue"][1] * game.cube["red"][1]
	}
	return multiple
}

type gameData struct {
	gameNum int
	cube    map[string][2]int
	// to get the number of color and the max of times
}

func ParseInput(input string) (games []gameData) {
	// Parse new line
	for _, line := range strings.Split(input, "\n") {
		gameNum := 0
		fmt.Sscanf(line, "Game %d:", &gameNum)
		line := strings.TrimPrefix(line, fmt.Sprintf("Game %d:", gameNum))
		rounds := map[string][]int{
			"green": make([]int, 0),
			"red":   make([]int, 0),
			"blue":  make([]int, 0),
			"total": make([]int, 0),
		}

		// Split the ;
		for _, round := range strings.Split(line, ";") {
			// Declare total count of cubes with each color each round
			total := 0
			// Game 1:
			// 1 green, 4 blue;
			// 1 blue, 2 green, 1 red;
			// 1 red, 1 green, 2 blue;
			// 1 green, 1 red;
			// 1 green;
			// 1 green, 1 blue, 1 red
			for _, cube := range strings.Split(round, ",") { // Split the ,
				cube = strings.TrimSpace(cube)
				if cube == "" {
					continue
				}
				color := ""
				count := 0
				fmt.Sscanf(cube, "%d %s", &count, &color)
				rounds[color] = append(rounds[color], count)
				total += count
			}
			rounds["total"] = append(rounds["total"], total)
		}

		// fmt.Println(rounds) for debugs

		for _, round := range rounds {
			sort.Ints(round)
		}

		for color, round := range rounds {
			if len(round) == 0 {
				rounds[color] = []int{0}
			}
		}

		//fmt.Println(rounds) for debugs

		game := gameData{
			gameNum: gameNum,
			cube: map[string][2]int{
				"green": {rounds["green"][0], rounds["green"][len(rounds["green"])-1]},
				"red":   {rounds["red"][0], rounds["red"][len(rounds["red"])-1]},
				"blue":  {rounds["blue"][0], rounds["blue"][len(rounds["blue"])-1]},
				"total": {rounds["total"][0], rounds["total"][len(rounds["total"])-1]},
			},
		}

		games = append(games, game)
		// fmt.Println(games) // for debugs
	}

	return
}

func main() {
	// var input string = "Game 1: 1 green, 4 blue; 1 blue, 2 green, 1 red; 1 red, 1 green, 2 blue; 1 green, 1 red; 1 green; 1 green, 1 blue, 1 red"
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.TrimRight(string(dat), "\n")
	// result := puzzle01(input)ls
	result := puzzle02(input)
	fmt.Println(result)
  fmt.Println("hello")
}
