package main

import (
	"fmt"
	"strconv"
)

var numRange = []int {265275, 781584}
var isPartTwo = true

func main() {
	err := run()
	if err != nil {
	    fmt.Printf("err: %v", err)
	    return
	}
}

func run() error {
	counter := 0
	for i := numRange[0]; i <= numRange[1]; i++ {
		code := convertCodeToIntArr(i)

		isValid := false
		prevNum := -1
		sameDigitCounter := 0
		hasPerfectDouble := false

		for i, num := range code {
			if num - prevNum == 0 {
				sameDigitCounter++

				if isPartTwo {
					if i == 5 && sameDigitCounter == 1 {
						hasPerfectDouble = true
						isValid = true
						break
					}

					if sameDigitCounter >= 2 && !hasPerfectDouble {
						isValid = false
						prevNum = num
						continue
					}
				}

				isValid = true
				prevNum = num
				continue
			} else if sameDigitCounter == 1 {
				sameDigitCounter = 0
				hasPerfectDouble = true
			} else if sameDigitCounter > 1 {
				sameDigitCounter = 0
				hasPerfectDouble = false
			}

			if num - prevNum < 0 {
				isValid = false
				break
			}

			prevNum = num
		}

		if isValid {
			counter++
		}
	}

	fmt.Print(counter)

	return nil
}

func convertCodeToIntArr(code int) []int {
	codeStr := strconv.Itoa(code)

	var slice []int
	for _, digit := range codeStr {
		slice = append(slice, int(digit)-int('0'))
	}
	
	return slice
}
