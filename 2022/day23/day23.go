package day23

import (
	readinput "advent/readInput"
	"fmt"
	"strings"
)

// valid_dir := [4]int{0,2,3,4}
type Elf struct {
	r       int
	c       int
	fst_dir []int
}

//	type Dir struct {
//		nr int
//		nc int
//	}
var gridSize = 30

var elves = make([]Elf, 0)
var elvesMap = make(map[string]bool)

// var posMap = make(map[string]int)
var check = [][]int{
	{1, 0},   // n 0
	{1, -1},  // nw 1
	{1, 1},   // ne 2
	{-1, 0},  // s 3
	{-1, 1},  // se 4
	{-1, -1}, // sw 5
	{0, -1},  // e 6
	{0, 1},   // w 7
}

func Solve() {
	pathOfInputText := "./2022/day23/testinput.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	parseInput(splitInput)
	// fmt.Println(elves)
	playRound(elves)
	fmt.Println(calculateEmptyGroundTiles())
}

func parseInput(input []string) {
	for r, row := range input {
		for c, ele := range row {
			if ele == '#' {
				elves = append(elves, Elf{r: r, c: c, fst_dir: []int{0, 3, 7, 6}})
				key := fmt.Sprintf("%d_%d", r, c)
				elvesMap[key] = true
			}
		}
	}
}

func playRound(elves []Elf) {
	for round := 0; round < 10; round++ {
		displayGrid()
		elves_will_move := make(map[int]bool)
		//First half consider Each 8 position
		for idx, elf := range elves {
			for _, c := range check {
				key := fmt.Sprintf("%d_%d", elf.r+c[0], elf.c+c[1])
				if elvesMap[key] {
					elves_will_move[idx] = true
				}
			}
		}

		//Propose direction
		proposeMap := make(map[string]int)

		for k, _ := range elves_will_move {
			currElf := elves[k]

			// {1, 0},   // n 0
			// {1, -1},  // nw 1
			// {1, 1},   // ne 2
			// {-1, 0},  // s 3
			// {-1, 1},  // se 4
			// {-1, -1}, // sw 5
			// {0, -1},  // e 6
			// {0, 1},   // w 7
			// 0, 3, 6, 7
			for _, toMoveIn := range currElf.fst_dir {
				switch toMoveIn {
				case 0:
					ne := getKey(currElf.r+check[2][0], currElf.c+check[2][1])
					nw := getKey(currElf.r+check[1][0], currElf.c+check[1][1])
					n := getKey(currElf.r+check[0][0], currElf.c+check[0][1])
					if !(elvesMap[n] && elvesMap[ne] && elvesMap[nw]) {
						proposeMap[n] += 1
					}
				case 3:
					sw := getKey(currElf.r+check[5][0], currElf.c+check[5][1])
					se := getKey(currElf.r+check[4][0], currElf.c+check[4][1])
					s := getKey(currElf.r+check[3][0], currElf.c+check[3][1])
					if !(elvesMap[s] && elvesMap[se] && elvesMap[sw]) {
						proposeMap[s] += 1
					}
				case 6:
					ne := getKey(currElf.r+check[2][0], currElf.c+check[2][1])
					se := getKey(currElf.r+check[4][0], currElf.c+check[4][1])
					e := getKey(currElf.r+check[6][0], currElf.c+check[6][1])
					if !(elvesMap[e] && elvesMap[ne] && elvesMap[se]) {
						proposeMap[e] += 1
					}
				case 7:
					nw := getKey(currElf.r+check[1][0], currElf.c+check[1][1])
					sw := getKey(currElf.r+check[5][0], currElf.c+check[5][1])
					w := getKey(currElf.r+check[6][0], currElf.c+check[6][1])
					if !(elvesMap[w] && elvesMap[nw] && elvesMap[sw]) {
						proposeMap[w] += 1
					}

				}
			}
		}

		for _, elf := range elves {
			// Move Each Elf in dir
			for _, dir := range elf.fst_dir {
				key := getKey(elf.r+check[dir][0], elf.c+check[dir][1])
				if proposeMap[key] == 1 {
					//Update elves map and elves arr
					delete(elvesMap, getKey(elf.r, elf.c))
					elf.r += check[dir][0]
					elf.c += check[dir][1]
					elvesMap[getKey(elf.r, elf.c)] = true
					break
				}
			}
			elf.fst_dir = append(elf.fst_dir[1:], elf.fst_dir[0])
		}
	}
}

func getKey(x, y int) string {
	return fmt.Sprintf("%d_%d", x, y)
}

func calculateEmptyGroundTiles() int {
	minR, maxR, minC, maxC := elves[0].r, elves[0].r, elves[0].c, elves[0].c
	for _, elf := range elves {
		minR = min(minR, elf.r)
		maxR = max(maxR, elf.r)
		minC = min(minC, elf.c)
		maxC = max(maxC, elf.c)
	}
	return (maxR-minR+1)*(maxC-minC+1) - len(elves)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func displayGrid() {
	for r := 0; r < gridSize; r++ {
		for c := 0; c < gridSize; c++ {
			if elvesMap[getKey(r, c)] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
