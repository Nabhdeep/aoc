package day4

import (
	readinput "advent/readInput"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	pathOfInputText := "./2022/day4/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sliceOfGames := strings.Split(input, "\n")
	ans1 := getPairs(sliceOfGames)
	ans2 := getPairsV2(sliceOfGames)
	fmt.Printf("Day 4 Part 1: %d\n", ans1)
	fmt.Printf("Day 4 Part 2: %d\n", ans2)
}

func getPairs(pairs []string) int {
	var count int
	for _, pair := range pairs {
		p := strings.Split(pair, ",")
		var prev []string
		for _, nums := range p {
			if prev == nil {
				prev = strings.Split(nums, "-")
				continue
			}
			curr := strings.Split(nums, "-")

			prevStart, err1 := strconv.Atoi(prev[0])
			prevEnd, err2 := strconv.Atoi(prev[1])
			currStart, err3 := strconv.Atoi(curr[0])
			currEnd, err4 := strconv.Atoi(curr[1])

			if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
				if (prevStart <= currStart && prevEnd >= currEnd) ||
					(prevStart >= currStart && prevEnd <= currEnd) {
					// fmt.Printf("Prev %s , Curr: %s\n", prev, curr)
					count++
				}
			} else {
				continue
			}

		}

	}
	return count
}

func getPairsV2(pairs []string) int {
	var count int
	for _, pair := range pairs {
		p := strings.Split(pair, ",")
		var prev []string
		for _, nums := range p {
			if prev == nil {
				prev = strings.Split(nums, "-")
				continue
			}
			curr := strings.Split(nums, "-")

			prevStart, err1 := strconv.Atoi(prev[0])
			prevEnd, err2 := strconv.Atoi(prev[1])
			currStart, err3 := strconv.Atoi(curr[0])
			currEnd, err4 := strconv.Atoi(curr[1])

			if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
				if prevStart <= currEnd && prevEnd >= currStart {
					// fmt.Printf("Prev %s , Curr: %s\n", prev, curr)
					count++
				}
			} else {
				continue
			}

		}

	}
	return count
}
