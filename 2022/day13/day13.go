package day13

import (
	readinput "advent/readInput"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type AllSlices [][]any

func (a AllSlices) Len() int           { return len(a) }
func (a AllSlices) Less(i, j int) bool { return isRightOrder(a[i], a[j]) }
func (a AllSlices) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Solve() {
	pathOfInputText := "./2022/day13/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	res, res2 := parseInput(strings.Split(input, "\n\n"))
	fmt.Printf("DAY 13 Ans : %d\n", res)
	fmt.Printf("DAY 13 Ans : %d\n", res2)
}

func parseInput(input []string) (int, int) {
	res := 0
	var allSlice [][]any
	for idx, i := range input {
		if idx+1%3 != 0 {
			part := strings.Split(i, "\n")
			pos := 0
			_left := parseSlice(part[0], &pos)
			pos = 0
			_right := parseSlice(part[1], &pos)
			allSlice = append(allSlice, _left)
			allSlice = append(allSlice, _right)
			//COMPARE
			b := compare(_left, _right)
			// fmt.Printf("%v\n", b)
			if b == nil {
				fmt.Println("NIL")
			} else {
				if *b {
					res += idx + 1
				}
			}
			// if isRightOrder(b) {
			// 	res += idx + 1
			// }

		}

	}
	allSlice = append(allSlice, []any{[]any{2}})
	allSlice = append(allSlice, []any{[]any{6}})
	sort.Sort(AllSlices(allSlice))
	res2 := 1
	for idx, l := range allSlice {
		if reflect.DeepEqual(l, []any{[]any{2}}) || reflect.DeepEqual(l, []any{[]any{6}}) {
			res2 *= idx + 1
		}

	}
	// fmt.Println(allSlice)

	return res, res2
}

func isRightOrder(left, right []any) bool {
	sl := compare(left, right)
	if sl != nil {
		return *sl
	}
	return false
}

func compare(left, right []any) *bool {
	var res bool
	for i := 0; i < len(left) && i < len(right); i++ {
		rightInt, isRightInt := right[i].(int)
		leftInt, isLeftInt := left[i].(int)
		rightSlice, isRightSlice := right[i].([]any)
		leftSlice, isLeftSlice := left[i].([]any)

		if isRightInt && isLeftInt {
			if leftInt < rightInt {
				res = true
				return &res
			} else if leftInt > rightInt {
				res = false
				return &res
			}
			continue
		}

		if isLeftSlice && isRightSlice {
			_res := compare(leftSlice, rightSlice)
			if _res != nil {
				return _res
			}
		}

		if isLeftSlice && isRightInt {
			_rightSlice := []any{rightInt}
			_res := compare(leftSlice, _rightSlice)
			if _res != nil {
				return _res
			}
		}

		if isLeftInt && isRightSlice {
			_leftSlice := []any{leftInt}
			_res := compare(_leftSlice, rightSlice)
			if _res != nil {
				return _res
			}
		}

	}
	if len(left) < len(right) {
		res = true
		return &res
	} else if len(left) > len(right) {
		res = false
		return &res
	}
	return nil
}

func parseSlice(str string, i *int) []any {
	res := []any{}
	if str[*i] == '[' {
		currNum := ""
		*i++
		for ; *i < len(str); (*i)++ {
			// println(string(str[i]))
			if str[*i] >= '0' && str[*i] <= '9' {
				currNum += string(rune(str[*i]))
				continue
			} else if str[*i] == '[' {
				value := parseSlice(str, i)
				res = append(res, value)
			} else if str[*i] == ',' {
				// println(currNum)
				if currNum != "" {
					num, _ := strconv.Atoi(currNum)
					res = append(res, num)
				}
				currNum = ""
			} else if str[*i] == ']' {
				if currNum != "" {
					num, _ := strconv.Atoi(currNum)
					res = append(res, num)
					currNum = ""
				}
				return res
			}
		}
	}
	return res
}

// func checkPacketOrder(st string) {
// 	fmt.Println(st == "\n")
// }

// func checkSlice(left, right []string) bool {
// 	_left := reflect.TypeOf(left)
// 	_right := reflect.TypeOf(right)
// 	return _left.Kind() == reflect.Slice && _right.Kind() == reflect.Slice
// }
