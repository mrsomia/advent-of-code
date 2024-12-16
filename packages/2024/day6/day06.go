package day6

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/mrsomia/advent-of-code/packages/utils"
)

func parseInput(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return lines
}

type Grid struct {
	g       [][]string
	dir     string
	tanks   [][2]int
	current [2]int
}

var directions = map[string][2]int{
	"MU": {0, -1},
	"RU": {1, -1},
	"LU": {-1, -1},
	"LM": {-1, 0},
	"RM": {1, 0},
	"LD": {-1, 1},
	"MD": {0, 1},
	"RD": {1, 1},
}

func NewGrid(lines []string) *Grid {
	grid := &Grid{}
	for y, line := range lines {
		chars := strings.Split(line, "")
		grid.g = append(grid.g, chars)
		// get x positions
		for x, char := range chars {
			switch char {
			case "#":
				grid.tanks = append(grid.tanks, [2]int{x, y})
			case "^", "v", ">", "<":
				grid.dir = char
				grid.current = [2]int{x, y}
			}
		}
	}
	return grid
}

func (grid *Grid) GetPoint(point [2]int) string {
	return grid.g[point[1]][point[0]]
}

func (grid *Grid) IsPointValid(src, chg [2]int) ([2]int, error) {
	dst := [2]int{src[0] + chg[0], src[1] + chg[1]}
	if dst[1] > -1 && dst[1] < len(grid.g) && dst[0] > -1 && dst[0] < len(grid.g[dst[1]]) {
		return dst, nil
	}
	return [2]int{-1, -1}, fmt.Errorf("Invalid Point, %#v", dst)
}

func MapDir(in string) string {
	switch in {
	case "^":
		return "MU"
	case "v":
		return "MD"
	case ">":
		return "RM"
	case "<":
		return "LM"
	}
	return ""
}

func (grid *Grid) TurnRight() {
	switch grid.dir {
	case "^":
		grid.dir = ">"
		return
	case "v":
		grid.dir = "<"
		return
	case ">":
		grid.dir = "v"
		return
	case "<":
		grid.dir = "^"
		return
	}
}

func (grid *Grid) Step() bool {
	chg := directions[MapDir(grid.dir)]
	newPoint, err := grid.IsPointValid(grid.current, chg)
	if err != nil {
		return false
	}
	if grid.GetPoint(newPoint) == "#" {
		grid.TurnRight()
		chg := directions[MapDir(grid.dir)]
		newPoint, err = grid.IsPointValid(grid.current, chg)
		if err != nil {
			return false
		}
	}
	grid.current = newPoint
	return true
}

func SolvePartA(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseInput(input)
	grid := NewGrid(lines)
	m := make(map[string]int)

	for grid.Step() {
		currStr := fmt.Sprintf("%d,%d", grid.current[0], grid.current[1])
		m[currStr] = 1
	}

	distinct := len(slices.Collect(maps.Keys(m)))

	return distinct
}

func SolvePartB(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseInput(input)
	grid := NewGrid(lines)
	return len(grid.g)
}
