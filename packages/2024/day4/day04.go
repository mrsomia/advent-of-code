package day4

import (
	"fmt"
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

func SolvePartA(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseInput(input)
	crossword := NewCrossWord(lines)
	count := crossword.GetXmases()
	fmt.Printf("count: %#v\n", count)
	return count
}
