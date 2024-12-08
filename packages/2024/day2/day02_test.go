package day2_test

import (
	"testing"

	"github.com/mrsomia/advent-of-code/packages/2024/day2"
)

func TestSolvePartA(t *testing.T) {
	expected := 2
	result := day2.SolvePartA("../../../input/2024/day02.txt")
	if result != expected {
		t.Errorf("Part A Failed, Expected: %v\nFound: %v\n", expected, result)
	}
}

func TestSolvePartB(t *testing.T) {
	expected := 4
	result := day2.SolvePartB("../../../input/2024/day02.txt")
	if result != expected {
		t.Errorf("Part B Failed, Expected: %v\nFound: %v\n", expected, result)
	}
}
