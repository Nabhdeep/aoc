package day3

import (
	readinput "advent/readInput"
	"fmt"
	"strings"
	"unicode"
)

func Solve() {
	pathOfInputText := "./2022/day3/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sliceOfGames := strings.Split(input, "\n")
	ans1 := getArrangement(sliceOfGames)
	ans2 := getBadge(sliceOfGames)
	fmt.Printf("Day 3 Part 1: %d\n", ans1)
	fmt.Printf("Day 3 Part 2: %d\n", ans2)
}

func getArrangement(sl []string) int {
	var res int
	for _, str := range sl {
		l := len(str)
		compartment_1, compartment_2 := str[:l/2], str[l/2:]
		common := getPriority(compartment_1, compartment_2)
		res += getScore(common)
	}
	return res
}

func getBadge(sl []string) int {
	var res int
	for i := 0; i < len(sl); i += 3 {
		group := sl[i:min(i+3, len(sl))]
		// Process the group of 3 strings
		if len(group) == 3 {
			var firstString = group[0]
			for i := 1; i < len(group); i++ {
				common := getPriority(firstString, group[i])
				firstString = string(common)
			}
			res += getScore([]rune(firstString))

		}
	}
	return res

}

func getPriority(comp_1 string, comp_2 string) []rune {
	charMap := make(map[rune]bool)
	var common []rune
	for _, char := range comp_1 {
		charMap[char] = true
	}
	for _, char := range comp_2 {
		if charMap[char] {
			common = append(common, char)
			delete(charMap, char)
		}
	}
	// fmt.Printf("GOT CHAR %s\n", string(common))
	return common
}

func getScore(arr []rune) int {
	common := arr[0]
	if unicode.IsUpper(common) {
		return (int(common) - 65) + 27
	} else {
		return (int(common) - 97) + 1
	}
}
