package day8

import "testing"

func TestPartA(t *testing.T) {
	expected := 14
	result := SolvePartA("../../../input/2024/day08.txt")
	if result != expected {
		t.Errorf("Part A failed, Expected: %v, Found: %v\n", expected, result)
	}
}

// func TestPartB(t *testing.T) {
// 	expected := 11387
// 	result := SolvePartB("../../../input/2024/day08.txt")
// 	if result != expected {
// 		t.Errorf("Part B failed, Expected: %v, Found: %v\n", expected, result)
// 	}
// }
