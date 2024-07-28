package day2

import (
	readinput "advent/readInput"
	"fmt"
	"strings"
)

//PUZZLE 1
//A for Rock => X => 1
//B for Paper => Y => 2
//C for Scissors => Z => 3
// 0 is Lost
// 3 is draw
// 6 is Won

// ROCK > SCISSORS
// SCISSORS > PAPER
// PAPER > ROCK

// PUZZLE 2
// X => loose
// Y => draw
// Z => win
func Solve() {
	pathOfInputText := "./2022/day2/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sliceOfGames := strings.Split(input, "\n")
	ans1, ans2 := GetEachGame(sliceOfGames)
	fmt.Printf("Day 2 Ans for total score of rock paper scissors: %d \n", ans1)
	fmt.Printf("Day 2 Ans2 for total score of rock paper scissors: %d \n", ans2)
}

func GetEachGame(sl []string) (int, int) {
	var totalScore int
	var totalScoreV2 int
	for _, game := range sl {
		gameSlice := strings.Split(game, " ")
		currScr := getScore(gameSlice[0], gameSlice[1])
		currScrAns2 := getScoreV2(gameSlice[0], gameSlice[1])
		// fmt.Printf("GAME %s , score %d \n", game, currScr)
		totalScore += currScr
		totalScoreV2 += currScrAns2
	}
	return totalScore, totalScoreV2
}

func getScore(oppo string, us string) int {
	var shapeScore int

	switch us {
	case "X":
		shapeScore = 1
	case "Y":
		shapeScore = 2
	case "Z":
		shapeScore = 3
	}

	switch {
	case (us == "X" && oppo == "A") || (us == "Y" && oppo == "B") || (us == "Z" && oppo == "C"):
		return 3 + shapeScore
	case us == "Z" && oppo == "B" || (us == "X" && oppo == "C") || (us == "Y" && oppo == "A"):
		return 6 + shapeScore

	default:
		return shapeScore
	}
}

func getScoreV2(oppo string, res string) int {
	shapeMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
	}
	var score int
	switch res {
	//LOSE
	case "X":
		if shapeMap[oppo] == 1 {
			score += 3
		} else {
			score += shapeMap[oppo] - 1
		}
	//DRAW
	case "Y":
		score += 3 + shapeMap[oppo]
	//WIN
	case "Z":
		score += 6
		if shapeMap[oppo] == 3 {
			score += 1
		} else {
			score += shapeMap[oppo] + 1
		}
	}
	return score
}
