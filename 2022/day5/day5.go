package day5

import (
	readinput "advent/readInput"
	"errors"
	"fmt"
	"slices"
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
	if len(s.items) == 0 {
		return true
	}
	return false
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

func Solve() {
	pathOfInputText := "./2022/day5/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	sliceOfGames := strings.Split(input, "\n\n")
	getStacksAndMoves(sliceOfGames)

}

func getStacksAndMoves(sl []string) {
	stacks := sl[0]
	// moves := sl[1]
	print(stacks + "\n")
	arr := make([]Stack, 9)
	var count int
	for idx, line := range stacks {
		if idx%4 == 1 {
			if line != 32 && !unicode.IsNumber(line) {
				arr[count].Push(string(line))
			} else if unicode.IsNumber(line) {
				arr[count].Reverse()
			}
			count++
			if count == 9 {
				count = 0
			}
		}
	}
	// fmt.Print(arr)
	fmt.Print(arr[0].Top())
}
