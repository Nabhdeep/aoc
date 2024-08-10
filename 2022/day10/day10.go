package day10

import (
	readinput "advent/readInput"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve() {
	pathOfInputText := "./2022/day10/input.txt"
	input := readinput.ReadFile(pathOfInputText)
	// fmt.Println(strings.Split(input, "\n"))
	// addxWait := []int{}
	addx := 1
	reg_X := 0
	reg_X2 := 1
	// addxWait := make(map[string]int)
	// iterateOnTape(&reg_X, addxWaitRegister, strings.Split(input, "\n"))
	iterateOnTapeV2(&reg_X, &addx, strings.Split(input, "\n"))
	iterateOnTapeV3(&reg_X2, strings.Split(input, "\n"))
	fmt.Printf("DAY 9 res1: -> %d\n", reg_X)
	// fmt.Printf("DAY 9 res2: -> %d\n", res2)
}

func iterateOnTapeV3(reg_x *int, input []string) {
	cycle := 0
	screen := [][]string{{}, {}, {}, {}, {}, {}}
	for i := 0; i < len(input); i++ {

		op := strings.Split(input[i], " ")
		// if cycle%20 == 0 || ((cycle+1)%20) == 0 || ((cycle+2)%20) == 0 {
		// 	fmt.Printf("Inner -> %d , %v\n", cycle, addxWaitRegister)

		// }
		switch op[0] {
		case "addx":
			for j := 0; j < 2; j++ {
				cycle++
				drawCycle(&cycle, reg_x, &screen)
			}

			v, _ := strconv.Atoi(op[1])
			(*reg_x) += v

		case "noop":
			cycle++
			drawCycle(&cycle, reg_x, &screen)
		}
	}
	for s := range screen {
		fmt.Println(screen[s])
	}
}

func iterateOnTapeV2(reg_x *int, addx *int, input []string) {
	cycle := 0
	for i := 0; i < len(input); i++ {

		op := strings.Split(input[i], " ")
		// if cycle%20 == 0 || ((cycle+1)%20) == 0 || ((cycle+2)%20) == 0 {
		// 	fmt.Printf("Inner -> %d , %v\n", cycle, addxWaitRegister)

		// }
		switch op[0] {
		case "addx":
			cycle += 1
			cycleCheck(&cycle, reg_x, addx)

			cycle += 1
			cycleCheck(&cycle, reg_x, addx)

			v, _ := strconv.Atoi(op[1])
			(*addx) += v

		case "noop":
			cycle++
			cycleCheck(&cycle, reg_x, addx)
		}
	}
}

func cycleCheck(cycle *int, reg_x *int, addx *int) {
	if ((*cycle)/20)%2 == 1 && ((*cycle)%20) == 0 {
		// fmt.Printf("CYCYLE NUMBER => %d , %d\n", *cycle, *addx)
		(*reg_x) += (*addx) * (*cycle)
		// fmt.Printf("CURR -> %d \n", (*reg_x))
		fmt.Println()
	}
}

// func cycleCheckV2(cycle *int, reg_x *int) {
// 	if ((*cycle)/20)%2 == 1 && ((*cycle)%20) == 0 {
// 		// fmt.Printf("CYCYLE NUMBER => %d , %d\n", *cycle, *addx)
// 		*reg_x = *reg_x * *cycle
// 		// fmt.Printf("CURR -> %d \n", (*reg_x))
// 		fmt.Println()
// 	}
// }

//	func checkStatus(addxWaitRegister *map[string]int) int {
//		var res int
//		if len((*addxWaitRegister)) == 0 {
//			return res
//		}
//		fmt.Println((*addxWaitRegister))
//		for key, value := range *addxWaitRegister {
//			parts := strings.Split(key, ":")
//			if value <= 0 {
//				op := strings.Split(parts[0], " ")
//				// Now key_parts is []string{"addx", "16"}
//				// You can access elements like this:
//				value, _ := strconv.Atoi(op[1])
//				// fmt.Println(value)
//				delete(*addxWaitRegister, key)
//				res += value
//			} else {
//				(*addxWaitRegister)[key] = (*addxWaitRegister)[key] - 1
//			}
//		}
//		// fmt.Println(res)
//		return res
//	}
func drawCycle(cycle *int, reg_x *int, screen *[][]string) {
	litPixel(cycle, reg_x, screen)
}
func litPixel(cycle *int, reg_x *int, screen *[][]string) {
	row := (*cycle - 1) / 40
	pixelPosition := (*cycle - 1) % 40
	if int(math.Abs(float64(*reg_x)-float64(pixelPosition))) < 2 {
		(*screen)[row] = append((*screen)[row], "#")
	} else {
		(*screen)[row] = append((*screen)[row], ".")
	}
}
