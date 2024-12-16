package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
	"github.com/mrsomia/advent-of-code/packages/utils"
)

func parseInput(input string) []string {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	return sections
}

type Rule struct {
	before int
	after  int
}

func ParseSection1(in string) []*Rule {
	result := []*Rule{}
	lines := strings.Split(strings.TrimSpace(in), "\n")
	for _, line := range lines {
		nums := strings.Split(line, "|")
		l, err := strconv.Atoi(strings.TrimSpace(nums[0]))
		if err != nil {
			fmt.Printf("Error parsing number: %v, err: %v\n", nums[0], err)
			return nil
		}
		r, err := strconv.Atoi(strings.TrimSpace(nums[1]))
		if err != nil {
			fmt.Printf("Error parsing number: %v, err: %v\n", nums[1], err)
			return nil
		}
		result = append(result, &Rule{l, r})
	}
	return result
}

func ParseSection2(in string) [][]int {
	result := make([][]int, 0)
	lines := strings.Split(strings.TrimSpace(in), "\n")
	for _, line := range lines {
		nums := strings.Split(line, ",")
		lineSlice := make([]int, 0)

		for _, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Printf("Error parsing num: %v, err:%v\n", num, err)
			}
			lineSlice = append(lineSlice, n)
		}
		result = append(result, lineSlice)
	}
	return result
}

func createGraph(rules []*Rule, line []int) graph.Graph[int, int] {
	g := graph.New(graph.IntHash, graph.Directed(), graph.Acyclic(), graph.PreventCycles())
	for _, r := range rules {
		if slices.Contains(line, r.before) && slices.Contains(line, r.after) {
			g.AddVertex(r.before)
			g.AddVertex(r.after)
			g.AddEdge(r.before, r.after)
		}
	}
	return g
}

func SolvePartA(filename string) int {
	input := utils.OpenFile(filename)
	sections := parseInput(input)
	rules := ParseSection1(sections[0])
	lines := ParseSection2(sections[1])

	correctLines := [][]int{}
	for _, line := range lines {
		g := createGraph(rules, line)
		sorted, err := graph.TopologicalSort(g)
		if err != nil {
			fmt.Printf("Error sorting line: %#v, err: %v\n", line, err)
		}

		compValue := slices.Compare(line, sorted)
		if compValue == 0 {
			// fmt.Printf("Correct, line: %v, idx: %v\n", line, idx)
			correctLines = append(correctLines, line)
		}
	}

	result := 0
	for _, line := range correctLines {
		idx := len(line) / 2
		result += line[idx]
		// fmt.Printf("idx: %v, line: %v, result: %v\n", idx, line, result)
	}
	return result
}

func SolvePartB(filename string) int {
	input := utils.OpenFile(filename)
	sections := parseInput(input)
	rules := ParseSection1(sections[0])
	lines := ParseSection2(sections[1])

	correctLines := [][]int{}
	for _, line := range lines {
		g := createGraph(rules, line)
		sorted, err := graph.TopologicalSort(g)
		if err != nil {
			fmt.Printf("Error sorting line: %#v, err: %v\n", line, err)
		}

		compValue := slices.Compare(line, sorted)
		if compValue != 0 {
			correctLines = append(correctLines, sorted)
		}
	}

	result := 0
	for _, line := range correctLines {
		idx := len(line) / 2
		result += line[idx]
	}
	return result
}
