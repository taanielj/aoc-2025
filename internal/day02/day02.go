package day02

import (
	"aoc2025/internal/utils"
	"fmt"
	"strings"
)

func SolveDay02() {
	input, err := utils.ReadLinesFromFile("inputs/day02.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	ranges := input[0]
	fmt.Println("Ranges:", ranges)
	rangeList := utils.SplitString(ranges, ",")
	invalidSum := 0
	for _, r := range rangeList {
		invalidIds, err := getInvalidIds(r)
		if err != nil {
			fmt.Println("Error parsing range:", err)
			return
		}
		for _, id := range invalidIds {
			invalidSum += id
		}
	}
	fmt.Printf("Sum of invalid IDs: %d\n", invalidSum)
}

func getInvalidIds(idRange string) ([]int, error) {
	var start, end int
	_, err := fmt.Sscanf(idRange, "%d-%d", &start, &end)
	if err != nil {
		return nil, fmt.Errorf("invalid range format: %s", idRange)
	}
	fmt.Printf("Processing range: %d-%d\n", start, end)
	var invalidIds []int
	for i := start; i <= end; i++ {
		if isInvalidID(i) {
			invalidIds = append(invalidIds, i)
		} else if isInvalidIDSimple(i) { // cross-check with part 1 simple check, should not happen
			return nil, fmt.Errorf("ID %d should be invalid by simple check but passed complex check", i)
		}
	}

	fmt.Printf("Invalid IDs in range %s: %v\n", idRange, invalidIds)
	return invalidIds, nil
}

func isInvalidID(id int) bool {
	idStr := fmt.Sprintf("%d", id)
	idLength := len(idStr)
	for chunkSize := 1; chunkSize <= idLength/2; chunkSize++ {
		if idLength%chunkSize != 0 {
			continue
		}
		chunk := idStr[:chunkSize]
		repeated := true
		for i := chunkSize; i < idLength; i += chunkSize {
			if idStr[i:i+chunkSize] != chunk {
				repeated = false
				break
			}
		}
		if repeated {
			return true
		}
	}
	return false
}

func isInvalidIDSimple(id int) bool {
	// if a sequence appears twice in a row in the ID, it is invalid
	idStr := fmt.Sprintf("%d", id)
	idLength := len(idStr)
	if idLength%2 == 0 {
		leftHalf := idStr[:idLength/2]
		rightHalf := idStr[idLength/2:]
		if strings.Compare(leftHalf, rightHalf) == 0 {
			return true
		}
	}
	return false
}
