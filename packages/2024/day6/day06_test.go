package day6

import "testing"

func TestPartA(t *testing.T) {
	expected := 41
	result := SolvePartA("../../../input/2024/day06.txt")
	if result != expected {
		t.Errorf("Part A failed, Expected: %v, Found: %v\n", expected, result)
	}
}

// func TestPartB(t *testing.T) {
// 	expected := 123
// 	result := SolvePartB("../../../input/2024/day06.txt")
// 	if result != expected {
// 		t.Errorf("Part B failed, Expected: %v, Found: %v\n", expected, result)
// 	}
// }
