package day7

import "testing"

func TestPartA(t *testing.T) {
	expected := 3749
	result := SolvePartA("../../../input/2024/day07.txt")
	if result != expected {
		t.Errorf("Part A failed, Expected: %v, Found: %v\n", expected, result)
	}
}

func TestPartB(t *testing.T) {
	expected := 11387
	result := SolvePartB("../../../input/2024/day07.txt")
	if result != expected {
		t.Errorf("Part B failed, Expected: %v, Found: %v\n", expected, result)
	}
}
