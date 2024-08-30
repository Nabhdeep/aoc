package day16

import (
	readinput "advent/readInput"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type memoKey struct {
	valve      string
	time       int
	openValves string
}

var cache = make(map[memoKey]int)

func Solve() {
	valve := make(map[string]int)
	paths := make(map[string][]string)
	openVal := make(map[string]bool)
	// cache := make(map[memoKey]int)
	pathOfInputText := "./2022/day16/testinput.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	parseInput(splitInput, &valve, &paths, &openVal)
	// fmt.Println(valve)
	// fmt.Println(paths)
	solvePart1(&valve, &paths, openVal)
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
			old_values := strings.Split(t[1], ",")
			values := []string{}
			for i := 0; i < len(old_values); i++ {
				values = append(values, strings.TrimSpace(string(old_values[i])))
			}
			(*tunnles)[valveName[1]] = values
		} else {
			t := strings.Split(part[1], "valve")
			val := strings.TrimSpace(t[len(t)-1])
			(*tunnles)[valveName[1]] = []string{val}
		}
		_, ok := (*openVal)[valveName[1]]
		if !ok {
			(*openVal)[valveName[1]] = false
		}
	}
}

func solvePart1(valve *map[string]int, paths *map[string][]string, openValves map[string]bool) {
	// if im at valve v1 and i have opened the set of openValves , and i have [t] time min left
	// and there are [players] which will play after me.
	// fmt.Printf("Starting valves: %v\n", *valve)
	// fmt.Printf("Starting paths: %v\n", *paths)
	res := findMostPressure("AA", openValves, 30, 0, valve, paths)
	// res2 := findMostPressure("AA", openValves, 26, 1, valve, paths)
	fmt.Println("Res ", res)
	// fmt.Println("Res ", res2)
}
func mapToString(m map[string]bool) string {
	var openValves []string
	for valve, isOpen := range m {
		if isOpen {
			openValves = append(openValves, valve)
		}
	}
	sort.Strings(openValves)
	return strings.Join(openValves, ",")
}
func findMostPressure(v1 string, openValve map[string]bool, t int, player int, valve *map[string]int, paths *map[string][]string) int {
	// fmt.Println("=====================================================================================================================")
	// fmt.Println("All Valves: ", (*valve))
	// fmt.Printf("Starting Valve: %s , time: %d \n", v1, t)
	// fmt.Printf("openValves: %v\n", openValve)
	if t == 0 {
		if player > 0 {
			return findMostPressure("AA", openValve, 26, player-1, valve, paths)
		} else {
			return 0
		}
	}
	key := memoKey{
		valve:      v1,
		time:       t,
		openValves: mapToString(openValve),
	}
	// fmt.Printf("Starting Valve: %s , time: %d \n", v1, t)
	// fmt.Printf("Cache: %v\n", cache[key])
	if cache[key] > 0 {
		return cache[key]
	}

	maxPressure := 0
	flowRate := (*valve)[v1]
	// fmt.Printf("Valve %s flow rate: %d\n", v1, (*valve)[v1])
	if flowRate > 0 && !openValve[v1] {
		newOpenValve := make(map[string]bool)
		for k, v := range openValve {
			newOpenValve[k] = v
		}
		newOpenValve[v1] = true
		pressureReleased := (t - 1) * flowRate
		// fmt.Printf("Opening valve %s, releasing %d pressure\n", v1, pressureReleased)
		maxPressure = pressureReleased + findMostPressure(v1, newOpenValve, t-1, player, valve, paths)
	}

	for _, p := range (*paths)[v1] {
		maxPressure = max(maxPressure, findMostPressure(p, openValve, t-1, player, valve, paths))
		// fmt.Printf("Parsed valve: %s, flow rate: %d, tunnels: %v , maxPressure: %d\n", p, (*valve)[p], (*paths)[v1], maxPressure)
	}

	cache[key] = maxPressure
	return maxPressure
}
