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
		bankJoltage, err := getBankJoltage(bank, 12) // change second argument for part 1 or part 2
		if err != nil {
			fmt.Println("Error calculating joltage:", err)
			return
		}

		totalJoltage += bankJoltage
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

// part 1
// func getBankJoltage(bank []int) int {
//	maxBankJoltage := 0
//	maxSeen := bank[len(bank)-1]
//	for i := len(bank) - 2; i >= 0; i-- { // we traverse backwards, and start at len-2 to avoid double counting the last element
//		if bankJoltage := bank[i]*10 + maxSeen; bankJoltage > maxBankJoltage {
//			maxBankJoltage = bankJoltage
//		}
//		if bank[i] > maxSeen {
//			maxSeen = bank[i]
//		}
//	}
//	return maxBankJoltage
//
// }

// part 2 - uses monotonic stack to keep largest possible number
func getBankJoltage(bank []int, batteries int) (int, error) {
	if batteries > len(bank) {
		return 0, fmt.Errorf("not enough batteries in bank: have %d, need %d", len(bank), batteries)
	}
	kept := []int{}
	skipsRemaining := len(bank) - batteries
	for _, joltage := range bank {
		for len(kept) > 0 && kept[len(kept)-1] < joltage && skipsRemaining > 0 {
			kept = kept[:len(kept)-1]
			skipsRemaining--
		}
		kept = append(kept, joltage)
	}
	kept = kept[:batteries] // trim from end if input was descending

	bankJoltage := 0
	for _, joltage := range kept {
		bankJoltage = bankJoltage*10 + joltage
	}
	return bankJoltage, nil
}
