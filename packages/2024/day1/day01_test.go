package day1_test

import (
	"testing"

	"github.com/mrsomia/advent-of-code/packages/2024/day1"
)

func TestPartA(t *testing.T) {
	result := day1.SolvePartA()
	expected := 11
	if result != expected {
		t.Fatalf("Expected: %v, Got: %v", expected, result)
	}
}
