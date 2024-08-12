package day11

import (
	readinput "advent/readInput"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Solve() {
	monkeys := [][]int{}

	test := []int{}
	passTo := [][]int{}
	operation := []string{}
	pathOfInputText := "./2022/day11/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	parseInput(strings.Split(input, "\n"), &monkeys, &test, &passTo, &operation)
	// fmt.Println(monkeys, passTo, operation, test)
	inspection := make([]int, len(monkeys))
	for i := 0; i < 10000; i++ {
		playMonkey(&monkeys, &operation, &test, &passTo, &inspection, false)
	}
	// fmt.Println(inspection)
	// for PART 2 change true -> false and 20 -> 10000 or PART 1 opposite of it
	res := getSum(&inspection)
	fmt.Printf("DAY 11 RES %d\n", res)
}

func getSum(inspection *[]int) int {
	slices.Sort(*inspection)
	slices.Reverse(*inspection)
	res := 1
	for i := 0; i < 2; i++ {
		res *= (*inspection)[i]
	}
	return res
}
func parseInput(arr []string, monkeys *[][]int, test *[]int, passTo *[][]int, operation *[]string) {
	for _, monkey := range arr {
		m := strings.Split(monkey, "Starting items: ")
		if len(m) >= 2 {
			worryLevels := strings.Split(m[1], ",")
			temp := []int{}
			for _, wL := range worryLevels {
				wL = strings.TrimSpace(wL)
				num, _ := strconv.Atoi(wL)
				temp = append(temp, num)
			}
			*monkeys = append((*monkeys), temp)
		}
		op := strings.Split(monkey, "Operation: ")
		if len(op) == 2 {
			(*operation) = append((*operation), op[1])
		}
		_test := strings.Split(monkey, "Test: ")
		if len(_test) == 2 {
			spTest := strings.Split(_test[1], " ")
			num, _ := strconv.Atoi(spTest[len(spTest)-1])
			*test = append(*test, num)
		}
		_if := strings.Split(monkey, "If ")
		if len(_if) >= 2 {
			_case := strings.Split(_if[1], "If ")
			_bool := strings.Split(_case[0], ":")
			_numStr := strings.Split(_bool[1], " ")
			_boolRes, _ := strconv.ParseBool(_bool[0])
			_num, _ := strconv.Atoi(_numStr[len(_numStr)-1])
			switch {
			case _boolRes:
				*passTo = append(*passTo, []int{_num})
			case !_boolRes:
				lastIndex := len(*passTo) - 1
				lastElem := (*passTo)[lastIndex]
				(*passTo)[lastIndex] = append(lastElem, _num)
			}
		}
	}
}

func doOp(operation string, old int) int {
	arrOp := strings.Split(operation, " ")
	var isDigit bool
	num, err := strconv.Atoi(arrOp[len(arrOp)-1])
	if err != nil {
		isDigit = false
	} else {
		isDigit = true
	}
	switch arrOp[len(arrOp)-2] {
	case "+":
		if isDigit {
			return old + num
		}
		return old + old
	case "*":
		if isDigit {
			return old * num
		}
		return old * old
	case "-":
		if isDigit {
			return old - num
		}
		return 0
	case "/":
		if isDigit {
			return old / num
		}
		return 1
	}
	return old
}

// func isIntegral(val float64) bool {
// 	return val == float64(int(val))
// }

func playMonkey(monkey *[][]int, operation *[]string, _test *[]int, passTo *[][]int, inspection *[]int, isPart1 bool) {
	big_mod := 1
	for _, t := range *_test {
		big_mod *= t
	}
	for i := 0; i < len(*monkey); i++ {
		//PLAY WL
		(*inspection)[i] += len((*monkey)[i])
		for _, wl := range (*monkey)[i] {
			newWL := doOp((*operation)[i], wl)
			//BORE
			if isPart1 {
				newWL = newWL / 3
			} else {
				newWL = newWL % big_mod
			}

			//TEST
			pass := (*passTo)[i]
			if newWL%(*_test)[i] == 0 {
				(*monkey)[pass[0]] = append((*monkey)[pass[0]], newWL)
			} else {
				(*monkey)[pass[1]] = append((*monkey)[pass[1]], newWL)
			}
		}

		(*monkey)[i] = []int{}
	}
}
