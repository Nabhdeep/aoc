package day5

import (
	readinput "advent/readInput"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Stack struct {
	items []string
}

func (s *Stack) Push(data string) {
	s.items = append(s.items, data)
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}
func (s *Stack) Pop() (string, error) {
	if s.isEmpty() {
		return "", errors.New("Stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}
func (s *Stack) Top() (string, error) {
	if s.isEmpty() {
		return "", errors.New("Stack is empty")
	}
	item := s.items[len(s.items)-1]
	return item, nil
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func (s *Stack) Reverse() {
	slices.Reverse(s.items)
}
func (s *Stack) PopMulti(num int) ([]string, error) {
	if s.isEmpty() {
		return nil, errors.New("Stack is empty")
	}

	items := s.items[len(s.items)-num:]
	s.items = s.items[:len(s.items)-num]
	return items, nil
}

func (s *Stack) PushMulti(num []string) {
	s.items = append(s.items, num...)
}

func Solve() {
	pathOfInputText := "./2022/day5/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sliceOfGames := strings.Split(input, "\n\n")
	getStacksAndMoves(sliceOfGames)

}

func getStacksAndMoves(sl []string) {
	stacks := sl[0]
	moves := sl[1]
	// print(moves + "\n")
	arr := make([]Stack, 9)
	arr2 := make([]Stack, 9)
	var parsed_moves [][]int
	var count int
	for idx, line := range stacks {
		if idx%4 == 1 {
			if line != 32 && !unicode.IsNumber(line) {
				arr[count].Push(string(line))
				arr2[count].Push(string(line))
			} else if unicode.IsNumber(line) {
				arr[count].Reverse()
				arr2[count].Reverse()
			}
			count++
			if count == 9 {
				count = 0
			}
		}
	}
	for _, move := range strings.Split(moves, "\n") {
		parsed_move := getMovefromLine(string(move))
		parsed_moves = append(parsed_moves, parsed_move)
	}
	// copy(arr2, arr)
	crateMover9000(parsed_moves, &arr)
	var ans1 string
	for _, i := range arr {
		res, _ := i.Top()
		ans1 += res
	}
	crateMover9001(parsed_moves, &arr2)
	var ans2 string
	for _, i := range arr2 {
		res, _ := i.Top()
		ans2 += res
	}
	fmt.Printf("Day 5 Puzzle 1 : %s \n", ans1)
	fmt.Printf("Day 5 Puzzle 2 : %s \n", ans2)
}

func getMovefromLine(line string) []int {
	var move []int
	for _, s := range strings.Split(line, " ") {
		if regexp.MustCompile(`\d`).MatchString(s) {
			val, _ := strconv.Atoi(s)
			move = append(move, val)
		}
	}
	return move
}

func crateMover9000(parsed_moves [][]int, arr *[]Stack) {
	for _, move := range parsed_moves {
		var reverse bool
		numCrates, from, to := move[0], move[1], move[2]
		if numCrates > 1 {
			reverse = true
		}
		popedItems, _ := (*arr)[from-1].PopMulti(numCrates)
		if reverse {
			slices.Reverse(popedItems)
		}
		(*arr)[to-1].PushMulti(popedItems)
	}
}

func crateMover9001(parsed_moves [][]int, arr *[]Stack) {
	for _, move := range parsed_moves {
		numCrates, from, to := move[0], move[1], move[2]
		popedItems, _ := (*arr)[from-1].PopMulti(numCrates)
		(*arr)[to-1].PushMulti(popedItems)
	}
}
