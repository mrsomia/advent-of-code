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
	add      *Node
	multiply *Node
	join     *Node
	value    int
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
			node.add = &Node{
				value: node.value + n,
			}
			node.multiply = &Node{
				value: node.value * n,
			}
			joined, _ := strconv.Atoi(fmt.Sprintf("%d%d", node.value, n))
			node.join = &Node{
				value: joined,
			}
			nextRow = append(nextRow, node.add, node.multiply, node.join)
		}
		currentNodes = nextRow
	}

	return head
}

func isTargetPresent(target int, head *Node) bool {
	if head == nil {
		return false
	}
	if target == head.value && head.add == nil {
		return true
	}
	if isTargetPresent(target, head.add) {
		return true
	} else {
		return isTargetPresent(target, head.multiply)
	}
}

func isTargetPresent2(target int, head *Node) bool {
	if head == nil {
		return false
	}
	if target == head.value && head.add == nil {
		return true
	}
	addNode := isTargetPresent2(target, head.add)
	if addNode {
		return true
	}
	multiplyNode := isTargetPresent2(target, head.multiply)
	if multiplyNode {
		return true
	}

	joinNode := isTargetPresent2(target, head.join)
	return joinNode
}

func SolvePartA(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseLines(parseInput(input))

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

	count := 0

	for _, line := range lines {
		target := line[0]
		treeHead := CreateTree(line[1:])
		if isTargetPresent2(target, treeHead) {
			count += target
		}
	}

	return count
}
