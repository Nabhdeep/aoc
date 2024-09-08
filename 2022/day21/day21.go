package day21

import (
	readinput "advent/readInput"
	"fmt"
	"strconv"
	"strings"
)

type OpMonkey struct {
	v1 string
	v2 string
	op string
}

type ConstMonkey struct {
	num int
}

var monkeyMap = make(map[string]any)

func Solve() {

	pathOfInputText := "./2022/day21/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	parseInput(splitInput)

	res := findRootRes("root")
	fmt.Println("RES . day21 ", res)

	// root, _ := monkeyMap["root"].(OpMonkey)

	// l := findRootRes(root.v1)
	// r := findRootRes(root.v2) // side which has humn in leaf node

	right := binarySearch()
	fmt.Println("RES . day21 ", right)
}

func parseInput(input []string) {
	for _, ele := range input {
		part := strings.Split(ele, ":")
		vari := strings.Split(part[1], " ")
		if len(vari) > 2 {
			monkeyMap[strings.TrimSpace(part[0])] = OpMonkey{v1: strings.TrimSpace(vari[1]), v2: strings.TrimSpace(vari[3]), op: strings.TrimSpace(vari[2])}
		} else {
			_num, _ := strconv.Atoi(strings.TrimSpace(vari[1]))
			monkeyMap[strings.TrimSpace(part[0])] = ConstMonkey{num: _num}
		}
	}

}

func findRootRes(node_name string) int {
	node := monkeyMap[node_name]
	// fmt.Println(node)
	switch n := node.(type) {
	case ConstMonkey:
		return n.num
	case OpMonkey:
		v1 := findRootRes(n.v1)
		v2 := findRootRes(n.v2)
		switch n.op {
		case "+":
			return v1 + v2
		case "*":
			return v1 * v2
		case "/":
			return v1 / v2
		case "-":
			return v1 - v2
		}
	}
	return 0
}

func binarySearch() int {
	lo := 0
	hi := 1000000000000000
	// fmt.Println(hi, lo)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		fmt.Println(lo, hi, mid)
		left, right := prepBinarySeach(mid)
		diff := left - right
		fmt.Println(left, right)
		if diff == 0 {
			return mid
		} else if diff < 0 {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}

func prepBinarySeach(humanValue int) (int, int) {
	tempMap := make(map[string]any)
	for k, v := range monkeyMap {
		tempMap[k] = v
	}

	tempMap["humn"] = ConstMonkey{num: int(humanValue)}
	root, _ := tempMap["root"].(OpMonkey)
	left := findRootResV2(root.v1, tempMap)
	right := findRootResV2(root.v2, tempMap)
	return left, right
}

func findRootResV2(node_name string, newMap map[string]any) int {
	node := newMap[node_name]
	// fmt.Println(node)
	switch n := node.(type) {
	case ConstMonkey:
		return n.num
	case OpMonkey:
		v1 := findRootResV2(n.v1, newMap)
		v2 := findRootResV2(n.v2, newMap)
		switch n.op {
		case "+":
			return v1 + v2
		case "*":
			return v1 * v2
		case "/":
			return v1 / v2
		case "-":
			return v1 - v2
		}
	}
	return 0
}
