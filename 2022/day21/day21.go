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

type Expression struct {
	isNum bool
	left  *Expression
	right *Expression
	num   int
	op    string
}

func Solve() {

	pathOfInputText := "./2022/day21/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	parseInput(splitInput)

	res := findRootRes("root")
	fmt.Println("RES . day21 ", res)

	root, _ := monkeyMap["root"].(OpMonkey)
	left := findRootResV3(root.v1)
	right := findRootResV3(root.v2)
	r := evaluateExpression(right)
	ans := solveLHS(left, r)

	fmt.Println(left, r, ans)
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

func findRootResV3(node_name string) *Expression {
	node := monkeyMap[node_name]
	// fmt.Println(node)
	if node_name == "humn" {
		return &Expression{isNum: false, num: -2}
	}
	switch n := node.(type) {
	case ConstMonkey:
		return &Expression{isNum: true, num: n.num}
	case OpMonkey:
		v1 := findRootResV3(n.v1)
		v2 := findRootResV3(n.v2)
		switch n.op {
		case "+":
			return &Expression{left: v1, right: v2, op: "+"}
		case "*":
			return &Expression{left: v1, right: v2, op: "*"}
		case "/":
			return &Expression{left: v1, right: v2, op: "/"}
		case "-":
			return &Expression{left: v1, right: v2, op: "-"}
		}
	}
	return &Expression{isNum: false, num: -1}
}

func evaluateExpression(expr *Expression) int {
	if expr.isNum {
		return expr.num
	}

	leftVal := evaluateExpression(expr.left)
	rightVal := evaluateExpression(expr.right)

	switch expr.op {
	case "+":
		return leftVal + rightVal
	case "*":
		return leftVal * rightVal
	case "/":
		return leftVal / rightVal
	case "-":
		return leftVal - rightVal
	default:
		fmt.Printf("Unexpected operator: %s\n", expr.op)
		return 0
	}
}

func solveLHS(expr *Expression, rhs int) int {

	if !expr.isNum && expr.num == -2 {
		// This is the "humn" node
		fmt.Println("RHS", rhs)
		return rhs
	}

	leftIsNum := expr.left.isNum
	rightIsNum := expr.right.isNum
	fmt.Println(expr, expr.left, expr.right)

	var knownVal int
	var unknownExpr *Expression

	if leftIsNum && rightIsNum {
		switch expr.op {
		case "+":

			return rhs - (expr.left.num + expr.right.num)
		case "*":
			return rhs / (expr.left.num * expr.right.num)
		case "-":
			return rhs - (expr.left.num - expr.right.num)
		case "/":
			return (rhs / expr.left.num) * expr.right.num
		}

	} else if !leftIsNum && !rightIsNum {
		// Both sides are non-numeric, solve right side first, then use that result for left side
		leftResult := solveLHS(expr.left, rhs)
		return solveLHS(expr.right, leftResult)
	} else if leftIsNum {
		knownVal = expr.left.num
		unknownExpr = expr.right
	} else {
		knownVal = expr.right.num
		unknownExpr = expr.left
	}
	fmt.Println("KNOWN VAL=>>>", knownVal, expr.op, rhs, leftIsNum)

	switch expr.op {
	case "+":
		return solveLHS(unknownExpr, rhs-knownVal)
	case "*":
		println("============", rhs/knownVal, leftIsNum)
		return solveLHS(unknownExpr, rhs/knownVal)
	case "-":
		if leftIsNum {
			return solveLHS(unknownExpr, rhs-knownVal)
		} else {
			return solveLHS(unknownExpr, rhs+knownVal)
		}
	case "/":
		if leftIsNum {
			return solveLHS(unknownExpr, knownVal/rhs)
		} else {
			return solveLHS(unknownExpr, rhs*knownVal)
		}
	}

	return 1
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

// Tired but facing int overflow bugs working with testinput
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
