package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func getInput(day int) string {
	inputDir := "inputs"
	inputFile := filepath.Join(inputDir, fmt.Sprintf("day%d.txt", day))

	if _, err := os.Stat(inputFile); err == nil {
		content, _ := os.ReadFile(inputFile)
		return string(content)
	}

	sessionCookie := os.Getenv("SESSION")

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, _ := http.NewRequest("GET", url, nil)

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})

	client := &http.Client{}
	resp, err := client.Do(req)

	fmt.Println(req)
	fmt.Println(err)

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	_ = os.MkdirAll(inputDir, os.ModePerm)
	_ = os.WriteFile(inputFile, body, os.ModePerm)

	return string(body)
}

func main() {
	godotenv.Load(".env")

	day := 12
	input := getInput(day)
	fmt.Println("Length of input is ", len(input))
	day12_1(input)
	day12_2(input)
}
