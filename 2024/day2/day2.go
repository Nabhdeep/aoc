package day2

import (
	readinput "advent/readInput"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	pathOfInputText := "./2024/day2/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sl := strings.Split(input, "\n\n")
	// fmt.Println(sl)
	parsedInput, safe := parseInput(sl)
	// fmt.Println(parsedInput)
	day2 := part2(parsedInput)
	fmt.Printf("Day2 Res => %d \n", safe)
	fmt.Printf("Day2 Res2 => %d \n", day2)
}

func parseInput(sl []string) ([][]int, int) {
	var res [][]int
	var safe int
	for _, line := range strings.Split(sl[0], "\n") {
		var tempArr []int
		for _, num := range strings.Split(line, " ") {
			n, _ := strconv.Atoi(num)
			tempArr = append(tempArr, n)
		}
		// fmt.Println(isAsending(tempArr, 1), isDesending(tempArr, 1))
		if isAsending(tempArr, 1) || isDesending(tempArr, 1) {
			safe++
		}
		res = append(res, tempArr)
	}
	return res, safe
}

// 1 3 6 7 9
// 3>1 , 3-1
func isAsending(a []int, i int) bool {
	if i >= len(a) {
		return true
	}
	if a[i] > a[i-1] && (a[i]-a[i-1]) >= 1 && (a[i]-a[i-1]) <= 3 {
		return isAsending(a, i+1)
	}

	return false
}

// 7 6 4 2 1 , 1
// 6<7
// 7 6 4 2 1 , 2
// 4<6 6-4 = 2
func isDesending(a []int, i int) bool {
	if i >= len(a) {
		return true
	}
	// fmt.Println(a[i], a[i-1])
	if a[i] < a[i-1] && (a[i-1]-a[i]) >= 1 && (a[i-1]-a[i]) <= 3 {
		return isDesending(a, i+1)
	}

	return false
}

func part2(parseInput [][]int) int {
	var safe int
	for _, arr := range parseInput {
		for i, _ := range arr {
			var new_a = make([]int, len(arr)-1)
			copy(new_a, arr[:i])
			copy(new_a[i:], arr[i+1:])
			if isAsending(new_a, 1) || isDesending(new_a, 1) {
				safe++
				break
			}

		}
	}
	return safe
}
