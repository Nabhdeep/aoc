package day7

import (
	readinput "advent/readInput"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type Dir struct {
	name     string
	path     string
	size     int
	children []*Dir
	parent   *Dir
	isDir    bool
}

func Solve() {
	pathOfInputText := "./2022/day7/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	ans1, ans2 := processDir(input)
	fmt.Printf("DAY 7 ANS 1: %d\n", ans1)
	fmt.Printf("DAY 7 ANS 2: %d\n", ans2)
}

func processDir(input string) (int, int) {
	currStack := []*Dir{}
	root := &Dir{}
	// currLS := []Dir{}
	for _, s := range strings.Split(input, "\n") {
		command := strings.Fields(s)
		switch {
		case string(s[0]) == "$":
			switch {
			case command[1] == "cd":
				if command[2] == ".." {
					currStack = currStack[:len(currStack)-1]
				} else if command[2] == "/" {
					root = &Dir{name: "root", path: "/", isDir: true}
					currStack = []*Dir{root}
				} else {
					topStack := currStack[len(currStack)-1]
					for _, children := range topStack.children {
						if children.name == command[2] {
							currStack = append(currStack, children)
						}
					}
				}
			}
		case command[0] == "dir":
			topStack := currStack[len(currStack)-1]
			newDir := &Dir{name: command[1], path: getCurrPath(currStack), parent: topStack, isDir: true}
			found := false
			for _, children := range topStack.children {
				if children.name == command[1] {
					found = true
					break
				}
			}
			if !found {
				topStack.children = append(topStack.children, newDir)
			}

		case unicode.IsNumber([]rune(command[0])[0]):
			topStack := currStack[len(currStack)-1]
			num, _ := strconv.Atoi(command[0])
			topStack.size = topStack.size + num
			updateParentSize(num, topStack)
			topStack.children = append(topStack.children, &Dir{name: command[1], path: getCurrPath(currStack), size: num, parent: topStack})
		}
	}
	return findDirSizeAtMost100k(root), freeUpDir(root)
}

func getCurrPath(stack []*Dir) string {
	_path := ""
	for _, s := range stack {
		_path += "/" + s.name
	}
	return _path
}

func updateParentSize(size int, dir *Dir) {
	if dir.parent == nil {
		return
	}
	dir.parent.size += size
	updateParentSize(size, dir.parent)
}

func findDirSizeAtMost100k(tree *Dir) int {
	sum := 0
	if tree.size <= 100000 {
		sum += tree.size
	}

	for i := range tree.children {
		child := tree.children[i]
		if len(child.children) > 0 {
			sum += findDirSizeAtMost100k(child)
		}
	}

	return sum
}

func freeUpDir(tree *Dir) int {
	total_unused := 70000000 - tree.size
	req := 30000000 - total_unused
	return slices.Min(findDir(tree, req, []int{}))
}

func findDir(tree *Dir, req int, arr []int) []int {
	for _, child := range tree.children {
		if child.size >= req && child.isDir {
			arr = append(arr, child.size)
		}
		arr = findDir(child, req, arr)
	}
	return arr
}
