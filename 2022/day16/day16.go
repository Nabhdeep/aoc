package day16

import (
	readinput "advent/readInput"
	"fmt"
	"strconv"
	"strings"
)

type memoKey struct {
	valve      string
	time       int
	openValves string
}

func Solve() {
	valve := make(map[string]int)
	tunnles := make(map[string][]string)
	openVal := make(map[string]bool)
	cache := make(map[memoKey]int)
	pathOfInputText := "./2022/day16/testinput.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	parseInput(splitInput, &valve, &tunnles, &openVal)
	fmt.Println(valve)
	fmt.Println(tunnles)
	// fmt.Println(openVal)
}

func parseInput(input []string, valve *map[string]int, tunnles *map[string][]string, openVal *map[string]bool) {
	for _, i := range input {
		part := strings.Split(i, ";")
		vl := strings.Split(part[0], "=")
		valveName := strings.Split(vl[0], " ")
		num, _ := strconv.Atoi(vl[1])
		(*valve)[valveName[1]] = num

		// tunnles
		has_values := strings.Contains(part[1], "valves")
		if has_values {
			t := strings.Split(part[1], "valves")
			(*tunnles)[valveName[1]] = strings.Split(t[1], ",")
		} else {
			t := strings.Split(part[1], "valve")
			(*tunnles)[valveName[1]] = []string{t[len(t)-1]}
		}
		_, ok := (*openVal)[valveName[1]]
		if !ok {
			(*openVal)[valveName[1]] = false
		}
	}
}
func mapToString(m map[string]bool) string {
	var openValves []string
	for valve, isOpen := range m {
		if isOpen {
			openValves = append(openValves, valve)
		}
	}
	return strings.Join(openValves, ",")
}

func openValves(pos string, time int, valve *map[string]int, tunnles *map[string][]string, openVal map[string]bool, memo *map[memoKey]int, curr_max int) int {
	if time <= 0 {
		return 0
	}

	key := memoKey{
		valve:      pos,
		time:       time,
		openValves: mapToString(openVal),
	}
	v, ok := (*memo)[key]
	if ok {
		return v
	}

	// currMax := 0
	// if (*valve)[pos] > 0 && !(openVal)[pos] {
	// 	allValves := append([]string{pos}, (*tunnles)[pos]...)
	// 	for _, v := range allValves {
	// 		pressure := (*valve)[v]*time - 1
	// 		openVal[v] = true
	// 		openValves(pos, time-1, valve, tunnles, openVal, memo, pressure)
	// 		openVal[v] = false
	// 	}

	// }

}
