package day1

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/mrsomia/advent-of-code/packages/utils"
)

func parseString(input string) ([]int, []int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	left := make([]int, 0)
	right := make([]int, 0)
	for idx, line := range lines {
		elems := strings.Split(line, "   ")
		n, err := strconv.Atoi(elems[0])
		if err != nil {
			log.Fatalf("On loop %v, Failed to convert to int: %v\n", idx, elems[0])
		}
		left = append(left, n)
		n, err = strconv.Atoi(elems[1])
		if err != nil {
			log.Fatalf("On loop %v, Failed to convert to int: %v\n", idx, elems[1])
		}
		right = append(right, n)
	}
	return left, right
}

func SolvePartA() int {
	f := utils.OpenFile("../../../input/2024/day01.txt")
	left, right := parseString(f)
	slices.Sort(left)
	slices.Sort(right)
	result := make([]int, 0)

	for idx, l := range left {
		r := right[idx]
		ab := int(math.Abs(float64(r - l)))
		result = append(result, ab)
	}

	t := 0

	for _, n := range result {
		t += n
	}

	return t
}
