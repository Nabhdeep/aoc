package day6

import (
	readinput "advent/readInput"
	"fmt"
	"slices"
	"strings"
)

type pos struct {
	r int
	c int
}
type guard struct {
	pos pos
	dir string
}

var visited = make(map[pos]bool)
var dirs = []string{"U", "R", "D", "L"}

var dirsMap = make(map[string][]int)

func Solve() {
	pathOfInputText := "./2024/day6/testinput.txt"
	input := readinput.ReadFile(pathOfInputText)
	sl := strings.Split(input, "\n\n")
	parsedInput, start := parseInput(sl[0])
	_currPos := pos{r: start[0], c: start[1]}
	_guard := guard{pos: _currPos, dir: "U"}
	// fmt.Println("+++")
	res1 := getUniquePositions(parsedInput, &_guard)
	res2 := findLoops(parsedInput, &_guard, start[0], start[1])
	fmt.Println(res1)
	fmt.Println(res2)

}

func parseInput(sl string) ([][]string, []int) {
	var parsedInput [][]string
	var start []int
	for i, line := range strings.Split(sl, "\n") {
		_line := strings.Split(line, "")
		parsedInput = append(parsedInput, _line)
		idx := slices.Index(_line, "^")
		if idx > 0 {
			start = append(start, i, idx)
		}
	}

	return parsedInput, start
}

func getUniquePositions(parsedInput [][]string, _guard *guard) int {
	dirsMap["U"] = []int{-1, 0}
	dirsMap["R"] = []int{0, 1}
	dirsMap["L"] = []int{0, -1}
	dirsMap["D"] = []int{1, 0}
	visited[(*_guard).pos] = true
	rows := len(parsedInput)
	cols := len(parsedInput[0])
	// fmt.Println(visited)
	for {
		// check guard position
		// update position
		// check in visited
		// if present break
		// otherwise add to visited
		_currDir := (*_guard).dir
		// fmt.Println()
		delta := dirsMap[_currDir]
		dr, dc := delta[0], delta[1]
		nr := (*_guard).pos.r + dr
		nc := (*_guard).pos.c + dc
		// fmt.Println(nr, nc, _currDir)
		// time.Sleep(200 * time.Millisecond)
		if nr >= 0 && nc >= 0 && nr < rows && nc < cols {
			if parsedInput[nr][nc] != "#" {
				(*_guard).pos.r = nr
				(*_guard).pos.c = nc
				visited[(*_guard).pos] = true
				// fmt.Println(visited)

			} else {
				// fmt.Println("CHANGE DIR")
				nDir := dirs[((slices.Index(dirs, (*_guard).dir) + 1) % 4)]
				(*_guard).dir = nDir

			}
		} else {
			break
		}
	}
	return len(visited)
}

func findLoops(parsedInput [][]string, _guard *guard, r, c int) int {

	// check for each possible positions
	// keep a obsiticle in one position
	// check if looop
	// reset the obsiticle
	var count int
	for cr, _ := range parsedInput {
		for cc, _ := range parsedInput[0] {
			var visited2 = make(map[guard]bool)
			if cr == r && cc == c || parsedInput[cr][cc] == "#" {
				continue
			}
			parsedInput[cr][cc] = "#"
			(*_guard).dir = "U"
			(*_guard).pos = pos{r: r, c: c}
			if loops(parsedInput, _guard, visited2) {
				// fmt.Println(count)
				count++
			}
			parsedInput[cr][cc] = "."

		}
	}

	return count
}

func loops(parsedInput [][]string, _guard *guard, visited2 map[guard]bool) bool {
	dirsMap["U"] = []int{-1, 0}
	dirsMap["R"] = []int{0, 1}
	dirsMap["L"] = []int{0, -1}
	dirsMap["D"] = []int{1, 0}
	rows := len(parsedInput)
	cols := len(parsedInput[0])
	// fmt.Println(visited)
	for {
		// check guard position
		// update position
		// check in visited (has the guard position and dir facing)
		// if present return true
		// otherwise add to visited
		// if out of area return false
		_currDir := (*_guard).dir
		// fmt.Println()
		delta := dirsMap[_currDir]
		dr, dc := delta[0], delta[1]
		nr := (*_guard).pos.r + dr
		nc := (*_guard).pos.c + dc
		// time.Sleep(50 * time.Millisecond)
		if nr >= 0 && nc >= 0 && nr < rows && nc < cols {
			if parsedInput[nr][nc] != "#" {
				(*_guard).pos.r = nr
				(*_guard).pos.c = nc
				// fmt.Println(nr, nc, _currDir, visited2[(*_guard)])
				if visited2[(*_guard)] {
					// for _, line := range parsedInput {
					// 	fmt.Println(line)
					// }
					// fmt.Println("=====================================")
					return true
				} else {
					visited2[(*_guard)] = true
				}

			} else {
				// fmt.Println("CHANGE DIR")
				nDir := dirs[((slices.Index(dirs, (*_guard).dir) + 1) % 4)]
				(*_guard).dir = nDir

			}
		} else {
			return false
		}
	}
}
