package main

import (
	"fmt"
)

var puzzleInput = []int {1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,10,1,19,1,19,9,23,1,23,13,27,1,10,27,31,2,31,13,35,1,10,35,39,2,9,39,43,2,43,9,47,1,6,47,51,1,10,51,55,2,55,13,59,1,59,10,63,2,63,13,67,2,67,9,71,1,6,71,75,2,75,9,79,1,79,5,83,2,83,13,87,1,9,87,91,1,13,91,95,1,2,95,99,1,99,6,0,99,2,14,0,0}
var expectedNumber = 19690720

var morphingInput = make([]int, 0)

func main() {
	isFound := findVerb(100)

	fmt.Printf("\n-- Part 1 --\nOutput: %v", morphingInput)

	if isFound {
		fmt.Printf("\n-- Part 2 --\nFound number: %v\nVerb: %v\nNoun: %v\nOutput: %v",
			morphingInput[0],
			morphingInput[1],
			morphingInput[2],
			100 * morphingInput[1] + morphingInput[2],
		)
	}

	fmt.Print("\n\nOPCODE 99: Program terminated\n")
}

func findVerb(searchSquared int) bool {
	for i := 0; i < searchSquared; i++ {
		for j := 0; j < searchSquared; j++ {
			morphingInput = append([]int {puzzleInput[0], i, j}, puzzleInput[3:]...)

			isFound, err := run()
			if err != nil {
				fmt.Printf("err: %v", err)
				return false
			}

			if isFound {
				return true
			}
		}
	}

	return false
}

func run() (bool, error) {
	for i := 0; i < len(morphingInput); i += 4 {
		if morphingInput[i] == 99 {
			break
		}

		if i + 3 >= len(morphingInput) {
			break
		}

		opcode := morphingInput[i]
		inPos1 := morphingInput[i + 1]
		inPos2 := morphingInput[i + 2]
		outPos := morphingInput[i + 3]

		if outPos >= len(morphingInput) {
			break
		}
		morphingInput[outPos] = getOpcodeValue(opcode, morphingInput[inPos1], morphingInput[inPos2])

		if morphingInput[outPos] == expectedNumber {
			return true, nil
		}
	}

	return false, nil
}

func getOpcodeValue(opcode int, num1 int, num2 int) int {
	switch opcode {
	case 1:
		return num1 + num2
	case 2:
		return num1 * num2
	default:
		return -69
	}
}
