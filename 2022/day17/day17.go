package day17

import (
	readinput "advent/readInput"
	"strings"
)

type rock struct {
	typeRock string
	edges    [][]int
}

var arrRock = []rock{

	{
		typeRock: "-",
		edges:    [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
	},
	{
		typeRock: "+",
		edges:    [][]int{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}},
	},
	{
		typeRock: "J",
		edges:    [][]int{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}},
	},
	{
		typeRock: "l",
		edges:    [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
	},
	{
		typeRock: "o",
		edges:    [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
	},
}

var hugeRock = make(map[[2]int]bool)

func Solve() {
	pathOfInputText := "./2022/day17/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "")
	// fmt.Println(splitInput)
	// parseInput(splitInput)
	res := startTetris(arrRock, splitInput)
	println("Res day16: ", res)
}

// func parseInput(jetStream []string) {
// out:
// 	for _, j := range jetStream {
// 		fmt.Println(j == j)
// 		break out
// 	}
// }

func startTetris(arrRocks []rock, jetStream []string) int {
	var i int
	prev_i := -1
	var round int
	var jet int
	var maxY int
	var currRock [][]int
	rockLen := len(arrRocks)
	jetLen := len(jetStream)

	for round < 2022 {
		i = i % rockLen
		jet = jet % jetLen

		// start
		if prev_i != i {
			currRock = deepCopy(arrRocks[i].edges)
			spawnRock(&currRock, maxY)
			prev_i = i
		}
		// fmt.Println("Before jet push:")
		// printGrid(maxY, currRock)
		//JET push
		jetPush(&currRock, jetStream[jet])

		// fmt.Println("After jet push:")
		// printGrid(maxY, currRock)

		//pushDown
		push := pushDown(&currRock, &maxY)

		// fmt.Println("After pushing down:")
		// printGrid(maxY, currRock)

		// fmt.Println(currRock.typeRock, hugeRock)
		if !push {
			// fmt.Println("After pushing down:")
			// printGrid(maxY, currRock)
			i++
			round++
		}
		jet++
		// fmt.Println(round)
	}

	return maxY
}

func spawnRock(r *[][]int, maxY int) {

	for _, s := range *r {
		s[0] += maxY + 3
		s[1] += 2
	}
}
func deepCopy(edges [][]int) [][]int {
	copyEdges := make([][]int, len(edges))
	for i := range edges {
		copyEdges[i] = make([]int, len(edges[i]))
		copy(copyEdges[i], edges[i])
	}
	return copyEdges
}

func jetPush(r *[][]int, jet string) {
	switch jet {
	case ">":
		for _, s := range *r {
			if s[1]+1 > 6 || hugeRock[[2]int{s[0], s[1] + 1}] {
				return
			}
		}
		for _, s := range *r {
			s[1] += 1
		}
	case "<":
		for _, s := range *r {
			if s[1]-1 < 0 || hugeRock[[2]int{s[0], s[1] - 1}] {
				return
			}
		}
		for _, s := range *r {
			s[1] -= 1
		}
	}
}
func pushDown(r *[][]int, maxY *int) bool {

	for _, s := range *r {
		if s[0] == 0 || hugeRock[[2]int{s[0] - 1, s[1]}] {
			// Rock has hit bottom or another rock
			for _, point := range *r {
				if *maxY <= point[0] {
					*maxY = point[0] + 1
				}
				hugeRock[[2]int{point[0], point[1]}] = true
			}
			return false
		}
	}
	// Move rock down
	for i := range *r {
		(*r)[i][0]--
	}
	return true
}

// func printGrid(maxY int, currRock [][]int) {
// 	gridWidth := 7
// 	gridHeight := maxY + 4

//
// 	grid := make([][]rune, gridHeight)
// 	for i := range grid {
// 		grid[i] = make([]rune, gridWidth)
// 		for j := range grid[i] {
// 			grid[i][j] = '.'
// 		}
// 	}

//
// 	for rockPos := range hugeRock {
// 		if rockPos[0] < gridHeight && rockPos[1] < gridWidth {
// 			grid[rockPos[0]][rockPos[1]] = '#'
// 		}
// 	}

//
// 	for _, pos := range currRock {
// 		if pos[0] < gridHeight && pos[1] < gridWidth {
// 			grid[pos[0]][pos[1]] = '@'
// 		}
// 	}

//
// 	for i := gridHeight - 1; i >= 0; i-- {
// 		fmt.Println(string(grid[i]))
// 	}
// 	fmt.Println()
// }
