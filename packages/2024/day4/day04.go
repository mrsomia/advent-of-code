package day4

import (
	"log/slog"
	"slices"
	"strings"

	"github.com/mrsomia/advent-of-code/packages/utils"
)

func parseInput(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return lines
}

var directions = map[string][]int{
	"MU": {0, -1},
	"RU": {1, -1},
	"LU": {-1, -1},
	"LM": {-1, 0},
	"RM": {1, 0},
	"LD": {-1, 1},
	"MD": {0, 1},
	"RD": {1, 1},
}

type CrossWord struct {
	mapped [][]string
	xPoses [][]int
	aPoses [][]int
}

func NewCrossWord(lines []string) *CrossWord {
	c := &CrossWord{}
	for y, line := range lines {
		chars := strings.Split(line, "")
		c.mapped = append(c.mapped, chars)
		// get x positions
		for x, char := range chars {
			if char == "X" {
				c.xPoses = append(c.xPoses, []int{x, y})
			}
			if char == "A" {
				c.aPoses = append(c.aPoses, []int{x, y})
			}
		}
	}
	return c
}

func (c *CrossWord) IsPointValid(src, chg []int) bool {
	dst := []int{src[0] + chg[0], src[1] + chg[1]}
	return dst[1] > -1 && dst[1] < len(c.mapped) && dst[0] > -1 && dst[0] < len(c.mapped[dst[1]])
}

func (c *CrossWord) IsXmas(src, dir []int) bool {
	TARGET := "XMAS"
	current := slices.Clone(src)
	for i := 1; i < 4; i++ {
		if !c.IsPointValid(current, dir) {
			return false
		}
		current = []int{current[0] + dir[0], current[1] + dir[1]}
		char := c.mapped[current[1]][current[0]]
		if char != string(TARGET[i]) {
			return false
		}
	}
	return true
}

func (c *CrossWord) IsCrossMas(aPos []int) (ret bool) {
	crossDirs := []string{"LU", "RD", "LD", "RU"}
	otherChars := make([]string, len(crossDirs))
	defer slog.Info("IsCrossMas", "aPos", aPos, "otherChars", otherChars)
	for idx, dir := range crossDirs {
		chg := directions[dir]
		if !c.IsPointValid(aPos, chg) {
			return false
		}
		pos := []int{aPos[0] + chg[0], aPos[1] + chg[1]}
		otherChars[idx] = c.mapped[pos[1]][pos[0]]
	}
	return (((otherChars[0] == "M" && otherChars[1] == "S") || (otherChars[1] == "M" && otherChars[0] == "S")) &&
		((otherChars[2] == "M" && otherChars[3] == "S") || (otherChars[3] == "M" && otherChars[2] == "S")))
}

func (c *CrossWord) GetXmases() int {
	count := 0

	for _, xPos := range c.xPoses {
		for _, dir := range directions {
			if c.IsXmas(xPos, dir) {
				count += 1
			}
		}
	}
	return count
}

func (c *CrossWord) GetCrossMases() int {
	count := 0

	for _, aPos := range c.aPoses {
		isTrue := c.IsCrossMas(aPos)
		slog.Info("IsCrossMas", "aPos", aPos, "isTrue", isTrue)
		if isTrue {
			count += 1
		}
	}
	return count
}

func SolvePartA(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseInput(input)
	crossword := NewCrossWord(lines)
	count := crossword.GetXmases()
	return count
}

func SolvePartB(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseInput(input)
	crossword := NewCrossWord(lines)
	count := crossword.GetCrossMases()
	return count
}
