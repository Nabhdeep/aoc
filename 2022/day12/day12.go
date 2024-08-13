package day12

import (
	readinput "advent/readInput"
	"fmt"
	"sort"
	"strings"
)

type Item struct {
	path  string
	steps int
	point Point
}
type Point struct {
	r, c int
}

type QueueCustom struct {
	items []Item
}

func (q *QueueCustom) Init() {
	sort.Slice(q.items, func(i, j int) bool {
		return q.items[i].steps < q.items[j].steps
	})
}
func (q *QueueCustom) Push(item Item) {
	q.items = append(q.items, item)
	q.Init()
}

func (q *QueueCustom) Pop() Item {
	item := q.items[0]
	q.items = q.items[1:]
	return item
}
func (q *QueueCustom) IsEmpty() bool {
	return len(q.items) == 0
}

func Solve() {
	visited := make(map[Point]bool)
	pathOfInputText := "./2022/day12/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sr, sc := getPoint(strings.Split(input, "\n"), "S")
	er, ec := getPoint(strings.Split(input, "\n"), "E")
	grid := getGrids(strings.Split(input, "\n"))
	rows := len(grid)
	cols := len(grid[0])
	start := Point{r: sr, c: sc}
	end := Point{r: er, c: ec}
	queue := QueueCustom{}
	var s string
	fmt.Println("Select '1' for Part 1 Else '2'")
	fmt.Scanln(&s)

	switch s {
	case "1":
		queue.Push(Item{point: start, steps: 0, path: "S"})
		visited[start] = true
		res1 := traverse(grid, &queue, &visited, end, rows, cols)
		fmt.Printf("DAY 12 RES1:%d \n", res1)
	case "2":
		queue.Push(Item{point: end, steps: 0, path: "E"})
		visited[end] = true
		res2 := traverseV2(grid, &queue, &visited, rows, cols)
		fmt.Printf("DAY 12 RES1:%d \n", res2)

	}
}

func traverse(grid [][]string, queue *QueueCustom, visited *map[Point]bool, end Point, rows, cols int) int {
	for !queue.IsEmpty() {
		current := queue.Pop()
		if current.point == end {
			return current.steps
		}

		for _, nighbr := range getNeighbors(current.point, rows, cols) {
			if !(*visited)[nighbr] {
				cr := grid[current.point.r][current.point.c]
				nr := grid[nighbr.r][nighbr.c]
				if getElevation(nr)-getElevation(cr) <= 1 {
					newItem := Item{
						point: nighbr,
						steps: current.steps + 1,
						path:  current.path + nr,
					}
					// current.point = nighbr
					// current.steps += 1
					// current.path = current.path + nr
					queue.Push(newItem)
					(*visited)[nighbr] = true
				}
			}
		}
	}
	return -1
}

func traverseV2(grid [][]string, queue *QueueCustom, visited *map[Point]bool, rows, cols int) int {
	for !queue.IsEmpty() {
		current := queue.Pop()
		if grid[current.point.r][current.point.c] == "a" {
			return current.steps
		}

		for _, nighbr := range getNeighbors(current.point, rows, cols) {
			if !(*visited)[nighbr] {
				cr := grid[current.point.r][current.point.c]
				nr := grid[nighbr.r][nighbr.c]
				if getElevation(cr)-1 <= getElevation(nr) {
					newItem := Item{
						point: nighbr,
						steps: current.steps + 1,
						path:  current.path + nr,
					}
					// current.point = nighbr
					// current.steps += 1
					// current.path = current.path + nr
					queue.Push(newItem)
					(*visited)[nighbr] = true
				}
			}
		}
	}
	return -1
}

func getPoint(input []string, searchSymbol string) (int, int) {
	for i := 0; i < len(input); i++ {
		newStr := strings.Split(input[i], "")
		for j := 0; j < len(newStr); j++ {
			if newStr[j] == searchSymbol {
				fmt.Println(newStr[j])
				return i, j
			}
		}
	}
	return 0, 0
}

func getGrids(input []string) [][]string {
	res := [][]string{}
	for i := 0; i < len(input); i++ {
		newStr := strings.Split(input[i], "")
		res = append(res, newStr)
	}
	return res
}

func getElevation(r string) rune {
	switch r {
	case "S":
		return []rune("a")[0]
	case "E":
		return []rune("z")[0]
	default:
		return []rune(r)[0]
	}
}
func getNeighbors(point Point, rows, cols int) []Point {
	paths := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	res := make([]Point, 0, 4)
	for _, p := range paths {
		dR := point.r + p[0]
		dC := point.c + p[1]
		if dR >= 0 && dR < rows && dC >= 0 && dC < cols {
			res = append(res, Point{r: dR, c: dC})
		}
	}
	return res
}
