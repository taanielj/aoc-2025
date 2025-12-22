package day06

import (
	"aoc2025/internal/utils"
	"fmt"
	// "strings"
)

// https://gobyexample.com/enums

type Operator int

const (
	Unset Operator = iota
	Add
	Multiply
)

type mathProblem struct {
	operands     []int
	operator     Operator
	position     int
	operandCount int
}

func SolveDay06() {
	lines, err := utils.ReadLinesFromFile("inputs/day06.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		panic(err)
	}
	problems := getProblems(lines)
	totalSum := 0
	for _, problem := range problems {
		result := solveProblem(problem)
		totalSum += result
	}
	fmt.Println("Day 06 part 2 solution:", totalSum)

}

func getProblems(lines []string) []mathProblem {

	problems := parseOperatorLine(lines[len(lines)-1])
	lines = lines[:len(lines)-1]
	for idx := range problems {
		operands := make([]int, 0)
		for i := problems[idx].operandCount - 1; i >= 0; i-- {
			operand := ""
			for _, line := range lines {
				char := rune(line[problems[idx].position+i])
				if char == ' ' {
					continue
				}
				operand += string(char)
			}
			value := utils.MustAtoi(operand)
			operands = append(operands, value)
		}
		problems[idx].operands = operands
	}

	return problems
}

func parseOperator(char rune) Operator {
	switch char {
	case '+':
		return Add
	case '*':
		return Multiply
	default:
		return Unset
	}
}

func parseOperatorLine(line string) []mathProblem {
	// parse line char by char, don't trim spaces, otherwise we lose position info
	problems := make([]mathProblem, 0)
	currentProblem := mathProblem{}
	currentProblem.operator = Unset
	currentProblem.operandCount = 1
	for i := len(line) - 1; i >= 0; i-- {
		operator := parseOperator(rune(line[i]))
		if operator == Unset {
			currentProblem.operandCount++
			continue
		}
		currentProblem.position = i
		currentProblem.operator = operator
		problems = append(problems, currentProblem)
		currentProblem = mathProblem{}
		currentProblem.operator = Unset
		currentProblem.operandCount = 1
		i--

	}
	currentProblem.position = 0
	currentProblem.operator = parseOperator(rune(line[0])) 
	problems = append(problems, currentProblem)
	return problems
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
		println("Unset operator in problem:", problem.operator)
		return 0
	}
	return result
}
