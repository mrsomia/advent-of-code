package day7

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mrsomia/advent-of-code/packages/utils"
)

func parseInput(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return lines
}

func parseLines(lines []string) [][]int {
	// Lines will come in the form [target, ...operands]
	result := [][]int{}
	for _, line := range lines {
		l := []int{}
		leftAndRight := strings.Split(line, ":")
		if len(leftAndRight) != 2 {
			fmt.Printf("Expected line length to be 2, line: %#v\n", line)
			continue
		}
		target, err := strconv.Atoi(leftAndRight[0])
		if err != nil {
			fmt.Printf("Error converting target: %v\n", leftAndRight[0])
			continue
		}
		l = append(l, target)

		operands := strings.Split(strings.TrimSpace(leftAndRight[1]), " ")
		for _, op := range operands {
			opp, err := strconv.Atoi(op)
			if err != nil {
				fmt.Printf("Err converting operand: %v\n", opp)
			}
			l = append(l, opp)
		}

		result = append(result, l)
	}
	return result
}

type Node struct {
	left  *Node
	right *Node
	value int
}

func CreateTree(line []int) *Node {
	head := &Node{
		value: line[0],
	}
	currentNodes := make([]*Node, 0)
	currentNodes = append(currentNodes, head)
	for _, n := range line[1:] {
		nextRow := make([]*Node, 0)
		for _, node := range currentNodes {
			node.left = &Node{
				value: node.value + n,
			}
			node.right = &Node{
				value: node.value * n,
			}
			nextRow = append(nextRow, node.left, node.right)
		}
		currentNodes = nextRow
	}

	return head
}

func isTargetPresent(target int, head *Node) bool {
	if target == head.value {
		return true
	}
	if head.left != nil {
		switch isTargetPresent(target, head.left) {
		case true:
			return true
		case false:
			if head.right != nil {
				return isTargetPresent(target, head.right)
			}
		}
	}
	return false
}

func SolvePartA(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseLines(parseInput(input))
	fmt.Printf("lines: %#v\n", lines)

	count := 0

	for _, line := range lines {
		target := line[0]
		treeHead := CreateTree(line[1:])
		if isTargetPresent(target, treeHead) {
			count += target
		}
	}

	return count
}

func SolvePartB(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseLines(parseInput(input))
	fmt.Printf("lines: %#v\n", lines)
	return len(lines)
}
