package day22

import (
	readinput "advent/readInput"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Solve() {
	pathOfInputText := "./2022/day22/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	board, moves := parseInput(splitInput)
	// res := moveOnBoard(board, moves)
	res2 := moveOnBoardV2(board, moves)
	fmt.Println("RES day", res2)
}

func parseInput(input []string) ([][]string, []string) {
	var board [][]string
	var moves []string
	var maxLen int
	for _, ele := range input {
		re := regexp.MustCompile(`[A-Z]`)
		r := re.Match([]byte(ele))
		if r {
			s := ""
			move := "R"
			for _, rune := range ele {
				x := fmt.Sprintf("%c", rune)
				if re.Match([]byte(x)) {
					if s == "" {
						continue
					}

					moves = append(moves, s+"-"+move)
					s = ""
					move = x
					continue
				}
				s += x
			}
			moves = append(moves, s+"-"+move)

		} else if len(ele) > 0 {
			eleArr := strings.Split(ele, "")
			maxLen = max(maxLen, len(eleArr))
			board = append(board, eleArr)
		}
	}
	board = normalizeBoard(board, maxLen)
	return board, moves
}

type Dir struct {
	nr int
	nc int
}

func normalizeBoard(board [][]string, maxLen int) [][]string {
	for idx, arr := range board {
		if len(arr) != maxLen {
			eleToFill := maxLen - len(arr)
			for eleToFill > 0 {
				arr = append(arr, " ")
				eleToFill--
			}
			board[idx] = arr
		}
	}
	return board
}

var dirMap = map[string]Dir{
	"R": {nr: 0, nc: 1},
	"L": {nr: 0, nc: -1},
	"D": {nr: 1, nc: 0},
	"U": {nr: -1, nc: 0},
}

var dirPassMap = map[string]int{
	"R": 0,
	"L": 1,
	"D": 2,
	"U": 3,
}

func moveOnBoard(board [][]string, moves []string) int {
	currPos := findStartingPos(board)
	currFacing := ""
	for _, move := range moves {
		arrMove := strings.Split(move, "-")
		steps, dir := arrMove[0], arrMove[1]
		numSteps, _ := strconv.Atoi(steps)
		if currFacing == "" {
			currFacing = dir
		} else {
			currFacing = nextFace(currFacing, dir)
		}
		i := 0
		for i < numSteps {
			nextPos := [2]int{currPos[0] + dirMap[currFacing].nr, currPos[1] + dirMap[currFacing].nc}
			//wrap around

			nextPos[0] = (nextPos[0] + len(board)) % len(board)
			nextPos[1] = ((nextPos[1] + len(board[0])) % len(board[nextPos[0]]))

			for board[nextPos[0]][nextPos[1]] == " " || board[nextPos[0]][nextPos[1]] == "" {
				nextPos[0] = (nextPos[0] + dirMap[currFacing].nr + len(board)) % len(board)
				nextPos[1] = (nextPos[1] + dirMap[currFacing].nc + len(board[0])) % len(board[nextPos[0]])
			}

			if board[nextPos[0]][nextPos[1]] == "#" {
				break
			}

			currPos = nextPos
			i++

		}

	}
	return (1000 * (currPos[0] + 1)) + (4 * (currPos[1] + 1)) + (dirPassMap[currFacing])
}

func moveOnBoardV2(board [][]string, moves []string) int {
	currPos := findStartingPos(board)
	currFacing := ""
	for _, move := range moves {
		arrMove := strings.Split(move, "-")
		steps, dir := arrMove[0], arrMove[1]
		numSteps, _ := strconv.Atoi(steps)

		if currFacing == "" {
			currFacing = dir
		} else {
			currFacing = nextFace(currFacing, dir)
		}
		saveFacing := currFacing
		for i := 0; i < numSteps; i++ {
			nextPos := [2]int{currPos[0] + dirMap[currFacing].nr, currPos[1] + dirMap[currFacing].nc}
			//wrap around

			nextPos[0], nextPos[1], currFacing = wrapCube(nextPos[0], nextPos[1], currFacing)

			fmt.Println(nextPos, currFacing, board[nextPos[0]][nextPos[1]])

			if board[nextPos[0]][nextPos[1]] == "#" {
				currFacing = saveFacing
				break
			}
			// 35289
			// 15266
			currPos = nextPos
			fmt.Println(i)

		}

	}

	fmt.Println(currPos)
	return (1000 * (currPos[0] + 1)) + (4 * (currPos[1] + 1)) + (dirPassMap[currFacing])
}

func wrapCube(nr int, nc int, currFace string) (int, int, string) {
	// fmt.Println("PREV", nr, nc, currFace)
	if nr < 0 && nc >= 50 && nc < 100 && currFace == "U" {
		nr, nc = nc+100, 0
		return nr, nc, "R"
	} else if nc < 0 && nr >= 150 && nr < 200 && currFace == "L" {
		nr, nc = 0, nr-100
		// nc = nr - 100
		// nr = 0
		return nr, nc, "D"
	} else if nr < 0 && nc >= 100 && nc < 150 && currFace == "U" {
		// nr = 199
		// nc = nc - 100
		nr, nc = 199, nc-100
		return nr, nc, "U"
	} else if nr >= 200 && nc >= 0 && nc < 50 && currFace == "D" {
		// nr = 0
		// nc = nc + 100
		nr, nc = 0, nc+100
		return nr, nc, "D"
	} else if nc >= 150 && nr >= 0 && nr < 50 && currFace == "R" {
		// nr = 149 - nr
		// nc = 99
		nr, nc = 149-nr, 99
		return nr, nc, "L"
	} else if nc == 100 && nr >= 100 && nr < 150 && currFace == "R" {
		// nr = 149 - nr
		// nc = 149
		nr, nc = 149-nr, 149
		return nr, nc, "L"
	} else if nr == 50 && nc >= 100 && nc < 150 && currFace == "D" {
		// nr = nc - 50
		// nc = 99
		nr, nc = nc-50, 99
		return nr, nc, "L"
	} else if nc == 100 && nr >= 50 && nr < 100 && currFace == "R" {
		// nr = 49
		// nc = nr + 50
		nr, nc = 49, nr+50
		return nr, nc, "U"
	} else if nr == 150 && nc >= 50 && nc < 100 && currFace == "D" {
		// nr = nc + 100
		// nc = 49
		nr, nc = nc+100, 49
		return nr, nc, "L"
	} else if nc == 50 && nr >= 150 && nr < 200 && currFace == "R" {
		// nr = 149
		// nc = nr - 100
		nr, nc = 149, nr-100
		return nr, nc, "U"
	} else if nr == 99 && nc >= 0 && nc < 50 && currFace == "U" {
		// nr = nc + 50
		// nc = 50
		nr, nc = nc+50, 50
		return nr, nc, "R"
	} else if nc == 49 && nr >= 50 && nr < 100 && currFace == "L" {
		// nr = 100
		// nc = nr - 50
		nr, nc = 100, nr-50
		return nr, nc, "D"
	} else if nc == 49 && nr >= 0 && nr < 50 && currFace == "L" {
		// nr = 149 - nr
		// nc = 0
		nr, nc = 149-nr, 0
		return nr, nc, "R"
	} else if nc < 0 && nr >= 100 && nr < 150 && currFace == "L" {
		// nr = 149 - nr
		// nc = 50
		nr, nc = 149-nr, 50
		return nr, nc, "R"
	}

	return nr, nc, currFace
}

func findStartingPos(board [][]string) [2]int {
	for j := 0; j < len(board[0]); j++ {
		if board[0][j] == "." {
			return [2]int{0, j}
		}

	}
	return [2]int{}
}

func nextFace(currFace, expectedFace string) string {
	dir := []string{"R", "D", "L", "U"}
	currIndex := 0
	for i, dir := range dir {
		if dir == currFace {
			currIndex = i
			break
		}
	}
	if expectedFace == "R" {
		return dir[(currIndex+1)%4]
	}
	return dir[(currIndex-1+4)%4]
}
