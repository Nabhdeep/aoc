package day14

import (
	readinput "advent/readInput"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func Solve() {
	allPoints := make(map[Point]bool)
	allPointsV2 := make(map[Point]bool)
	pathOfInputText := "./2022/day14/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	parseInput(strings.Split(input, "\n"), &allPoints)
	parseInput(strings.Split(input, "\n"), &allPointsV2)
	// rocks := len(allPoints)
	res := simulateSand(&allPoints)
	res2 := simulateSandV2(&allPointsV2)

	fmt.Printf("Day 14 Res: %d \n", res)
	fmt.Printf("Day 14 Res2: %d \n", res2)
}

func parseInput(sl []string, allPoints *map[Point]bool) {
	for _, line := range sl {
		points := strings.Split(line, " -> ")
		addPoint(points, allPoints)
		// fmt.Println((*allPoints))
	}

}

func simulateSand(allPoints *map[Point]bool) int {
	sandStart := Point{x: 500, y: 0}
	currSand := sandStart
	countSand := 0
	maxDepth := getMaxDepth(allPoints)
	for {
		if currSand.y >= maxDepth {
			// fmt.Printf("Sand fell into abyss at %v\n", currSand)
			return countSand
		}

		//Move
		nx := currSand.x
		ny := currSand.y

		//CHECK
		// |
		// V
		down := Point{x: nx, y: ny + 1}
		left := Point{x: nx - 1, y: ny + 1}
		right := Point{x: nx + 1, y: ny + 1}
		isLeftPreset := (*allPoints)[left]
		isRightPresent := (*allPoints)[right]
		isDownPresent := (*allPoints)[down]
		if !isDownPresent {
			currSand = down
		} else if !isLeftPreset {
			currSand = left
			// fmt.Printf("Sand moving left to %v\n", currSand)
		} else if !isRightPresent {
			currSand = right
			// fmt.Printf("Sand moving right to %v\n", currSand)
		} else {

			(*allPoints)[currSand] = true
			currSand = sandStart
			countSand++
			// fmt.Printf("Sand came to rest at %v\n", currSand)
		}
	}

}

func simulateSandV2(allPoints *map[Point]bool) int {
	sandStart := Point{x: 500, y: 0}
	currSand := sandStart
	countSand := 0
	maxDepth := getMaxDepth(allPoints) + 2
	for {
		if (*allPoints)[sandStart] {
			break
		}
		if currSand.y == maxDepth {
			restPoint := Point{x: currSand.x, y: maxDepth}
			(*allPoints)[restPoint] = true
			currSand = sandStart
			continue
		}
		//Move
		nx := currSand.x
		ny := currSand.y
		//CHECK
		// |
		// V
		down := Point{x: nx, y: ny + 1}
		left := Point{x: nx - 1, y: ny + 1}
		right := Point{x: nx + 1, y: ny + 1}
		isLeftPreset := (*allPoints)[left]
		isRightPresent := (*allPoints)[right]
		isDownPresent := (*allPoints)[down]
		if !isDownPresent {
			currSand = down
		} else if !isLeftPreset {
			currSand = left
			// fmt.Printf("Sand moving left to %v\n", currSand)
		} else if !isRightPresent {
			currSand = right
			// fmt.Printf("Sand moving right to %v\n", currSand)
		} else {
			(*allPoints)[currSand] = true
			// fmt.Printf("Sand came to rest at %v\n", currSand)
			currSand = sandStart
			countSand++
		}

	}
	return countSand

}

func getMaxDepth(allPoints *map[Point]bool) int {
	res := 0
	for k, _ := range *allPoints {
		res = int(math.Max(float64(res), float64(k.y)))
	}
	return res
}

// func addPoint(points []string, allPoints *map[Point]bool) {
// 	prev_x := 0
// 	prev_y := 0
// 	for _, p := range points {
// 		println(p)
// 		xy := strings.Split(p, ",")
// 		x, _ := strconv.Atoi(xy[0])
// 		y, _ := strconv.Atoi(xy[1])
// 		if prev_x == 0 && prev_y == 0 {
// 			prev_x = x
// 			prev_y = y
// 			(*allPoints)[Point{x: x, y: y}] = true
// 			continue
// 		}
// 		dX := prev_x - x
// 		dY := prev_y - y
// 		if dY > 0 {
// 			i := 0
// 			for i <= dY {
// 				(*allPoints)[Point{x: x, y: prev_y + i}] = true
// 				i++
// 			}
// 		}
// 		if dY < 0 {
// 			i := 0
// 			for i >= dY {
// 				(*allPoints)[Point{x: x, y: prev_y - i}] = true
// 				i--
// 			}
// 		}

// 		if dX > 0 {
// 			i := 0
// 			for i <= dX {
// 				(*allPoints)[Point{x: prev_x - i, y: y}] = true
// 				i++
// 			}
// 		}
// 		if dX < 0 {
// 			i := 0
// 			for i >= dX {
// 				(*allPoints)[Point{x: prev_x + i, y: y}] = true
// 				i--
// 			}
// 		}

// 		prev_x = x
// 		prev_y = y
// 	}
// }

func addPoint(points []string, allPoints *map[Point]bool) {
	prev_x, prev_y := 0, 0
	for _, p := range points {
		xy := strings.Split(p, ",")
		x, _ := strconv.Atoi(xy[0])

		y, _ := strconv.Atoi(xy[1])

		if prev_x == 0 && prev_y == 0 {
			prev_x, prev_y = x, y
			(*allPoints)[Point{x: x, y: y}] = true
			continue
		}
		dx, dy := x-prev_x, y-prev_y
		length := int(math.Max(math.Abs(float64(dx)), math.Abs(float64(dy))))
		for i := 0; i <= length; i++ {
			(*allPoints)[Point{x: prev_x + dx*i/length, y: prev_y + dy*i/length}] = true
		}
		prev_x, prev_y = x, y
	}
}
