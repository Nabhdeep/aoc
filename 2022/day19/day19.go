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
	geode               [2]int
	maxObsidianRequired int
	maxOreRequired      int
	maxClayRequired     int
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

func (s state) hash() string {
	return fmt.Sprintf("%d:%d:%d:%d:%d:%d:%d:%d", s.ore_robot, s.clay_robot, s.obsidian_robot, s.geode_robot, s.ore, s.clay, s.obsidian, s.geode)
}

func (s state) new() state {
	return state{ore_robot: 1, clay_robot: 0, obsidian_robot: 0, geode_robot: 0, ore: 0, clay: 0, obsidian: 0, geode: 0}

}

func (s state) buildOreRobot(bp blueprint) state {
	newState := s
	newState.ore -= bp.ore
	newState = newState.canHarvestResource()
	newState.ore_robot++
	//newState.minute--
	return newState
}

func (s state) buildClayRobot(bp blueprint) state {
	newState := s
	newState.ore -= bp.clay
	newState = newState.canHarvestResource()
	newState.clay_robot++
	return newState
}

func (s state) buildObsidianRobot(bp blueprint) state {
	newState := s
	newState.ore -= bp.obsidian[0]
	newState.clay -= bp.obsidian[1]
	newState = newState.canHarvestResource()
	newState.obsidian_robot++
	//newState.minute--
	return newState
}

func (s state) buildGeodeRobot(bp blueprint) state {
	newState := s
	newState.ore -= bp.geode[0]
	newState.obsidian -= bp.geode[1]
	newState = newState.canHarvestResource()
	newState.geode_robot++
	//newState.minute--
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
var cache map[string]int

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
		// blueprints = append(blueprints, blueprint{ore: atoi(record[0]), clay: atoi(record[1]), obsidian: [2]int{atoi(record[2]), atoi(record[3])}, geode: [2]int{atoi(record[4]), atoi(record[5])}})
		blueprints = append(blueprints,
			blueprint{
				ore:                 atoi(record[0]),
				clay:                atoi(record[1]),
				obsidian:            [2]int{atoi(record[2]), atoi(record[3])},
				geode:               [2]int{atoi(record[4]), atoi(record[5])},
				maxOreRequired:      max(atoi(record[0]), atoi(record[1]), atoi(record[2]), atoi(record[4])),
				maxClayRequired:     atoi(record[3]),
				maxObsidianRequired: atoi(record[5])})
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
	totalQualityLevel := 0
	for i, bp := range bps {
		cache = make(map[string]int)
		maxGeodes := dfs(state{}.new(), 24, bp)
		qualityLevel := (i + 1) * maxGeodes
		totalQualityLevel += qualityLevel
		fmt.Printf("Blueprint %d: Max Geodes = %d, Quality Level = %d\n", i+1, maxGeodes, qualityLevel)
	}
	fmt.Printf("Total Quality Level: %d\n", totalQualityLevel)
}

func dfs(s state, t int, bp blueprint) int {
	if t <= 0 {
		return s.geode
	}

	key := fmt.Sprintf("%s:%d", s.hash(), t)
	if num, exist := cache[key]; exist {
		return num
	}

	maxGeode := s.geode + s.geode_robot*t
	// Try building each type of robot
	if s.ore >= bp.geode[0] && s.obsidian >= bp.geode[1] {
		newState := s.buildGeodeRobot(bp)
		maxGeode = max(maxGeode, dfs(newState, t-1, bp))
	} else {
		if s.ore >= bp.obsidian[0] && s.clay >= bp.obsidian[1] && s.obsidian_robot < bp.maxObsidianRequired {
			newState := s.buildObsidianRobot(bp)
			maxGeode = max(maxGeode, dfs(newState, t-1, bp))
		}
		if s.ore >= bp.clay && s.clay_robot < bp.maxClayRequired {
			newState := s.buildClayRobot(bp)
			maxGeode = max(maxGeode, dfs(newState, t-1, bp))
		}
		if s.ore >= bp.ore && s.ore_robot < bp.maxOreRequired {
			newState := s.buildOreRobot(bp)
			maxGeode = max(maxGeode, dfs(newState, t-1, bp))
		}

		maxGeode = max(maxGeode, dfs(s.canHarvestResource(), t-1, bp))
	}

	cache[key] = maxGeode
	return maxGeode
}
func max(a int, b ...int) int {
	result := a
	for _, v := range b {
		if v > result {
			result = v
		}
	}
	return result
}
