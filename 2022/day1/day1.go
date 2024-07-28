package day1

import (
	readinput "advent/readInput"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	pathOfInputText := "./2022/day1/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sl := strings.Split(input, "\n\n")
	var maxArray []int
	calculateElfWithMostCalories(&sl, &maxArray)
	sort.Slice(maxArray, func(i, j int) bool {
		return maxArray[i] > maxArray[j]
	})
	topThreeCals := maxArray[0:3]

	fmt.Printf("Ans 1 Calories with ELF is %d\n", maxArray[0])
	fmt.Printf("Ans 2 Top 3 Calories with ELF is %d\n", sumSlice(&topThreeCals))
}

func calculateElfWithMostCalories(sl *[]string, maxArr *[]int) {
	for _, sub := range *sl {
		// fmt.Println(idx, sub)
		localMax := 0
		for _, strNum := range strings.Split(sub, "\n") {
			num, err := strconv.Atoi(strNum)
			if err != nil {
				fmt.Println(sub)
				fmt.Printf("Due to %s", strNum)
				panic("Wrong num")
			}
			localMax += num
			*maxArr = append(*maxArr, localMax)
		}
	}
}

func sumSlice(sl *[]int) int {
	var res int
	for _, num := range *sl {
		res += num
	}
	return res
}
