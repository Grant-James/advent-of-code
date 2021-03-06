package main

import (
	"fmt"
	"strconv"
)

var puzzleInput = []int {3,225,1,225,6,6,1100,1,238,225,104,0,1102,17,65,225,102,21,95,224,1001,224,-1869,224,4,224,1002,223,8,223,101,7,224,224,1,224,223,223,101,43,14,224,1001,224,-108,224,4,224,102,8,223,223,101,2,224,224,1,223,224,223,1101,57,94,225,1101,57,67,225,1,217,66,224,101,-141,224,224,4,224,102,8,223,223,1001,224,1,224,1,224,223,223,1102,64,34,225,1101,89,59,225,1102,58,94,225,1002,125,27,224,101,-2106,224,224,4,224,102,8,223,223,1001,224,5,224,1,224,223,223,1102,78,65,225,1001,91,63,224,101,-127,224,224,4,224,102,8,223,223,1001,224,3,224,1,223,224,223,1102,7,19,224,1001,224,-133,224,4,224,102,8,223,223,101,6,224,224,1,224,223,223,2,61,100,224,101,-5358,224,224,4,224,102,8,223,223,101,3,224,224,1,224,223,223,1101,19,55,224,101,-74,224,224,4,224,102,8,223,223,1001,224,1,224,1,224,223,223,1101,74,68,225,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,107,677,677,224,102,2,223,223,1006,224,329,1001,223,1,223,1008,226,677,224,102,2,223,223,1006,224,344,1001,223,1,223,7,226,677,224,102,2,223,223,1005,224,359,1001,223,1,223,8,226,226,224,102,2,223,223,1006,224,374,1001,223,1,223,1007,226,226,224,102,2,223,223,1006,224,389,101,1,223,223,8,677,226,224,1002,223,2,223,1005,224,404,101,1,223,223,1108,677,226,224,102,2,223,223,1006,224,419,1001,223,1,223,1108,226,677,224,102,2,223,223,1006,224,434,101,1,223,223,1108,677,677,224,1002,223,2,223,1005,224,449,101,1,223,223,1008,677,677,224,1002,223,2,223,1006,224,464,101,1,223,223,7,677,226,224,1002,223,2,223,1006,224,479,101,1,223,223,108,677,677,224,1002,223,2,223,1005,224,494,101,1,223,223,107,226,677,224,1002,223,2,223,1006,224,509,101,1,223,223,107,226,226,224,102,2,223,223,1006,224,524,1001,223,1,223,1107,226,677,224,1002,223,2,223,1006,224,539,101,1,223,223,1008,226,226,224,102,2,223,223,1006,224,554,1001,223,1,223,8,226,677,224,1002,223,2,223,1006,224,569,101,1,223,223,1007,677,677,224,102,2,223,223,1005,224,584,1001,223,1,223,1107,677,226,224,1002,223,2,223,1006,224,599,101,1,223,223,7,226,226,224,1002,223,2,223,1005,224,614,101,1,223,223,108,677,226,224,1002,223,2,223,1005,224,629,1001,223,1,223,108,226,226,224,1002,223,2,223,1005,224,644,101,1,223,223,1007,677,226,224,1002,223,2,223,1006,224,659,101,1,223,223,1107,226,226,224,102,2,223,223,1005,224,674,1001,223,1,223,4,223,99,226}
var morphingInput = make([]int, 0)

var step = 1

func main() {
	morphingInput = puzzleInput

	err := run()
	if err != nil {
	    fmt.Printf("err: %v", err)
	    return
	}

	fmt.Printf("[OUTPUT] - %v\n", morphingInput)
}

func run() error {
	for i := 0; i < len(morphingInput); i += step {
		// Returns 4 number instruction instructions [0, 0, 0, 00]
		instruction := getOpcodeInstructions(morphingInput[i])
		opcode := instruction[3]
		param1Mode := instruction[2]
		param2Mode := instruction[1]

		// Halt on opcode 99
		if opcode == 99 {
			fmt.Printf("\n[OPCODE 99] Halting.\n")
			break
		}

		// Input/Output opcodes handles=d and returned first
		if opcode == 3 || opcode == 4 {
			if i + 1 >= len(morphingInput) {
				fmt.Printf("\nNot enough parameters to continue. Halting.\n")
				break
			}

			// Set next step
			step = 2

			// Output is always in position mode
			outPos := morphingInput[i + 1]

			var value int
			switch opcode {
			case 3:
				fmt.Println("[OPCODE 3] Enter input: ")
				_, err := fmt.Scan(&value)
				if err != nil {
					return err
				}
				morphingInput[outPos] = value
				continue
			case 4:
				if param1Mode == 0 {
					outPos = morphingInput[outPos]
				}
				fmt.Printf("%v ", outPos)
				continue
			}
		}

		// Validate rest of array length
		if i + 3 >= len(morphingInput) {
			fmt.Printf("\nNot enough parameters to continue. Halting.\n")
			break
		}

		// Get value / value of position for params
		var param1 = morphingInput[i + 1]
		if param1Mode == 0 {
			param1 = morphingInput[param1]
		}
		var param2 = morphingInput[i + 2]
		if param2Mode == 0 {
			param2 = morphingInput[param2]
		}

		// Output is always in position mode
		outPos := morphingInput[i + 3]

		// Jump if true - jump to instruction pointer if non-zero
		step = 3
		if opcode == 5 {
			if param1 != 0 {
				i = param2 - step
			}

			continue
		}

		// Jump if false - jump to instruction pointer if is zero
		if opcode == 6 {
			if param1 == 0 {
				i = param2 - step
			}

			continue
		}

		// Set output
		step = 4
		morphingInput[outPos] = getOpcodeValue(opcode, param1, param2)
	}

	return nil
}

func getOpcodeValue(opcode int, num1 int, num2 int) int {
	switch opcode {
	case 1:
		return num1 + num2
	case 2:
		return num1 * num2
	case 7:
		if num1 < num2 {
			return 1
		}
		return 0
	case 8:
		if num1 == num2 {
			return 1
		}
		return 0
	default:
		return -69
	}
}

func getOpcodeInstructions(num int) []int {
	codeStr := strconv.Itoa(num)

	var slice []int
	for _, digit := range codeStr {
		slice = append(slice, int(digit)-int('0'))
	}

	if num < 100 {
		slice = append([]int {0,0,0}, num)
	} else if num < 1000 {
		slice = append([]int {0,0, slice[0]}, sliceToNum(slice[1:3]))
	} else if num < 10000 {
		slice = append([]int {0, slice[0], slice[1]}, sliceToNum(slice[2:4]))
	}

	return slice
}

func sliceToNum(slice []int) int {
	var numStr string
	for _, num := range slice {
		numStr += strconv.Itoa(num)
	}
	num, _ := strconv.Atoi(numStr)
	return num
}
