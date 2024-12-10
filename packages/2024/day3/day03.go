package day3

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/mrsomia/advent-of-code/packages/utils"
)

func parseInput(input string) []string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return lines
}

func SolvePartA(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseInput(input)
	re, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	if err != nil {
		log.Fatalf("regexp failed to compile, %v\n", err)
	}

	m := make([][]int, 0)
	for _, line := range lines {
		matches := re.FindAllSubmatch([]byte(line), -1)
		// fmt.Printf("matches:%q\n", matches)
		for _, matchGroup := range matches {
			n := make([]int, 0)
			num1s := string(matchGroup[1])
			num1, err := strconv.Atoi(num1s)
			if err != nil {
				fmt.Printf("Failed to convert num: %v\n", err)
			}
			n = append(n, num1)

			num2s := string(matchGroup[2])
			num2, err := strconv.Atoi(num2s)
			if err != nil {
				fmt.Printf("Failed to convert num: %v\n", err)
			}
			n = append(n, num2)

			m = append(m, n)
		}
	}

	result := 0
	for _, mm := range m {
		result = result + (mm[0] * mm[1])
	}

	return result
}

func SolvePartB(filename string) int {
	input := utils.OpenFile(filename)
	lines := parseInput(input)

	re, err := regexp.Compile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	if err != nil {
		log.Fatalf("regexp failed to compile, %v\n", err)
	}
	// numRe, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	// if err != nil {
	// 	log.Fatalf("regexp failed to compile, %v\n", err)
	// }

	muls := make([]string, 0)
	m := make([][]int, 0)
	for _, line := range lines {
		matches := re.FindAllString(line, -1)
		muls = append(muls, matches...)
	}

	fmt.Printf("MULS:\n %#v\n", muls)
	enabled := true
	for _, mul := range muls {
		if mul == "do()" {
			enabled = true
			continue
		}
		if mul == "don't()" {
			enabled = false
			continue
		}

		if !enabled {
			continue
		}
		matchGroup := strings.Split(mul[4:len(mul)-1], ",")
		fmt.Printf("matchGroup:%s\n", matchGroup)
		n := make([]int, 0)
		num1s := string(matchGroup[0])
		num1, err := strconv.Atoi(num1s)
		if err != nil {
			fmt.Printf("Failed to convert num: %v\n", err)
		}
		n = append(n, num1)

		num2s := string(matchGroup[1])
		num2, err := strconv.Atoi(num2s)
		if err != nil {
			fmt.Printf("Failed to convert num: %v\n", err)
		}
		n = append(n, num2)

		m = append(m, n)
	}

	result := 0
	for _, mm := range m {
		result = result + (mm[0] * mm[1])
	}

	return result
}
