package day23

import (
	readinput "advent/readInput"
	"fmt"
	"strings"
)

type Elf struct {
	r int
	c int
}

var elvesMap = make(map[Elf]bool)
var prirorityArr = []int{0, 3, 7, 6}

var check = [][]int{
	{-1, 0},  // n 0
	{-1, -1}, // nw 1
	{-1, 1},  // ne 2
	{1, 0},   // s 3
	{1, 1},   // se 4
	{1, -1},  // sw 5
	{0, 1},   // e 6
	{0, -1},  // w 7
}

func Solve() {
	pathOfInputText := "./2022/day23/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	parseInput(splitInput)
	// fmt.Println(elves)
	playRound(&elvesMap)
	// res := playRoundV2(&elvesMap)
	fmt.Println(calculateEmptyGroundTiles(elvesMap))
}

func parseInput(input []string) {
	for r, row := range input {
		for c, ele := range row {
			if ele == '#' {
				// elves = append(elves, Elf{r: r, c: c})
				elvesMap[Elf{r: r, c: c}] = true
			}
		}
	}
}

func shouldMove(elf Elf) bool {
	for _, c := range check {
		dr, dc := c[0], c[1]
		if elvesMap[Elf{r: elf.r + dr, c: elf.c + dc}] {
			return true
		}
	}
	return false
}

func playRoundV2(elvesMap *map[Elf]bool) int {
	round := 1
	for true {
		var posCount = make(map[Elf]int)
		var elfMoveIn = make(map[Elf]Elf)

		//First half consider Each 8 position

		for currElf, _ := range *elvesMap {
			if shouldMove(currElf) {

				// {1, 0},   // n 0
				// {1, -1},  // nw 1
				// {1, 1},   // ne 2
				// {-1, 0},  // s 3
				// {-1, 1},  // se 4
				// {-1, -1}, // sw 5
				// {0, -1},  // e 6
				// {0, 1},   // w 7
				// 0, 3, 6, 7
				for _, toMoveIn := range prirorityArr {
					was := false
					switch toMoveIn {
					case 0:
						ne := Elf{r: currElf.r + check[2][0], c: currElf.c + check[2][1]}
						nw := Elf{r: currElf.r + check[1][0], c: currElf.c + check[1][1]}
						n := Elf{r: currElf.r + check[0][0], c: currElf.c + check[0][1]}
						if !(*elvesMap)[n] && !(*elvesMap)[ne] && !(*elvesMap)[nw] {
							was = true
							elfMoveIn[currElf] = n
							posCount[n] += 1
						}
					case 3:
						sw := Elf{r: currElf.r + check[5][0], c: currElf.c + check[5][1]}
						se := Elf{r: currElf.r + check[4][0], c: currElf.c + check[4][1]}
						s := Elf{r: currElf.r + check[3][0], c: currElf.c + check[3][1]}
						if !(*elvesMap)[s] && !(*elvesMap)[se] && !(*elvesMap)[sw] {
							was = true
							elfMoveIn[currElf] = s
							posCount[s] += 1
						}
					case 6:
						ne := Elf{r: currElf.r + check[2][0], c: currElf.c + check[2][1]}
						se := Elf{r: currElf.r + check[4][0], c: currElf.c + check[4][1]}
						e := Elf{r: currElf.r + check[6][0], c: currElf.c + check[6][1]}
						if !(*elvesMap)[e] && !(*elvesMap)[ne] && !(*elvesMap)[se] {
							was = true
							elfMoveIn[currElf] = e
							posCount[e] += 1
						}
					case 7:
						nw := Elf{r: currElf.r + check[1][0], c: currElf.c + check[1][1]}
						sw := Elf{r: currElf.r + check[5][0], c: currElf.c + check[5][1]}
						w := Elf{r: currElf.r + check[7][0], c: currElf.c + check[7][1]}
						if !(*elvesMap)[w] && !(*elvesMap)[nw] && !(*elvesMap)[sw] {
							was = true
							elfMoveIn[currElf] = w
							posCount[w] += 1
						}

					}
					if was {
						was = false
						break
					}
				}
			}

		}
		if len(elfMoveIn) == 0 {
			return round
		}
		for elf, nxElf := range elfMoveIn {
			// Move Each Elf in dir
			if !(*elvesMap)[nxElf] {

				// fmt.Print(copyElvesMap[key])
				if posCount[nxElf] == 1 {
					delete(*elvesMap, elf)
					(*elvesMap)[nxElf] = true
				}

			}
		}
		prirorityArr = append(prirorityArr[1:], prirorityArr[0])
		round++
	}
	return 0
}

