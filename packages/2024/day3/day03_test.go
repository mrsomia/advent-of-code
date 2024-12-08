package day3

import "testing"

func TestPartA(t *testing.T) {
	expected := 161
	result := SolvePartA("../../../input/2024/day03.txt")
	if result != expected {
		t.Errorf("Part A failed, Expected: %v, Found: %v\n", expected, result)
	}
}