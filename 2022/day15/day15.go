package day15

import (
	readinput "advent/readInput"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Sensor struct {
	x int
	y int
}
type Beacon struct {
	x int
	y int
}

func Solve() {
	pathOfInputText := "./2022/day15/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	arr := parseInput(splitInput)
	const target int = 2000000
	//2000000
	insections := getIntersections(arr, target)
	leftMost := getLeftMostX(arr)
	res := itearateOverIntersections(insections, target, leftMost)
	fmt.Println("Day 15 Res: ", int(math.Abs(float64(res[0]))+math.Abs(float64(res[1]))))
}

func parseInput(input []string) [][]interface{} {
	var arr [][]interface{}

	for _, line := range input {

		splitLine := strings.Split(line, " ")

		//		fmt.Println(splitLine)
		SensorX, _ := strconv.Atoi(strings.TrimPrefix(strings.Split(splitLine[2], ",")[0], "x="))
		SensorY, _ := strconv.Atoi(strings.TrimPrefix(strings.Split(splitLine[3], ":")[0], "y="))
		BeaconX, _ := strconv.Atoi(strings.TrimPrefix(strings.Split(splitLine[8], ",")[0], "x="))
		BeaconY, _ := strconv.Atoi(strings.TrimPrefix(splitLine[9], "y="))
		sensor := Sensor{x: SensorX, y: SensorY}
		beacon := Beacon{x: BeaconX, y: BeaconY}
		arr = append(arr, []interface{}{sensor, beacon})

	}
	return arr
}
func getIntersections(arr [][]interface{}, line int) [][]interface{} {
	// fmt.Println(arr)
	maxX := getMaxX(arr)
	println("Max is", maxX)
	var res [][]interface{}
	for _, pair := range arr {
		sensor, _ := pair[0].(Sensor)
		beacon, _ := pair[1].(Beacon)
		dist := findManhatten(sensor, beacon)
		if int(math.Abs(float64(sensor.y-line))) < dist {
			res = append(res, []interface{}{sensor, beacon})
		}
	}
	return res
}

func getMaxX(arr [][]interface{}) int {
	var currMax int
	for _, pair := range arr {
		sensor, _ := pair[0].(Sensor)
		beacon, _ := pair[1].(Beacon)
		currMax = max(sensor.x, beacon.x, currMax)
	}

	return currMax
}

func findManhatten(sensor Sensor, beacon Beacon) int {
	return int(math.Abs(float64(sensor.x)-float64(beacon.x)) + math.Abs(float64(sensor.y)-float64(beacon.y)))
}

func itearateOverIntersections(arr [][]interface{}, target int, leftMost int) [2]int {
	// fmt.Println(arr)
	// visited := make(map[int]bool)
	window := [2]int{math.MaxInt, math.MinInt}
	for _, pair := range arr {
		sensor, _ := pair[0].(Sensor)
		beacon, _ := pair[1].(Beacon)
		getLinePoints(sensor, beacon, target, leftMost, &window)
	}
	// fmt.Println(visited)
	return window
}
func getLeftMostX(arr [][]interface{}) int {
	minPoint := 0
	for _, pair := range arr {
		sensor, _ := pair[0].(Sensor)
		beacon, _ := pair[1].(Beacon)
		// fmt.Println(sensor.x, beacon.x, minPoint)
		minPoint = min(sensor.x, beacon.x, minPoint)
	}
	// fmt.Printf("LEFT MOST %d\n", minPoint)
	return minPoint
}
func getLinePoints(sensor Sensor, beacon Beacon, line int, leftMost int, window *[2]int) {
	// nt(math.Abs(float64(sensor.y - line)))
	dY := int(math.Abs(float64(sensor.y - line)))
	dist := findManhatten(sensor, beacon)
	remaning := int(math.Abs(float64(dY - dist)))
	// if dY > dist {
	// 	return
	// }
	dXl := sensor.x - remaning
	dXr := sensor.x + remaning
	if dXr < leftMost {
		dXr = leftMost
	}
	fmt.Printf("Sensor at (%d, %d) covers range %d to %d on line %d\n",
		sensor.x, sensor.y, dXl, dXr, line)

	wL := (*window)[0]
	wR := (*window)[1]
	if wL > dXl {
		(*window)[0] = dXl
	}
	if wR < dXr {
		(*window)[1] = dXr
	}
	// for i := dXr; i <= dXl; i++ {
	// 	if !(*visited)[i] {
	// 		(*visited)[i] = true
	// 	}
	// }
}
