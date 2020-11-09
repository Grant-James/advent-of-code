package main

import (
	"fmt"
)

var puzzleInput = []int {1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,9,1,19,1,19,5,23,1,9,23,27,2,27,6,31,1,5,31,35,2,9,35,39,2,6,39,43,2,43,13,47,2,13,47,51,1,10,51,55,1,9,55,59,1,6,59,63,2,63,9,67,1,67,6,71,1,71,13,75,1,6,75,79,1,9,79,83,2,9,83,87,1,87,6,91,1,91,13,95,2,6,95,99,1,10,99,103,2,103,9,107,1,6,107,111,1,10,111,115,2,6,115,119,1,5,119,123,1,123,13,127,1,127,5,131,1,6,131,135,2,135,13,139,1,139,2,143,1,143,10,0,99,2,0,14,0}
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
