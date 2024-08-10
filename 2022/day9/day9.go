package day9

import (
	readinput "advent/readInput"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve() {
	pathOfInputText := "./2022/day9/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	res1, res2 := findTailsVisited(strings.Split(input, "\n"))
	fmt.Printf("DAY 9 res1: -> %d\n", res1)
	fmt.Printf("DAY 9 res2: -> %d\n", res2)
}

func findTailsVisited(input []string) (int, int) {
	//fmt.Println(input)
	visited := make(map[string]bool)
	visited2 := make(map[string]bool)
	visited["0-0"] = true
	visited2["0-0"] = true
	currPositionHead := []int{0, 0}
	currPositionTail := []int{0, 0}
	currLongRope := make([][]int, 10)
	for i := range currLongRope {
		currLongRope[i] = []int{0, 0}
	}
	for _, s := range input {
		parts := strings.Split(s, " ")

		move := parts[0]
		steps, _ := strconv.Atoi(parts[1])

		for i := 0; i < steps; i++ {
			// fmt.Println(move, steps)
			switch move {
			case "R":
				moveInDir(0, 1, &currPositionHead, &currPositionTail, visited)
				moveInDirV2(0, 1, &currLongRope, visited2)
			case "L":
				moveInDir(0, -1, &currPositionHead, &currPositionTail, visited)
				moveInDirV2(0, -1, &currLongRope, visited2)
			case "U":
				moveInDir(1, 0, &currPositionHead, &currPositionTail, visited)
				moveInDirV2(1, 0, &currLongRope, visited2)
			case "D":
				moveInDir(-1, 0, &currPositionHead, &currPositionTail, visited)
				moveInDirV2(-1, 0, &currLongRope, visited2)
			}
			// fmt.Println(currPositionHead, currPositionTail)
		}
		// fmt.Println()
	}
	return len(visited), len(visited2)

}

func getKey(s []int) string {
	s_zero := strconv.Itoa(s[0])
	s_one := strconv.Itoa(s[1])
	return fmt.Sprintf("%s-%s", s_zero, s_one)
}

func moveInDir(dr, dc int, currPositionHead *[]int, currPositionTail *[]int, visited map[string]bool) {
	(*currPositionHead)[0] += dr
	(*currPositionHead)[1] += dc
	b1, _, _ := isTouching((*currPositionTail)[0], (*currPositionTail)[1], (*currPositionHead)[0], (*currPositionHead)[1])
	// fmt.Printf("Touching -> %d , %d\n", r1, r2)
	if !b1 {
		var nt1 int
		var nt2 int
		if (*currPositionHead)[0] == (*currPositionTail)[0] {
			nt1 = 0
		} else {
			nt1 = ((*currPositionHead)[0] - (*currPositionTail)[0]) / int(math.Abs(float64((*currPositionHead)[0])-float64((*currPositionTail)[0])))
		}
		if (*currPositionHead)[1] == (*currPositionTail)[1] {
			nt2 = 0
		} else {
			nt2 = ((*currPositionHead)[1] - (*currPositionTail)[1]) / int(math.Abs(float64((*currPositionHead)[1])-float64((*currPositionTail)[1])))
		}
		// fmt.Printf("NEXT Tail -> %d , %d\n", nt1, nt2)
		(*currPositionTail)[0] += nt1
		(*currPositionTail)[1] += nt2
	}
	if !visited[getKey(*currPositionTail)] {
		visited[getKey(*currPositionTail)] = true
	}
}

func moveInDirV2(dr, dc int, currLongRope *[][]int, visited2 map[string]bool) {
	(*currLongRope)[0][0] += dr
	(*currLongRope)[0][1] += dc
	for i := 1; i < (len(*currLongRope)); i++ {
		b1, _, _ := isTouching((*currLongRope)[i][0], (*currLongRope)[i][1], (*currLongRope)[i-1][0], (*currLongRope)[i-1][1])
		// fmt.Printf("Touching -> %d , %d\n", r1, r2)
		if !b1 {
			var nt1 int
			var nt2 int
			if (*currLongRope)[i-1][0] == (*currLongRope)[i][0] {
				nt1 = 0
			} else {
				nt1 = ((*currLongRope)[i-1][0] - (*currLongRope)[i][0]) / int(math.Abs(float64((*currLongRope)[i-1][0])-float64((*currLongRope)[i][0])))
			}
			if (*currLongRope)[i-1][1] == (*currLongRope)[i][1] {
				nt2 = 0
			} else {
				nt2 = ((*currLongRope)[i-1][1] - (*currLongRope)[i][1]) / int(math.Abs(float64((*currLongRope)[i-1][1])-float64((*currLongRope)[i][1])))
			}
			// fmt.Printf("NEXT Tail -> %d , %d\n", nt1, nt2)
			(*currLongRope)[i][0] += nt1
			(*currLongRope)[i][1] += nt2
		}
	}
	if !visited2[getKey((*currLongRope)[len((*currLongRope))-1])] {
		visited2[getKey((*currLongRope)[len((*currLongRope))-1])] = true
	}
}
func isTouching(t1, t2, h1, h2 int) (bool, int, int) {
	r1 := int(math.Abs(float64(t1) - float64(h1)))
	r2 := int(math.Abs(float64(h2) - float64(t2)))
	if r1 == 0 && r2 == 1 || r1 == 1 && r2 == 0 || r1 == 1 && r2 == 1 {
		return true, r1, r2
	}
	return false, r1, r2
}
