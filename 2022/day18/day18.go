package day18

import (
	readinput "advent/readInput"
	"fmt"
	"strconv"
	"strings"
)

type cube struct {
	x, y, z int
}

var exposedSides []int

var mapOfCube = make(map[cube]bool)
var delta = [][]int{{0, 0, 1}, {0, 0, -1}, {1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}}

func Solve() {
	pathOfInputText := "./2022/day18/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	// println(splitInput)
	parseInput(splitInput)
	res := findExposedSides()
	res2 := findExposedSidesV2()
	fmt.Println(res)
	fmt.Println(res2)
}

func parseInput(input []string) {
	for _, i := range input {
		coords := strings.Split(i, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		new_cube := cube{x, y, z}
		mapOfCube[new_cube] = true
	}

	exposedSides = make([]int, len(mapOfCube))
}

func findExposedSides() int {
	//making 6 imaginary cubes to check if those are present in the exposedSides if no then add
	idx := 0
	has_cube := 0
	for _cube, _ := range mapOfCube {
		for _, next := range delta {
			nX := _cube.x + next[0]
			nY := _cube.y + next[1]
			nZ := _cube.z + next[2]
			if mapOfCube[cube{x: nX, y: nY, z: nZ}] {
				has_cube++
			}
		}
		exposedSides[idx] += has_cube
		has_cube = 0
	}
	return len(mapOfCube)*6 - sumWithForLoop(exposedSides)
}

func findExposedSidesV2() int {
	// added dfs for each imaginary cube
	// dfs will do a floodfill and when it exceedes the min or max limit will return true

	ans := 0
	for _cube, _ := range mapOfCube {
		for _, next := range delta {
			nX := _cube.x + next[0]
			nY := _cube.y + next[1]
			nZ := _cube.z + next[2]
			var visited = make(map[cube]bool)
			if dfs(cube{x: nX, y: nY, z: nZ}, visited) {
				ans++
			}

		}
	}
	return ans
}

func dfs(_cube cube, visited map[cube]bool) bool {
	if _cube.x >= 30 || _cube.y >= 30 || _cube.z >= 30 || _cube.x < 0 || _cube.y < 0 || _cube.z < 0 {
		return true
	}
	if mapOfCube[_cube] || visited[_cube] {
		return false
	}
	visited[_cube] = true

	for _, next := range delta {
		nX := _cube.x + next[0]
		nY := _cube.y + next[1]
		nZ := _cube.z + next[2]

		if dfs(cube{nX, nY, nZ}, visited) {
			return true
		}
	}
	return false
}

func sumWithForLoop(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
