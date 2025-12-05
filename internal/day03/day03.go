package day03

import (
	"aoc2025/internal/utils"
	"fmt"
)

func SolveDay03() {
	lines, err := utils.ReadLinesFromFile("inputs/day03.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	totalJoltage := 0
	for _, line := range lines {
		bank, err := parseLine(line)
		if err != nil {
			fmt.Println("Error parsing joltage:", err)
			return
		}
		joltage := getJoltage(bank)

		totalJoltage += joltage
	}
	fmt.Printf("Total Joltage: %d\n", totalJoltage)

}

func parseLine(line string) ([]int, error) {
	var bank []int
	for _, char := range line {
		var val int
		_, err := fmt.Sscanf(string(char), "%d", &val)
		if err != nil {
			return nil, fmt.Errorf("invalid character in joltage bank: %s", string(char))
		}
		bank = append(bank, val)
	}
	return bank, nil
}

func getJoltage(bank []int) int {
	maxJoltage := 0
	maxSeen := bank[len(bank)-1]
	for i := len(bank) - 2; i >= 0; i-- { // we traverse backwards, and start at len-2 to avoid double counting the last element
		if joltage := bank[i]*10 + maxSeen; joltage > maxJoltage {
			maxJoltage = joltage
		}
		if bank[i] > maxSeen {
			maxSeen = bank[i]
		}
	}
	return maxJoltage

}
