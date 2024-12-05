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
