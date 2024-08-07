package day8

import (
	readinput "advent/readInput"
	"fmt"
	"strconv"
	"strings"
)

//30373

func Solve() {
	pathOfInputText := "./2022/day8/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	processed_input := processInputText(strings.Split(input, "\n"))
	//fmt.Printf("PROCESED %v\n", processed_input)
	ans1, ans2 := findVisibleTree(processed_input)
	fmt.Printf("Day 8 Res1 : %d  \n", ans1)
	fmt.Printf("Day 8 Res2 : %d \n", ans2)
}

func findVisibleTree(input [][]int) (int, int) {
	count := 0
	rows := len(input)
	cols := len(input[0])
	maxScore := 0
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			currheight := input[r][c]
			visible := isTreeVisible(input, r, c, 0, -1, currheight) || isTreeVisible(input, r, c, 0, 1, currheight) || isTreeVisible(input, r, c, 1, 0, currheight) || isTreeVisible(input, r, c, -1, 0, currheight)
			currMaxScore := scenicScore(input, r, c, 0, -1, currheight) * scenicScore(input, r, c, 0, 1, currheight) * scenicScore(input, r, c, 1, 0, currheight) * scenicScore(input, r, c, -1, 0, currheight)
			if maxScore < currMaxScore {
				maxScore = currMaxScore
			}
			if visible {
				count++
			}

		}
	}

	return count + rows*2 + cols*2 - 4, maxScore
}

func isTreeVisible(input [][]int, r, c, dr, dc, currheight int) bool {
	for {
		r += dr
		c += dc

		if r < 0 || r >= len(input) || c < 0 || c >= len(input[0]) {
			return true
		}
		if input[r][c] >= currheight {
			return false
		}
	}
}

func scenicScore(input [][]int, r, c, dr, dc, currheight int) int {
	count := 0
	for {
		r += dr
		c += dc

		if r < 0 || r >= len(input) || c < 0 || c >= len(input[0]) {
			return count
		}
		if input[r][c] >= currheight {
			return count + 1
		}
		count++
	}
}

func processInputText(s []string) [][]int {
	result := [][]int{}
	for _, num_arr := range s {
		temp := []int{}
		for _, num := range num_arr {
			_num, _ := strconv.Atoi(string(num))
			temp = append(temp, _num)
		}

		if len(temp) > 0 {
			result = append(result, temp)
		}
	}

	return result

}
