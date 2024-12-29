package day3

import (
	readinput "advent/readInput"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Solve() {
	pathOfInputText := "./2024/day3/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sl := strings.Split(input, "\n\n")
	day3Res := regexSol(sl)
	day3Res2 := regexSol2(sl)
	fmt.Println(day3Res)
	fmt.Println(day3Res2)
}

func regexSol(sl []string) int {
	var _rs int
	re := regexp.MustCompile(`mul\((\d*,\d*)\)`)
	_mul := re.FindAllString(sl[0], -1)
	for _, m := range _mul {
		_rs += mul(m)
	}
	return _rs
}

func regexSol2(sl []string) int {
	var _rs int
	re := regexp.MustCompile(`mul\((\d*,\d*)\)|do\(\)|don't\(\)`)
	_mul_do_dont := re.FindAllString(sl[0], -1)
	var _f bool = true // willbe flag
	for _, s := range _mul_do_dont {
		switch {
		case s == "don't()":
			_f = false
			continue
		case s == "do()":
			_f = true
			continue
		}
		if _f {
			_rs += mul(s)
		}
	}
	return _rs
}

func mul(s string) int {
	re := regexp.MustCompile(`(\d*,\d*)`)
	_nums := re.FindAllString(s, -1)
	if len(_nums) != 1 {
		fmt.Println(_nums)
		panic("Error nums not in len 1")
	}
	nums := strings.Split(_nums[0], ",")

	mul := 1

	for _, n := range nums {
		i, _ := strconv.Atoi(n)
		mul *= i
	}
	return mul
}
