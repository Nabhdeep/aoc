package day7

//2238
import (
	readinput "advent/readInput"
	"fmt"
	"strconv"
	"strings"
)

type _rule struct {
	result int
	nums   []int
}

func Solve() {
	pathOfInputText := "./2024/day7/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sl := strings.Split(input, "\n\n")
	rulesArr := parseInput(sl[0])
	res := 0
	for _, r := range rulesArr {
		isRuleTrue := checkIsRuleTrue(r.nums[0], 0, r.result, r.nums)
		if isRuleTrue {
			res += r.result
			// fmt.Printf("Added %d to result\n", k)
		}
	}
	fmt.Println(res)
}

func parseInput(sl string) []_rule {
	var ruleArr = []_rule{}
	lines := strings.Split(sl, "\n")
	for _, line := range lines {
		rule := strings.Split(line, ": ")
		expected, _ := strconv.Atoi(rule[0])
		var tempArr []int
		for _, num := range strings.Split(rule[1], " ") {
			iNum, _ := strconv.Atoi(num)
			tempArr = append(tempArr, iNum)
		}
		ruleArr = append(ruleArr, _rule{result: expected, nums: tempArr})

	}
	return ruleArr
}

// for part 1 comment concat operation
func checkIsRuleTrue(actual int, index int, expected int, numArr []int) bool {
	next := index + 1
	if actual == expected && index == len(numArr)-1 {
		return true
	}

	if actual > expected || next >= len(numArr) {
		return false
	}

	plus := checkIsRuleTrue(actual+numArr[next], next, expected, numArr)
	concat := checkIsRuleTrue(getConcat(actual, numArr[next]), next, expected, numArr)
	product := checkIsRuleTrue(actual*numArr[next], next, expected, numArr)
	return (plus || product || concat)

}

func checkIsRuleTrueV2(target int, arr []int) bool {
	if len(arr) == 1 {
		return target == arr[0]
	}

	if target%arr[len(arr)-1] == 0 && checkIsRuleTrueV2((target/arr[len(arr)-1]), arr[:len(arr)-1]) {
		return true
	}

	if target > arr[len(arr)-1] && checkIsRuleTrueV2((target-arr[len(arr)-1]), arr[:len(arr)-1]) {
		return true
	}

	return false
}

func getConcat(a, b int) int {
	_a := strconv.Itoa(a)
	_b := strconv.Itoa(b)
	_c, _ := strconv.Atoi(_a + _b)
	return _c
}
