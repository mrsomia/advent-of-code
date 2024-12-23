package day8

import (
	"image"
	"log/slog"
	"maps"
	"slices"
	"strings"

	"github.com/mrsomia/advent-of-code/packages/utils"
)

type Grid struct {
	nodeMap map[string][]image.Point
	grid    [][]string
}

func NewGrid(lines []string) *Grid {
	grid := &Grid{
		grid:    make([][]string, 0),
		nodeMap: make(map[string][]image.Point),
	}
	for y, line := range lines {
		chars := strings.Split(line, "")
		grid.grid = append(grid.grid, chars)
		// get x positions
		for x, char := range chars {
			if char == "." {
				continue
			}
			grid.nodeMap[char] = append(grid.nodeMap[char], image.Point{x, y})
		}
	}
	return grid
}

func (grid *Grid) GetPoint(point image.Point) string {
	return grid.grid[point.Y][point.X]
}

// func (point *Point) isIn(grid *Grid) bool {
// 	return point.y < len(grid.grid) && point.x < len(grid.grid[0])
// }

func parseInput(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return lines
}

func SolvePartA(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseInput(input)
	g := NewGrid(lines)
	bounds := image.Rect(0, 0, len(g.grid[0]), len(g.grid))
	m := make(map[image.Point]bool)

	for _, antennas := range g.nodeMap {
		for i, a1 := range antennas {
			for j := i + 1; j < len(antennas); j++ {
				a2 := antennas[j]

				diff1 := a1.Sub(a2)

				if p1 := a1.Add(diff1); p1.In(bounds) {
					m[p1] = true
				}

				diff2 := a2.Sub(a1)

				if p2 := a2.Add(diff2); p2.In(bounds) {
					m[p2] = true
				}
			}
		}
	}

	return len(slices.Collect(maps.Keys(m)))
}

func SolvePartB(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseInput(input)
	g := NewGrid(lines)
	bounds := image.Rect(0, 0, len(g.grid[0]), len(g.grid))
	m := make(map[image.Point]bool)

	for freq, antennas := range g.nodeMap {
		slog.Info("freq", "freq", freq)
		for i, a1 := range antennas {
			for j := i + 1; j < len(antennas); j++ {
				a2 := antennas[j]
				slog.Info("antennas", "a1", a1, "a2", a2)
				m[a1] = true
				m[a2] = true

				diff1 := a1.Sub(a2)
				diff2 := a2.Sub(a1)
				slog.Info("diffs", "diff1", diff1, "diff2", diff2)

				for p1 := a1.Add(diff1); p1.In(bounds); p1 = p1.Add(diff1) {
					slog.Info("point1", "p1", p1)
					m[p1] = true
				}

				for p2 := a2.Add(diff2); p2.In(bounds); p2 = p2.Add(diff2) {
					slog.Info("point2", "p2", p2)
					m[p2] = true
				}
			}
		}
	}

	return len(slices.Collect(maps.Keys(m)))
}
