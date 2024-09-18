package day24

import (
	readinput "advent/readInput"
	"fmt"
	"regexp"
	"strings"
)

// R , L , U , D
type storm struct {
	r   int
	c   int
	dir string
}

var eMoves = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}, {0, 0}}
var stormMap = make(map[storm]bool)
var stormArr = []storm{{r: 1, c: 0, dir: "v"}, {r: -1, c: 0, dir: "^"}, {r: 0, c: -1, dir: "<"}, {r: 0, c: 1, dir: ">"}}

func Solve() {
	pathOfInputText := "./2022/day24/testinput.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	parseInput(splitInput)
	start := []int{-1, 0}
	end := []int{len(splitInput) - 1, len(splitInput[0]) - 2}
	r := iterateOnBoard(start, end, len(splitInput)-1, len(splitInput[0])-1)
	fmt.Println(r)
}

func parseInput(input []string) [][]string {
	res := [][]string{}
	for r, ele := range input {
		for c, e := range strings.Split(ele, "") {
			re := regexp.MustCompile(`(v|\^|>|<)`)
			if re.Match([]byte(e)) {
				// fmt.Println(e, r, c)
				stormMap[storm{r: r, c: c, dir: e}] = true
			}
		}
	}
	return res
}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func iterateOnBoard(start, end []int, r, c int) int {
	queue := [][]int{}
	queue = append(queue, append(start, 0))
	lcm := (r * c) / GCD(r, c)
	var visited = make(map[string]bool)
	fmt.Println(queue, start, end, r, c, lcm)
	for len(queue) > 0 {

		cr := queue[0][0]
		cc := queue[0][1]
		time := queue[0][2]
		queue = queue[1:]
		time += 1
		for _, d := range eMoves {
			nr := cr + d[0]
			nc := cc + d[1]

			if nr == end[0] && nc == end[1] {
				fmt.Println("Time -> ", time)
				return time
			}

			if (nr < 0 || nc < 0 || nr >= r || nc >= c) && !(nr == -1 && nc == 0) {
				continue
			}
			hasStorm := false
			for _, s := range stormArr {
				sr := (nr - s.r*time) % r
				sc := (nc - s.c*time) % c
				if stormMap[storm{r: sr, c: sc, dir: s.dir}] {
					hasStorm = true
					break
				}
			}
			if hasStorm {
				continue
			}
			k := fmt.Sprintf("%d_%d_%d", nr, nc, time%lcm)
			if visited[k] {
				continue
			}
			visited[k] = true
			queue = append(queue, []int{nr, nc, time})
		}
	}

	return 0
}
