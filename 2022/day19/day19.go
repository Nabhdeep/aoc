package day19

import (
	readinput "advent/readInput"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type blueprint struct {
	ore  int
	clay int
	//obsidian -> ore , clay
	obsidian [2]int
	//geode -> ore , obsidian
	geode [2]int
}

type state struct {
	//minute         int
	ore_robot      int
	clay_robot     int
	obsidian_robot int
	geode_robot    int
	ore            int
	clay           int
	obsidian       int
	geode          int
}

func (s state) new() state {
	return state{0, 1, 0, 0, 0, 0, 0, 0}
}

func (s state) buildOreRobot(bp blueprint) state {
	newState := s
	if bp.ore <= newState.ore {
		newState.ore_robot++
		newState.ore -= bp.ore
		//newState.minute--
	}
	return newState
}

func (s state) buildClayRobot(bp blueprint) state {
	newState := s
	if bp.clay <= newState.ore {
		newState.clay_robot++
		newState.ore -= bp.ore
		//newState.minute--
	}
	return newState
}

func (s state) buildObsidianRobot(bp blueprint) state {
	newState := s
	if bp.obsidian[0] <= newState.ore && bp.obsidian[1] <= newState.clay {
		newState.obsidian_robot++
		newState.ore -= bp.obsidian[0]
		newState.clay -= bp.obsidian[1]
		//newState.minute--
	}
	return newState
}

func (s state) buildGeodeRobot(bp blueprint) state {
	newState := s
	if bp.geode[0] <= newState.ore && bp.geode[1] <= newState.obsidian {
		newState.geode_robot++
		newState.ore -= bp.geode[0]
		newState.obsidian -= bp.geode[1]
		//newState.minute--
	}
	return newState
}

func (s state) canHarvestResource() state {
	newState := s
	if newState.geode_robot > 0 {
		newState.geode += newState.geode_robot
	}

	if newState.obsidian_robot > 0 {
		newState.obsidian += newState.obsidian_robot
	}

	if newState.clay_robot > 0 {
		newState.clay += newState.clay_robot
	}
	if newState.ore_robot > 0 {
		newState.ore += newState.ore_robot
	}
	return newState
}

var blueprints []blueprint

// state -> num of geode
var cache map[state]int

func Solve() {
	pathOfInputText := "./2022/day19/testinput.txt"
	input := readinput.ReadFile(pathOfInputText)
	splitInput := strings.Split(input, "\n")
	parseInput(splitInput)
	startFactory(blueprints)
}

func parseInput(input []string) {
	for _, i := range input {
		parts := strings.Split(i, ":")
		bpTxt := parts[1]
		re := regexp.MustCompile(`\d{1,2}`)
		record := re.FindAllString(bpTxt, -1)
		blueprints = append(blueprints, blueprint{ore: atoi(record[0]), clay: atoi(record[1]), obsidian: [2]int{atoi(record[2]), atoi(record[3])}, geode: [2]int{atoi(record[4]), atoi(record[5])}})
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func startFactory(bps []blueprint) {
	for _, bp := range bps {
		gotMaxgeode := dfs(state{}.new(), 24, bp)
		fmt.Println(gotMaxgeode)
	}

}

func dfs(s state, t int, bp blueprint) int {
	if t == 0 {
		return s.geode
	}

	if num, exist := cache[s]; exist {
		return num
	}

	maxGeode := 0
	fmt.Println(s)
	// MAKE ROBOT
	//HARVEST
	cache[s] = maxGeode

	return maxGeode
}