func playRound(elvesMap *map[Elf]bool) {
	round := 0
	for round < 10 {
		var posCount = make(map[Elf]int)
		var elfMoveIn = make(map[Elf]Elf)

		//First half consider Each 8 position

		for currElf, _ := range *elvesMap {
			if shouldMove(currElf) {

				// {1, 0},   // n 0
				// {1, -1},  // nw 1
				// {1, 1},   // ne 2
				// {-1, 0},  // s 3
				// {-1, 1},  // se 4
				// {-1, -1}, // sw 5
				// {0, -1},  // e 6
				// {0, 1},   // w 7
				// 0, 3, 6, 7
				for _, toMoveIn := range prirorityArr {
					was := false
					switch toMoveIn {
					case 0:
						ne := Elf{r: currElf.r + check[2][0], c: currElf.c + check[2][1]}
						nw := Elf{r: currElf.r + check[1][0], c: currElf.c + check[1][1]}
						n := Elf{r: currElf.r + check[0][0], c: currElf.c + check[0][1]}
						if !(*elvesMap)[n] && !(*elvesMap)[ne] && !(*elvesMap)[nw] {
							was = true
							elfMoveIn[currElf] = n
							posCount[n] += 1
						}
					case 3:
						sw := Elf{r: currElf.r + check[5][0], c: currElf.c + check[5][1]}
						se := Elf{r: currElf.r + check[4][0], c: currElf.c + check[4][1]}
						s := Elf{r: currElf.r + check[3][0], c: currElf.c + check[3][1]}
						if !(*elvesMap)[s] && !(*elvesMap)[se] && !(*elvesMap)[sw] {
							was = true
							elfMoveIn[currElf] = s
							posCount[s] += 1
						}
					case 6:
						ne := Elf{r: currElf.r + check[2][0], c: currElf.c + check[2][1]}
						se := Elf{r: currElf.r + check[4][0], c: currElf.c + check[4][1]}
						e := Elf{r: currElf.r + check[6][0], c: currElf.c + check[6][1]}
						if !(*elvesMap)[e] && !(*elvesMap)[ne] && !(*elvesMap)[se] {
							was = true
							elfMoveIn[currElf] = e
							posCount[e] += 1
						}
					case 7:
						nw := Elf{r: currElf.r + check[1][0], c: currElf.c + check[1][1]}
						sw := Elf{r: currElf.r + check[5][0], c: currElf.c + check[5][1]}
						w := Elf{r: currElf.r + check[7][0], c: currElf.c + check[7][1]}
						if !(*elvesMap)[w] && !(*elvesMap)[nw] && !(*elvesMap)[sw] {
							was = true
							elfMoveIn[currElf] = w
							posCount[w] += 1
						}

					}
					if was {
						was = false
						break
					}
				}
			}

		}
		for elf, nxElf := range elfMoveIn {
			// Move Each Elf in dir
			if !(*elvesMap)[nxElf] {

				// fmt.Print(copyElvesMap[key])
				if posCount[nxElf] == 1 {
					delete(*elvesMap, elf)
					(*elvesMap)[nxElf] = true
				}

			}
		}
		prirorityArr = append(prirorityArr[1:], prirorityArr[0])
		round++
	}
}

func calculateEmptyGroundTiles(elvesMap map[Elf]bool) int {
	var minR, maxR, minC, maxC int
	first := true

	for elf := range elvesMap {
		if first {
			minR, maxR, minC, maxC = elf.r, elf.r, elf.c, elf.c
			first = false
		} else {
			minR = min(minR, elf.r)
			maxR = max(maxR, elf.r)
			minC = min(minC, elf.c)
			maxC = max(maxC, elf.c)
		}
	}

	return (maxR-minR+1)*(maxC-minC+1) - len(elvesMap)
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
