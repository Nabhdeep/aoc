package day11

import (
	readinput "advent/readInput"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve() {
	var stack []int

	pathOfInputText := "./2024/day11/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sl := strings.Split(input, "\n")
	buildInitialStack(sl, &stack)
	// part 1
	// blink(25, &stack)

	// uncomment part 1 and commnet part2 for running part 1
	//part 2
	part2 := blink2(75, &stack)
	fmt.Println(len(stack), part2)
}

var memo = make(map[string]int)

func buildInitialStack(sl []string, _s *[]int) {
	for _, ele := range strings.Split(sl[0], " ") {
		num, _ := strconv.Atoi(ele)
		// fmt.Println(ele)
		*_s = append(*_s, num)
	}
	// -1 to check the end of the list
	// *_s = append(*_s, -1)
}
func hasEvenDigits(n int) (bool, int, int) {
	_num := strconv.Itoa(int(math.Abs(float64(n))))
	_len := len(_num)
	isEven := _len%2 == 0

	if !isEven {
		return false, 0, 0
	}

	_mid := _len / 2
	left, _ := strconv.Atoi(_num[:_mid])
	right, _ := strconv.Atoi(_num[_mid:])

	return true, left, right
}

func splitRock(num int) []int {
	isEven, left, right := hasEvenDigits(num)
	if num == 0 {
		return []int{1}
	} else if isEven {
		return []int{left, right}
	}
	return []int{num * 2024}
}

func blink(numOfBlink int, arr *[]int) {

	for i := 0; i < numOfBlink; i++ {
		var newStack []int
		for _, num := range *arr {
			_rocks := splitRock(num)
			newStack = append(newStack, _rocks...)
		}
		(*arr) = newStack
	}
}

func splitRock2(num int, steps int) int {
	if steps == 0 {
		return 1
	}
	_key := fmt.Sprintf("%d_%d", num, steps)

	if val, found := memo[_key]; found {
		return val
	}

	isEven, left, right := hasEvenDigits(num)
	var result int
	if num == 0 {
		result = splitRock2(1, steps-1)
	} else if isEven {
		result = splitRock2(left, steps-1) + splitRock2(right, steps-1)
	} else {
		result = splitRock2(num*2024, steps-1)
	}
	memo[_key] = result
	return result
}

func blink2(numOfBlink int, arr *[]int) int {
	count := 0
	for _, num := range *arr {
		_l := splitRock2(num, numOfBlink)
		count += _l
	}
	return count
}
