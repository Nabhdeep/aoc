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
	openValves uint64
	player     int
}

func (mk memoKey) hash() string {
	return fmt.Sprintf("%s.%d.%d.%d", mk.valve, mk.time, mk.openValves, mk.player)
}

var cache = make(map[string]int)
var valvePos = make(map[string]int)

func Solve() {
	valve := make(map[string]int)
	paths := make(map[string][]string)
	openVal := make(map[string]bool)
	// cache := make(map[memoKey]int)
	pathOfInputText := "./2022/day16/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	parseInput(splitInput, &valve, &paths, &openVal, &valvePos)
	// fmt.Println(valve)
	// fmt.Println(paths)
	solvePart1(&valve, &paths, openVal)
	// fmt.Println(openVal)
}

func parseInput(input []string, valve *map[string]int, tunnles *map[string][]string, openVal *map[string]bool, valvePos *map[string]int) {
	var position int
	for _, i := range input {
		part := strings.Split(i, ";")
		vl := strings.Split(part[0], "=")
		valveName := strings.Split(vl[0], " ")
		num, _ := strconv.Atoi(vl[1])
		(*valve)[valveName[1]] = num

		if _, exists := (*valvePos)[valveName[1]]; !exists {
			(*valvePos)[valveName[1]] = position
			position++
		}
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
	initialMask := createBitmask(openValves)
	// res := findMostPressure("AA", initialMask, 30, 0, valve, paths)
	res2 := findMostPressure("AA", initialMask, 26, 1, valve, paths)
	// fmt.Println("Res ", res)
	fmt.Println("Res ", res2)
}
func createBitmask(openValves map[string]bool) uint64 {
	var mask uint64
	for valve, isOpen := range openValves {
		if isOpen {
			mask |= 1 << uint(valvePos[valve])
		}
	}
	return mask
}
func addToMask(mask uint64, valve string) uint64 {
	return mask | (1 << uint(valvePos[valve]))
}

func isValveOpen(mask uint64, valve string) bool {
	return (mask & (1 << uint(valvePos[valve]))) != 0
}
func findMostPressure(v1 string, openValve uint64, t int, player int, valve *map[string]int, paths *map[string][]string) int {
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
		openValves: openValve,
		player:     player,
	}
	// fmt.Printf("Starting Valve: %s , time: %d \n", v1, t)
	// fmt.Printf("Cache: %v\n", cache[key])
	hashKey := key.hash()
	if cache[hashKey] > 0 {
		return cache[hashKey]
	}

	maxPressure := 0
	flowRate := (*valve)[v1]
	// fmt.Printf("Valve %s flow rate: %d\n", v1, (*valve)[v1])
	if flowRate > 0 && !isValveOpen(openValve, v1) {
		// newOpenValve := make(map[string]bool)
		// for k, v := range openValve {
		// 	newOpenValve[k] = v
		// }
		newMask := addToMask(openValve, v1)
		pressureReleased := (t - 1) * flowRate
		// fmt.Printf("Opening valve %s, releasing %d pressure\n", v1, pressureReleased)
		maxPressure = pressureReleased + findMostPressure(v1, newMask, t-1, player, valve, paths)
	}

	for _, p := range (*paths)[v1] {
		maxPressure = max(maxPressure, findMostPressure(p, openValve, t-1, player, valve, paths))
		// fmt.Printf("Parsed valve: %s, flow rate: %d, tunnels: %v , maxPressure: %d\n", p, (*valve)[p], (*paths)[v1], maxPressure)
	}

	cache[hashKey] = maxPressure
	return maxPressure
}
