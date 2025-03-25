package day12

import (
	readinput "advent/readInput"
	"fmt"
	"strconv"
	"strings"
)

var moves = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type Dim struct {
	area      int
	parameter int
	sides     int
}

func (dim *Dim) measure() int {
	return dim.area * dim.parameter
}
func (dim *Dim) bulkDiscount() int {
	return dim.area * dim.sides
}

func Solve() {
	pathOfInputText := "./2024/day12/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sl := strings.Split(input, "\n")
	grid := makeGrid(&sl)
	// measurePlot(grid, "I", []int{0, 4})
	part1, part2 := iterateOnGrid(grid)
	fmt.Println(part1, part2)
}

func makeGrid(sl *[]string) [][]string {
	var result [][]string
	for _, line := range *sl {
		var _line []string
		for _, ele := range strings.Split(line, "") {
			_line = append(_line, ele)
		}
		result = append(result, _line)
	}
	return result
}
func getKey(coor []int) string {
	return fmt.Sprintf("%d_%d", coor[0], coor[1])
}

func getKey2(coor []float32) string {
	return fmt.Sprintf("%.1f_%.1f", coor[0], coor[1])
}

func calculateParameter(coor []int, grid [][]string) int {
	parameter := 0
	pMoves := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for _, m := range pMoves {
		nx := m[0] + coor[0]
		ny := m[1] + coor[1]

		if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || grid[nx][ny] != grid[coor[0]][coor[1]] {
			parameter += 1
		}
	}
	return parameter
}

// func calculateSides (coor []int ,)

func measurePlot(grid [][]string, plot string, start []int, visited_parent map[string]bool) Dim {
	queue := [][]int{start}
	visited := make(map[string]bool)
	visited_parent[getKey(start)] = true
	dimensions := Dim{area: 1, parameter: calculateParameter(start, grid)}

	for len(queue) > 0 {
		curr := queue[0]
		cx, cy := curr[0], curr[1]
		queue = queue[1:]
		_key := getKey([]int{cx, cy})
		if !visited[_key] {
			for _, m := range moves {
				nx := m[0] + cx
				ny := m[1] + cy
				newKey := getKey([]int{nx, ny})
				if nx >= 0 && ny >= 0 && nx < len(grid) && ny < len(grid[0]) && grid[nx][ny] == plot {
					if !visited_parent[newKey] {
						dimensions.area += 1
						dimensions.parameter += calculateParameter([]int{nx, ny}, grid)
						queue = append(queue, []int{nx, ny})
						visited_parent[newKey] = true
					}
				}
			}
			visited[_key] = true
		}
	}
	// fmt.Println(visited)
	findSides(visited, &dimensions)
	return dimensions
}

func getMapKeys(_map map[string]bool) [][]int {
	keys := make([][]int, len(_map))

	i := 0
	for k := range _map {
		nums := strings.Split(k, "_")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		keys[i] = []int{x, y}
		i++
	}

	return keys
}

func getMapKeys2(_map map[string]bool) [][]float32 {
	keys := make([][]float32, len(_map))

	i := 0
	for k := range _map {
		nums := strings.Split(k, "_")
		x64, _ := strconv.ParseFloat(nums[0], 32)
		y64, _ := strconv.ParseFloat(nums[1], 32)
		x := float32(x64)
		y := float32(y64)

		keys[i] = []float32{x, y}
		i++
	}

	return keys
}
func getCadidates(r, c int) [][]float32 {
	return [][]float32{{float32(r) - 0.5, float32(c) - 0.5}, {float32(r) - 0.5, float32(c) + 0.5}, {float32(r) + 0.5, float32(c) + 0.5}, {float32(r) + 0.5, float32(c) - 0.5}}
}
func slicesEqual(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

/*
For part 2

 1. get all the regions

 2. consider a cartisian plane
    for example 0,0 is a point of plot
    now to calculare all the corners we can consider
    (-+0.5 , -+0.5) the distance from the point
    so plot 0,0 will have 4 points / corners

    -0.5 , -0.5 left up point
    -0.5 , 0.5	down left
    0.5, -0.5 	up right
    0.5,0.5 	down right

    each region will have 4 points

3. now using the points we can get the corresponding region with it
4. we can make arr of all 4 sides of the point region up down left and right
5. now if a point has 1 region with which means it has one side , 2 region means no side ,3 region means one side , 4 is in the middle of the 4 regions
*/
func findSides(regionMap map[string]bool, dimensions *Dim) {
	corner_candidates := make(map[string]bool)
	_keysMap := getMapKeys(regionMap)
	for _, ele := range _keysMap {
		r := ele[0]
		c := ele[1]

		for _, corner := range getCadidates(r, c) {
			corner_candidates[getKey2(corner)] = true
		}
	}
	// fmt.Println(corner_candidates)
	_c := getMapKeys2(corner_candidates)
	// fmt.Println(_c)
	for _, ele := range _c {
		cr := ele[0]
		cc := ele[1]

		corner := [][]int{{int(cr - 0.5), int(cc - 0.5)}, {int(cr - 0.5), int(cc + 0.5)}, {int(cr + 0.5), int(cc + 0.5)}, {int(cr + 0.5), int(cc - 0.5)}}
		// fmt.Println(ele, corner)
		_config := []bool{}
		for _, corner_ele := range corner {
			nr := corner_ele[0]
			nc := corner_ele[1]
			// _k := regionMap[getKey([]int{nr, nc})]
			// fmt.Println(regionMap)
			// fmt.Println(_k, getKey([]int{nr, nc}), nr, nc)
			_config = append(_config, regionMap[getKey([]int{nr, nc})])
		}
		// fmt.Println(_config)
		num := 0
		_sum_corner := 0
		for _, b := range _config {
			if b {
				num++
			}
		}

		if num == 1 {
			_sum_corner++
		} else if num == 2 {
			if slicesEqual(_config, []bool{true, false, true, false}) || slicesEqual(_config, []bool{false, true, false, true}) {
				_sum_corner += 2
			}
		} else if num == 3 {
			_sum_corner++
		}
		dimensions.sides += _sum_corner
	}
}

func iterateOnGrid(grid [][]string) (int, int) {
	total := 0
	total2 := 0
	var visited_parent = make(map[string]bool)
	for x, x_coor := range grid {
		for y, ele := range x_coor {
			if !visited_parent[getKey([]int{x, y})] {
				dim := measurePlot(grid, ele, []int{x, y}, visited_parent)
				// fmt.Println("For", ele)
				// fmt.Println(dim.sides)
				total += dim.measure()
				total2 += dim.bulkDiscount()
			}
		}
	}
	return total, total2
}
