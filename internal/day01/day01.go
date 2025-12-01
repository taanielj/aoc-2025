package day01

import (
	"aoc2025/internal/utils"
	"fmt"
)

func SolveDay01() {
	initialPosition := 50
	initialPassword := 0
	lines, err := utils.ReadLinesFromFile("inputs/day01.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	position := initialPosition
	password := initialPassword
	for _, line := range lines {
		direction, steps, err := getDirectionAndSteps(line)
		if err != nil {
			fmt.Println("Error parsing instruction:", err)
			return
		}
		position, err = rotate(position, direction, steps)
		if err != nil {
			fmt.Println("Error rotating:", err)
			return
		}
		password = updatePassword(position, password)
	}
	fmt.Printf("Final Position: %d, Password: %d\n", position, password)
}

func getDirectionAndSteps(instruction string) (string, int, error) {
	// line format is like "R10" or "L5"
	if len(instruction) < 2 {
		return "", 0, fmt.Errorf("invalid instruction: %s", instruction)
	}
	direction := string(instruction[0])
	var steps int
	_, err := fmt.Sscanf(instruction[1:], "%d", &steps)
	if err != nil {
		return "", 0, fmt.Errorf("invalid steps in instruction: %s", instruction)
	}
	return direction, steps, nil
}

func rotate(position int, direction string, steps int) (int, error) {
	switch direction {
	case "R":
		position = (position + steps) % 100
	case "L":
		position = (position - steps + 100) % 100
	default:
		return 0, fmt.Errorf("invalid direction: %s", direction)
	}
	return position, nil
}

func updatePassword(position int, current_password int) int {
	// if final position is 0, then increment current code by 1, else return current code
	if position == 0 {
		return current_password + 1
	}
	return current_password
}
