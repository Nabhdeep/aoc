package day1

import (
	readinput "advent/readInput"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Solve() {
	pathOfInputText := "./2024/day1/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sl := strings.Split(input, "\n\n")
	left, right, rightMap := parseTwoSlices(&sl)
	fmt.Printf("day1 res => %d", getDiff(left, right))
	fmt.Printf("day1 res2 => %d", getCount(left, rightMap))

}

func parseTwoSlices(s *[]string) ([]int, []int, map[int]int) {
	var left = []int{}
	var right = []int{}
	var rightMap = make(map[int]int)
	for _, st := range *s {
		for _, numPair := range strings.Split(st, "\n") {
			for i, num := range strings.Split(numPair, "   ") {
				n, _ := strconv.Atoi(num)
				if i == 0 {
					left = append(left, n)
				} else {
					rightMap[n] = rightMap[n] + 1
					right = append(right, n)
				}
			}
		}
	}
	// fmt.Print(left, right)2
	return left, right, rightMap
}

func getDiff(l []int, r []int) int {
	slices.Sort(l)
	slices.Sort(r)
	var dist int = 0
	for idx, num := range l {
		dist += int(math.Abs(float64(num - r[idx])))
	}
	return dist
}

func getCount(l []int, r map[int]int) int {
	var res int = 0
	for _, num := range l {
		res += r[num] * num
	}
	return res
}
