package day5

import (
	readinput "advent/readInput"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// after - before list
type ab struct {
	a map[int]bool
	b map[int]bool
}

var numMap = make(map[int]ab)
var pagesCheck = [][]int{}

// make a map with num -> a [nums] , b [nums]
// for pages
// check each page in num map whether after and before nums are present in the map

func Solve() {
	pathOfInputText := "./2024/day5/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sl := strings.Split(input, "\n\n")
	parseInput(sl)
	// fmt.Println(pagesCheck)
	res := 0
	res2 := 0
	for _, check := range pagesCheck {
		ar := solvePart1(check)
		if ar {
			res += getMid(check)
		} else {
			a := sortArr(check)
			// fmt.Println(a)
			res2 += getMid(a)
		}
	}
	fmt.Println(res)
	fmt.Println(res2)
}

func parseInput(sl []string) {

	rule := strings.Split(sl[0], "\n")

	pages := strings.Split(sl[1], "\n")
	for _, p := range pages {
		ele := strings.Split(p, ",")
		temp := []int{}
		for _, e := range ele {
			num, _ := strconv.Atoi(e)
			temp = append(temp, num)
		}
		pagesCheck = append(pagesCheck, temp)
	}
	for _, s := range rule {
		sp := strings.Split(s, "|")
		if len(sp) == 2 {

			_b, _ := strconv.Atoi(sp[0])
			_a, _ := strconv.Atoi(sp[1])
			b_val, b_ok := numMap[_b]
			a_val, a_ok := numMap[_a]
			if b_ok {
				b_val.a[_a] = true
				numMap[_b] = b_val
			} else {
				numMap[_b] = ab{a: map[int]bool{_a: true}, b: map[int]bool{}}
			}

			if a_ok {
				a_val.b[_b] = true
				numMap[_a] = a_val
			} else {
				numMap[_a] = ab{b: map[int]bool{_b: true}, a: map[int]bool{}}
			}
		}
	}

}

func solvePart1(arr []int) bool {

	for i := 0; i < len(arr); i++ {
		before := arr[:i]
		curr := arr[i]
		after := arr[i+1:]
		mapCurr := numMap[curr]
		if len(before) > 0 {
			for _, b := range before {
				_, ok := mapCurr.b[b]
				if !ok {
					return false
				}
			}
		}
		if len(after) > 0 {
			for _, a := range after {
				_, ok := mapCurr.a[a]
				if !ok {
					return false
				}
			}
		}

	}

	return true

}

func getMid(arr []int) int {
	i := len(arr) / 2
	return arr[int(math.Ceil(float64(i)))]
}

func sortArr(arr []int) []int {
	i := 0
	for i < len(arr) {
		j := i + 1
		for j < len(arr) {
			arrMap := numMap[arr[i]]
			_, ok := arrMap.a[arr[j]]
			// fmt.Println(ok, arr, arr[i], arr[j], arrMap.a)
			if len(arrMap.a) > 0 {
				if !ok {
					_, b_ok := arrMap.b[arr[j]]
					if b_ok {
						arr[i], arr[j] = arr[j], arr[i]
					}
				}
			} else {
				arr[i], arr[j] = arr[j], arr[i]
			}
			j++
		}
		i++
	}
	return arr
}
