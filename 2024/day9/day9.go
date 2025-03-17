package day9

import (
	readinput "advent/readInput"
	"fmt"
	"strconv"
	"strings"
)

type attr struct {
	start  int
	length int
}

func Solve() {
	pathOfInputText := "./2024/day9/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sl := strings.Split(input, "")
	_map, _breaks := iteratePart2(&sl)
	move(&_map, _breaks)
	fmt.Println(_map)
	part2 := calculateCheckSum2(&_map)
	fmt.Println(part2)
	// mappedMemory := iterate(&sl)
	// part1 := compress(&mappedMemory)
	// fmt.Println(part1)

}
func swap(s []string, i, j int) {
	if i < 0 || j < 0 || i >= len(s) || j >= len(s) {
		fmt.Println("Index out of bounds")
		return
	}
	s[i], s[j] = s[j], s[i] // Swaps elements
}

func iterate(s *[]string) []string {
	var res []string
	for idx, ele := range *s {
		curNum, _ := strconv.Atoi(ele)

		if idx%2 == 0 {
			curID := idx / 2
			for curNum > 0 {
				res = append(res, strconv.Itoa(curID))
				curNum--
			}
		} else {
			for curNum > 0 {
				res = append(res, ".")
				curNum--
			}
		}
	}
	return res
}

func compress(memo *[]string) int {
	var i = 0
	var j = len(*memo) - 1

	for i < j {
		if (*memo)[i] != "." {
			i++
			continue
		} else {
			if (*memo)[j] != "." {
				swap((*memo), i, j)
			}
			j--
		}
	}
	return calculateCheckSum(memo)
}

func calculateCheckSum(memo *[]string) int {
	var count int
	for idx, ele := range *memo {
		curNum, _ := strconv.Atoi(ele)
		count += curNum * idx
	}
	return count
}

func iteratePart2(s *[]string) (map[int]attr, []attr) {
	var count = 0
	_map := make(map[int]attr)
	_breaks := []attr{}
	for idx, ele := range *s {
		curNum, _ := strconv.Atoi(ele)

		if idx%2 == 0 {
			curID := idx / 2
			_map[curID] = attr{start: count, length: curNum}
		} else {
			_breaks = append(_breaks, attr{start: count, length: curNum})
		}
		count += curNum
	}
	return _map, _breaks
}

func move(_map *map[int]attr, _breaks []attr) {
	for key := len(*_map) - 1; key >= 0; key-- {
		atr, _ := (*_map)[key]
		// fmt.Println(atr, key)
		for idx, _space := range _breaks {
			if _space.start < atr.start+atr.length && atr.length <= _space.length {
				(*_map)[key] = attr{start: _space.start, length: atr.length}
				if _space.length > atr.length {
					_breaks[idx] = attr{start: _space.start + atr.length, length: _space.length - atr.length}
				} else {
					_breaks = append(_breaks[:idx], _breaks[idx+1:]...)
				}

				break
			}
		}
	}
}

func calculateCheckSum2(_map *map[int]attr) int {
	var count = 0
	for key, atr := range *_map {

		for i := atr.start; i < atr.start+atr.length; i++ {
			count += key * i
		}
	}
	return count
}
