package main

import (
	"advent/2022/day1"
	"advent/2022/day10"
	"advent/2022/day11"
	"advent/2022/day12"
	"advent/2022/day13"
	"advent/2022/day14"
	"advent/2022/day15"
	"advent/2022/day2"
	"advent/2022/day3"
	"advent/2022/day4"
	"advent/2022/day5"
	"advent/2022/day6"
	"advent/2022/day7"
	"advent/2022/day8"
	"advent/2022/day9"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func getPuzzleFromURL(day string) (string, error) {
	client := &http.Client{}

	url := fmt.Sprintf("https://adventofcode.com/2022/day/%s/input", day)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", errors.New("failed to create request")
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: os.Getenv("COOKIE_TOKEN"),
	}
	req.AddCookie(cookie)
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("failed to send request")
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp)
		return "", errors.New("received non-200 response code")
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("failed to read response body")
	}

	return string(body), nil

}

func main() {
	godotenv.Load()
	var option int
	var day int
	fmt.Println("Enter Option")
	fmt.Println("1. New Puzzle download")
	fmt.Println("2. Solve downloaded Puzzle")
	fmt.Scan(&option)
	switch option {
	case 1:
		getAndSavePuzzle()
	case 2:
		fmt.Print("Solution day?: ")
		fmt.Scan(&day)
		switch {
		case day == 1:
			day1.Solve()
		case day == 2:
			day2.Solve()
		case day == 3:
			day3.Solve()
		case day == 4:
			day4.Solve()
		case day == 5:
			day5.Solve()
		case day == 6:
			day6.Solve()
		case day == 7:
			day7.Solve()
		case day == 8:
			day8.Solve()
		case day == 9:
			day9.Solve()
		case day == 10:
			day10.Solve()
		case day == 11:
			day11.Solve()
		case day == 12:
			day12.Solve()
		case day == 13:
			day13.Solve()
		case day == 14:
			day14.Solve()
		case day == 15:
			day15.Solve()
		}
	}
}

func getAndSavePuzzle() {
	var day string
	fmt.Println("Enter day of the puzzle")
	fmt.Scanln(&day)

	fmt.Printf("Selected day is %s\n", day)
	response, err := getPuzzleFromURL(day)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Response:", response)
	}

	dir := filepath.Join("2022", fmt.Sprintf("day%s", day))
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("Failed to create directory:", err)
		return
	}

	// Create the file in the new directory
	path := filepath.Join(dir, "input.txt")
	goFileName := fmt.Sprintf("/day%s.go", day)
	goFile, _ := filepath.Abs(dir + goFileName)
	os.Create(goFile)

	err = os.WriteFile(path, []byte(response), 0644)

	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return
	}
	fmt.Printf("Wrote puzzle input to file: %s\n", path)
}
