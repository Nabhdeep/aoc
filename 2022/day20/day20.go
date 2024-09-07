package day20

import (
	readinput "advent/readInput"
	"fmt"

	//	"regexp"
	"strconv"
	"strings"
)

func Solve() {
	pathOfInputText := "./2022/day20/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	parsed_input := parseInput(splitInput)
	newArr := findGroveCoord(parsed_input, 1, 1)
	zeroPos := findZeroPos(newArr)
	s := sumAll(newArr, zeroPos)
	newArr2 := findGroveCoord(parsed_input, 10, 811589153)
	zeroPos2 := findZeroPos(newArr2)
	s2 := sumAll(newArr2, zeroPos2)
	fmt.Println(s)
	fmt.Println(s2)
}

func parseInput(input []string) [][]int {
	arr := [][]int{}

	for i, ele := range input {
		num, _ := strconv.Atoi(ele)
		arr = append(arr, []int{i, num})
	}
	return arr
}

func findGroveCoord(arr [][]int, repeat int, mul int) [][]int {
	newArr := deepCopyArr(arr, mul)
	// fmt.Println(newArr)
	r := 0
	for r < repeat {
		for i := 0; i < len(arr); i++ {
			for j, ele := range newArr {
				if i == ele[0] {
					removeEle := newArr[j]
					newArr = append(newArr[:j], newArr[j+1:]...)

					newPos := (removeEle[1] + j) % (len(newArr))
					if newPos <= 0 {
						newPos += len(newArr)
					}

					newArr = append(newArr[:newPos], append([][]int{removeEle}, newArr[newPos:]...)...)
					break
				}
			}
		}
		r++
	}

	return newArr
}

func findZeroPos(arr [][]int) int {
	zeroPos := 0
	for i, ele := range arr {
		if ele[1] == 0 {
			zeroPos = i
			break
		}
	}
	return zeroPos
}

func sumAll(newArr [][]int, zeroPos int) int {
	sum := 0
	for _, n := range []int{1000, 2000, 3000} {
		pos := (zeroPos + n) % len(newArr)
		sum += newArr[pos][1]
	}
	return sum
}

func deepCopyArr(arr [][]int, mul int) [][]int {
	newArr := make([][]int, len(arr))
	for i, ele := range arr {
		newArr[i] = make([]int, len(ele))
		ele[1] = mul * ele[1]
		copy(newArr[i], ele)
	}
	return newArr
}
