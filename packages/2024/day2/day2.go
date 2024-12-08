package day2

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/mrsomia/advent-of-code/packages/utils"
)

func parseInput(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return lines
}

func SolvePartA(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseInput(input)
	valid := make(map[int]bool)
	for idx, line := range lines {
		nums := make([]int, 0)
		for _, n := range strings.Split(line, " ") {
			nn, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalf("Error converting: %v to an int", n)
			}
			nums = append(nums, nn)
		}

		isSafe := partAIsSafe(nums)
		// slog.Info("Process single report", "nums", nums, "isSafe", isSafe)
		valid[idx] = isSafe
	}

	result := 0
	for _, isVal := range valid {
		if isVal {
			result += 1
		}
	}
	return result
}

func Abs(i int) int {
	return int(math.Abs(float64(i)))
}

func partAIsSafe(report []int) (ret bool) {
	diff := 0
	last := 0
	for idx, n := range report {
		if idx == 0 {
			last = n
			continue
		}
		d := last - n
		if idx == 1 {
			diff = d
		}
		if Abs(d) > 3 || d == 0 {
			return false
		}
		// if direction changes
		if d*diff < 0 {
			return false
		}
		last = n
	}
	return true
}

// func partBIsSafeSinglePass(report []int) int {
// 	diff := 0
// 	last := 0
// 	for idx, n := range report {
// 		if idx == 0 {
// 			last = n
// 			continue
// 		}
// 		d := last - n
// 		if idx == 1 {
// 			diff = d
// 		}
// 		// if direction changes
// 		if d*diff < 0 {
// 			return idx
// 		}
// 		// if greater than 3 or 0
// 		if Abs(d) > 3 || d == 0 {
// 			return idx
// 		}
// 		last = n
// 	}
// 	return -1
// }

func PartBIsSafe(report []int) (ret bool) {
	whole := partAIsSafe(report)
	if whole {
		return true
	}
	for idx := range report {
		subReport := removeItemFromSlice(report, idx)
		isSafe := partAIsSafe(subReport)
		if isSafe {
			return true
		}
	}
	return false
}

func removeItemFromSlice(report []int, s int) []int {
	slice := make([]int, len(report))
	copy(slice, report)
	return append(slice[:s], slice[s+1:]...)
}

func SolvePartB(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseInput(input)
	valid := make(map[int]bool)
	for idx, line := range lines {
		nums := make([]int, 0)
		for _, n := range strings.Split(line, " ") {
			nn, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalf("Error converting: %v to an int", n)
			}
			nums = append(nums, nn)
		}

		isSafe := PartBIsSafe(nums)
		valid[idx] = isSafe
	}

	result := 0
	for _, isVal := range valid {
		if isVal {
			result += 1
		}
	}
	return result
}
