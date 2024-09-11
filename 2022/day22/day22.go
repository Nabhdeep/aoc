package day22

import (
	readinput "advent/readInput"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Solve() {
	pathOfInputText := "./2022/day22/testinput.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	board, moves := parseInput(splitInput)
	moveOnBoard(board, moves)
}

func parseInput(input []string) ([][]string, []string) {
	var board [][]string
	var moves []string
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
			board = append(board, eleArr)
		}
	}
	fmt.Println(moves)
	return board, moves
}

type Dir struct {
	nr int
	nc int
}

var dirMap = map[string]Dir{
	"R": {nr: 0, nc: 1},
	"L": {nr: 0, nc: -1},
	"D": {nr: 1, nc: 0},
	"U": {nr: -1, nc: 0},
}

func moveOnBoard(board [][]string, moves []string) {
	currPos := findStartingPos(board)
	currFacing := "R"
	fmt.Println("Starting POS=>>", currPos)
	for _, move := range moves {
		arrMove := strings.Split(move, "-")
		steps, dir := arrMove[0], arrMove[1]
		numSteps, _ := strconv.Atoi(steps)
		i := 0
		for i < numSteps {
			nextPos := [2]int{currPos[0] + dirMap[currFacing].nr, currPos[1] + dirMap[currFacing].nc}
			fmt.Println(nextPos, currFacing)
			//wrap around

			nextPos[0] = (nextPos[0] % len(board))
			nextPos[1] = (nextPos[1] % len(board[nextPos[0]]))

			for board[nextPos[0]][nextPos[1]] == " " || board[nextPos[0]][nextPos[1]] == "" {
				nextPos[0] += dirMap[currFacing].nr
				nextPos[1] += dirMap[currFacing].nc
			}

			if board[nextPos[0]][nextPos[1]] == "#" {
				break
			}

			currPos = nextPos

		}
		currFacing = nextFace(currFacing, dir)

	}

	fmt.Println(currPos, currFacing)
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
