package day25

import (
	readinput "advent/readInput"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve() {
	pathOfInputText := "./2022/day25/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	res := runInput(splitInput)
	fmt.Println(res)
}

func runInput(input []string) string {
	var number int
	for _, ele := range input {
		number += SNAFUToDecimal(ele)
	}
	return DecimalToSNAFU(number)
}

func SNAFUToDecimal(s string) int {
	l := len(s) - 1
	runningTotal := 0

	for _, ele := range s {
		num, err := strconv.Atoi(string(ele))
		if err != nil {
			switch {
			case ele == '=':
				runningTotal += -2 * int(math.Pow(5, float64(l)))
			case ele == '-':
				runningTotal += -1 * int(math.Pow(5, float64(l)))
			}
		}

		runningTotal += num * int(math.Pow(5, float64(l)))
		l--
	}

	return runningTotal
}

func DecimalToSNAFU(i int) string {
	remaining := 0
	res := ""
	for i > 0 {
		remaining = i % 5
		i = i / 5

		if remaining <= 2 {
			s := strconv.Itoa(remaining)
			res = s + res
		} else {
			if remaining == 4 {
				res = "-" + res
			} else {
				res = "=" + res
			}
			i += 1
		}

	}

	return res
}
