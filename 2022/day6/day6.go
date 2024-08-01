package day6

import (
	readinput "advent/readInput"
	"fmt"
)

func Solve() {
	pathOfInputText := "./2022/day6/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	// fmt.Println(input)
	str, idx := findMarker(input, 4)
	str1, idx2 := findMarker(input, 14)
	fmt.Printf("Day 6 ans1 : %s ,%d \n", str, idx)
	fmt.Printf("Day 6 ans2 : %s ,%d \n", str1, idx2)
}

func findMarker(sl string, marker int) (string, int) {
	charCount := make(map[rune]int)

	for i, char := range sl {
		charCount[char]++

		if i >= marker {
			oldChar := rune(sl[i-marker])
			charCount[oldChar]--
			if charCount[oldChar] == 0 {
				delete(charCount, oldChar)
			}
		}

		if len(charCount) == marker {
			return string(char), i + 1
		}
	}

	return "", 0
}
