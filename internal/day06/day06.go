package day06

import (
	"aoc2025/internal/utils"
	"fmt"
	// "sort"
	// "strconv"
	"strings"
)

// https://gobyexample.com/enums

type Operator int

const (
	Add Operator = iota
	Multiply
)

var operationSymbols = map[Operator]string{
	Add:      "+",
	Multiply: "*",
}

type mathProblem struct {
	operands []int
	operator Operator
}

func SolveDay06() {
	lines, err := utils.ReadLinesFromFile("inputs/day06.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		panic(err)
	}
	problems := getPart1Problems(lines)
	totalSum := 0
	for _, problem := range problems {
		result := solveProblem(problem)
		totalSum += result
	}
	fmt.Println("Day 06 part 1 solution:", totalSum)
}

func getPart1Problems(lines []string) []mathProblem {
	// last line is operators
	operatorLine := lines[len(lines)-1]
	operators := parseOperatorLine(operatorLine)
	problems := make([]mathProblem, 0, len(operators))
	for i, operator := range operators {
		problem := mathProblem{
			operator: operator,
			operands: make([]int, 0),
		}
		for _, line := range lines[:len(lines)-1] {
			fields := strings.Fields(line)
			problem.operands = append(problem.operands, utils.MustAtoi(fields[i]))
		}
		problems = append(problems, problem)
	}
	return problems
}

func parseOperatorLine(line string) []Operator {
	symbols := strings.Fields(line)
	operators := make([]Operator, 0, len(symbols))
	for _, symbol := range symbols {
		for op, sym := range operationSymbols {
			if sym == symbol {
				operators = append(operators, op)
			}
		}
	}
	return operators
}

func solveProblem(problem mathProblem) int {
	result := 0
	switch problem.operator {
	case Add:
		for _, operand := range problem.operands {
			result += operand
		}
	case Multiply:
		result = 1
		for _, operand := range problem.operands {
			result *= operand
		}
	default:
		panic("Unknown operator")
	}
	return result
}
