package day4

import "testing"

func TestPartA(t *testing.T) {
	expected := 18
	result := SolvePartA("../../../input/2024/day04.txt")
	if result != expected {
		t.Errorf("Part A failed, Expected: %v, Found: %v\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	expected := 9
	result := SolvePartB("../../../input/2024/day04.txt")
	if result != expected {
		t.Errorf("Part A failed, Expected: %v, Found: %v\n", expected, result)
	}
}
