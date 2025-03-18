package day10

// NOTE for part 2 remove the visited map checking in the findAllPaths function

import (
	readinput "advent/readInput"
	"fmt"
	"strconv"
	"strings"
)

var moves = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

// trail 	map		0:1 		5
var trail = make(map[string]int)

func Solve() {
	pathOfInputText := "./2024/day10/testinput.txt"
	input := readinput.ReadFile(pathOfInputText)
	sl := strings.Split(input, "\n")
	grid := makeGrid(sl)
	part1 := moveOnAllPaths(grid)
	// fmt.Println(all)
	// path := findAllPaths([]int{0, 4}, 9, grid, 0, make(map[string]bool))
	fmt.Println(part1)
}

func makeGrid(sl []string) [][]int {
	var newGrid [][]int
	for _, ele := range sl {
		nums := strings.Split(ele, "")
		var lane []int
		for _, _num := range nums {
			_n, _ := strconv.Atoi(_num)
			lane = append(lane, _n)
		}
		newGrid = append(newGrid, lane)
	}
	return newGrid
}

func findAllPaths(start []int, numToSeach int, grid [][]int, currCount int, visited map[string]bool) int {
	cx, cy := start[0], start[1]
	_key := getKey(start)
	if grid[cx][cy] == numToSeach && !visited[_key] {
		visited[_key] = true
		return currCount + 1
	}

	for _, move := range moves {
		dx, dy := move[0], move[1]
		nx, ny := dx+cx, dy+cy
		if nx >= 0 && ny >= 0 && nx < len(grid[0]) && ny < len(grid) && grid[nx][ny] > grid[cx][cy] && grid[nx][ny]-grid[cx][cy] == 1 {
			// fmt.Println("currNum", grid[cx][cy], grid[nx][ny])
			currCount = findAllPaths([]int{nx, ny}, numToSeach, grid, currCount, visited)
		}
	}
	return currCount
}

func getKey(s []int) string {
	s_zero := strconv.Itoa(s[0])
	s_one := strconv.Itoa(s[1])
	return fmt.Sprintf("%s-%s", s_zero, s_one)
}

func moveOnAllPaths(grid [][]int) int {
	// trailMap := make(map[string]int)
	count := 0
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if grid[x][y] == 0 {
				_res := findAllPaths([]int{x, y}, 9, grid, 0, make(map[string]bool))
				// trailMap[getKey([]int{x, y})] = _res
				count += _res
			}
		}
	}
	// fmt.Println(trailMap)
	return count
}
