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
	res := moveOnBoard(board, moves)
	fmt.Println("RES day", res)
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
