package day05

import (
	"aoc2025/internal/utils"
	"fmt"
	"sort"
	"strconv"
)

type freshRange struct {
	start int
	end   int
}

func SolveDay05() {
	lines, err := utils.ReadLinesFromFile("inputs/day05.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		panic(err)
	}
	ranges := getFreshRanges(lines)
	items := getItems(lines)
	freshItemCount := countFreshItems(items, ranges)
	fmt.Println("Day 05 Part 1 Solution:")
	fmt.Printf("Number of ranges: %d\n", len(ranges))
	fmt.Printf("Number of items: %d\n", len(items))
	fmt.Printf("Number of fresh items: %d\n", freshItemCount)

	mergedRanges := mergeRanges(ranges)
	possibleFreshItemCount := countPossibleFreshItems(mergedRanges)

	fmt.Println("Day 05 Part 2 Solution:")
	fmt.Printf("Number of merged ranges: %d\n", len(mergedRanges))
	fmt.Printf("Number of possible fresh items: %d\n", possibleFreshItemCount)
}

func getFreshRanges(lines []string) []freshRange {
	freshRanges := []freshRange{}
	for _, line := range lines {
		parts := utils.SplitString(line, "-")
		if len(parts) != 2 {
			continue
		}
		newRange := freshRange{
			start: utils.MustAtoi(parts[0]),
			end:   utils.MustAtoi(parts[1]),
		}
		freshRanges = append(freshRanges, newRange)

	}
	return freshRanges
}

func mergeRanges(ranges []freshRange) []freshRange {
	// sort by start
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := []freshRange{ranges[0]}
	for _, currentRange := range ranges[1:] {
		lastRange := &merged[len(merged)-1]
		if currentRange.start <= lastRange.end {
			if currentRange.end > lastRange.end {
				lastRange.end = currentRange.end
			}

		} else {
			merged = append(merged, currentRange)
		}
	}

	return merged
}

func getItems(lines []string) []int {
	items := []int{}
	for _, line := range lines {
		item, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		items = append(items, item)
	}
	return items
}

func countFreshItems(items []int, freshRanges []freshRange) int {
	freshCount := 0
	for _, item := range items {
		for _, fr := range freshRanges {
			if item >= fr.start && item <= fr.end {
				freshCount++
				break
			}
		}
	}
	return freshCount
}

func countPossibleFreshItems(freshRanges []freshRange) int {
	count := 0
	for _, fr := range freshRanges {
		count += fr.end - fr.start + 1
	}
	return count
}
