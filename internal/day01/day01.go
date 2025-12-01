package day01

import (
	"aoc2025/internal/utils"
	"fmt"
)

type Instruction struct {
	Direction string
	Steps     int
}

type State struct {
	Position int
	Password int
}

func SolveDay01() {
	initialPosition := 50
	initialPassword := 0
	lines, err := utils.ReadLinesFromFile("inputs/day01.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	state := State{Position: initialPosition, Password: initialPassword}
	for _, line := range lines {
		instruction, err := getInstruction(line)
		if err != nil {
			fmt.Println("Error parsing instruction:", err)
			return
		}
		state, err = rotate(state, instruction)
		if err != nil {
			fmt.Println("Error rotating:", err)
			return
		}

	}
	fmt.Printf("Final Position: %d, Final Password: %d\n", state.Position, state.Password)
}

func getInstruction(instruction string) (Instruction, error) {
	if len(instruction) < 2 {
		return Instruction{}, fmt.Errorf("invalid instruction format: %s", instruction)
	}
	direction := string(instruction[0])
	var steps int
	_, err := fmt.Sscanf(instruction[1:], "%d", &steps)
	if err != nil {
		return Instruction{}, fmt.Errorf("invalid steps in instruction: %s", instruction)
	}
	return Instruction{Direction: direction, Steps: steps}, nil
}

func rotate(state State, instruction Instruction) (State, error) {
	crossings := 0
	for i := 0; i < instruction.Steps; i++ {
		switch instruction.Direction {
		case "R":
			state.Position++
			if state.Position > 99 {
				state.Position = 0
			}
		case "L":
			state.Position--
			if state.Position < 0 {
				state.Position = 99
			}
		default:
			return State{}, fmt.Errorf("invalid direction: %s", instruction.Direction)
		}
		if state.Position == 0 {
			crossings++
		}
	}
	state.Password += crossings
	return state, nil

}
