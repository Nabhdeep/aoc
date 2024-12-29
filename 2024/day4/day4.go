package day4

import (
	readinput "advent/readInput"
	"fmt"
	"strings"
)

var XMAS = []string{"X", "M", "A", "S"}
var Paths = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func Solve() {
	pathOfInputText := "./2024/day4/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sl := strings.Split(input, "\n\n")
	parsedInput := parseInput(sl[0])
	// fmt.Println(parsedInput)
	day1_res1 := countXMAS(parsedInput)
	fmt.Println(day1_res1)
	day2_res2 := countX_MAS(parsedInput)
	fmt.Println(day2_res2)
}

func parseInput(sl string) [][]string {
	var arr [][]string
	for _, line := range strings.Split(sl, "\n") {
		var a []string
		a = append(a, strings.Split(line, "")...)
		arr = append(arr, a)
	}
	return arr
}

func countXMAS(arr [][]string) int {
	var count int = 0
	for r := range arr {
		for c := range arr[0] {
			if arr[r][c] == "X" {
				// fmt.Println("FOR", r, c)
				for _, path := range Paths {
					dr := path[0]
					dc := path[1]
					flag := ""
					// fmt.Println("DIR", dr, dc)
					for _, d := range []int{0, 1, 2, 3} {
						nr := (r + dr*d)
						nc := (c + dc*d)
						if nr >= 0 && nr < len(arr) && nc >= 0 && nc < len(arr[0]) {
							flag += arr[nr][nc]
						}
					}
					if flag == "XMAS" {
						count++
					}

				}
			}
		}
	}
	return count
}

var xPath = [][]int{{-1, -1}, {1, -1}, {1, 1}, {-1, 1}}

func countX_MAS(arr [][]string) int {
	var count int = 0
	for r := range arr {
		for c := range arr[0] {
			if arr[r][c] == "A" {
				m := 0
				s := 0
				for _, p := range xPath {
					nr := p[0] + r
					nc := p[1] + c
					if nr >= 0 && nr < len(arr) && nc >= 0 && nc < len(arr[0]) {
						if arr[nr][nc] == "S" {
							s++
						} else if arr[nr][nc] == "M" {
							m++
						}
					}

				}
				if s == 2 && m == 2 && arr[r+1][c+1] != arr[r-1][c-1] {
					count++
				}
			}
		}
	}
	return count
}
